package main

import (
	"fmt"
	"log"
	"net"
	"os"

	"github.com/fakhripraya/whatsapp-service/data"
	"github.com/fakhripraya/whatsapp-service/entities"
	protos "github.com/fakhripraya/whatsapp-service/protos/whatsapp"
	"github.com/fakhripraya/whatsapp-service/server"
	"github.com/hashicorp/go-hclog"
	"github.com/joho/godotenv"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

var err error

func main() {
	logger := hclog.Default()

	// load configuration from env file
	logger.Info("Loading env")
	err = godotenv.Load(".env")

	if err != nil {
		// log the fatal error if load env failed
		log.Fatal(err)
	}

	// Initialize app configuration
	logger.Info("Initialize application configuration")
	var appConfig entities.Configuration
	err = data.ConfigInit(&appConfig)

	// create a whatsapp login / fetch current session
	logger.Info("Creating a new WhatsApp connection")
	waConfig, err := data.NewWA(logger)
	if err != nil {
		logger.Error("Error while establishing WhatsApp connection", "error", err.Error())
		log.Fatal(err)
	}

	// create a new gRPC server, use WithInsecure to allow http connections
	logger.Info("Creating a new gRPC server")
	gs := grpc.NewServer()

	// create an instance of the WA sender server
	logger.Info("Creating a new WA sender instance")
	sender := server.NewSender(logger, waConfig)

	// register the WA sender server
	logger.Info("Registering the new WA sender into the gRPC server")
	protos.RegisterWhatsAppServer(gs, sender)

	// register the reflection service which allows clients to determine the methods
	// for this gRPC service
	logger.Info("Registering reflection service")
	reflection.Register(gs)

	// create a TCP socket for inbound server connections
	logger.Info("Creating TCP socket")
	listener, err := net.Listen("tcp", fmt.Sprintf(":"+appConfig.AppConfig.Port))
	if err != nil {
		logger.Error("Unable to create listener", "error", err.Error())
		os.Exit(1)
	}

	// listen for requests
	logger.Info("Successfully creating TCP socket")
	gs.Serve(listener)
}
