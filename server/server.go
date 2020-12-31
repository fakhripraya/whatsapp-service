package server

import (
	"context"
	"strconv"
	"strings"

	"github.com/Rhymen/go-whatsapp"
	protos "github.com/fakhripraya/whatsapp-service/protos/whatsapp"

	"github.com/fakhripraya/whatsapp-service/data"
	"github.com/hashicorp/go-hclog"
)

// Sender is a gRPC server, it implements the methods defined by the WhatsAppSenderServer interface
type Sender struct {
	protos.UnimplementedWhatsAppServer
	logger hclog.Logger
	config *data.Whatsapp
}

// NewSender creates a new WA Sender server
func NewSender(logger hclog.Logger, config *data.Whatsapp) *Sender {
	newSender := &Sender{
		logger: logger,
		config: config}

	return newSender
}

// SendWhatsApp is a function to send a WhatsApp message based on the WhatsAppRequest
func (sender *Sender) SendWhatsApp(ctx context.Context, wr *protos.WARequest) (*protos.WAResponse, error) {

	// create the existance result instance
	exResult := &data.ExistanceResult{}

	// filter the WhatsApp number into Indonesian based WhatsApp number
	if strings.HasPrefix(wr.RemoteJid, "+") {
		wr.RemoteJid = strings.Replace(wr.RemoteJid, "+", "", 1)
	} else if strings.HasPrefix(wr.RemoteJid, "0") {
		wr.RemoteJid = strings.Replace(wr.RemoteJid, "0", "62", 1)
	}

	// check if WA number whether exist or or not
	ok, err := sender.config.Wac.Exist(wr.RemoteJid)
	if err != nil {
		return &protos.WAResponse{
				ErrorCode:    "500",
				ErrorMessage: err.Error()},
			nil
	}

	exist := <-ok

	// parse the existance check result to the given instance
	err = data.UnmarshalJSON(exist, exResult)
	if err != nil {
		return &protos.WAResponse{
				ErrorCode:    "400",
				ErrorMessage: err.Error()},
			nil
	}

	// return bad request if WA number doesn't exist
	if strconv.Itoa(exResult.Status) != "200" {
		return &protos.WAResponse{
				ErrorCode:    "400",
				ErrorMessage: "Nomor WhatsApp tidak dapat ditemukan"},
			nil
	}

	// set the WA target info
	text := whatsapp.TextMessage{
		Info: whatsapp.MessageInfo{
			RemoteJid: wr.RemoteJid,
		},
		Text: wr.Text,
	}

	// send the WA text
	_, err = sender.config.Wac.Send(text)
	if err != nil {
		return &protos.WAResponse{
				ErrorCode:    "400",
				ErrorMessage: err.Error()},
			nil
	}

	// send the ok response if succeed
	return &protos.WAResponse{
			ErrorCode:    "200",
			ErrorMessage: "Status Accepted"},
		nil
}
