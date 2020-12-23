package server

import (
	"context"

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
	// set the WA target info
	text := whatsapp.TextMessage{
		Info: whatsapp.MessageInfo{
			RemoteJid: wr.RemoteJid,
		},
		Text: wr.Text,
	}

	// send the WA text
	_, err := sender.config.Wac.Send(text)
	if err != nil {
		return &protos.WAResponse{
				ErrorCode:    "404",
				ErrorMessage: err.Error()},
			nil
	}

	return nil, nil
}
