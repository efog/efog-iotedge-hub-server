package main

import (
	"fmt"
	"log"
	"os"

	efogIotEdgeHubServer "github.com/efog/efog-iotedge-hub"
	zap "go.uber.org/zap"
)

func main() {
	logger := zap.NewExample()
	defer logger.Sync()

	undo := zap.RedirectStdLog(logger)
	defer undo()

	var backendHost string
	var backendPort string
	var frontendHost string
	var frontendPort string

	backendHost = os.Getenv("BACKEND_HOST")
	backendPort = os.Getenv("BACKEND_PORT")
	frontendHost = os.Getenv("FRONTEND_HOST")
	frontendPort = os.Getenv("FRONTEND_PORT")

	if backendHost == "" {
		backendHost = "localhost"
	}
	if backendPort == "" {
		backendPort = "56789"
	}
	if frontendHost == "" {
		frontendHost = "localhost"
	}
	if frontendPort == "" {
		frontendPort = "12345"
	}

	log.Printf("%s", backendHost)
	log.Printf("%s", backendPort)
	log.Printf("%s", frontendHost)
	log.Printf("%s", frontendPort)

	log.Print("Redirected standard library")
	log.Print("Starting server")
	wantFrontEndBind := fmt.Sprintf("tcp://*:%s", frontendPort)
	wantFrontEndConnect := fmt.Sprintf("tcp://%s:%s", frontendHost, frontendPort)
	wantBackEndBind := fmt.Sprintf("tcp://*:%s", backendPort)
	wantBackEndConnect := fmt.Sprintf("tcp://%s:%s", backendHost, backendPort)
	server := efogIotEdgeHubServer.NewServer(&wantBackEndBind, &wantBackEndConnect, &wantFrontEndBind, &wantFrontEndConnect)

	log.Printf("Frontend bind endpoint %s", wantFrontEndBind)
	log.Printf("Frontend connect endpoint %s", wantFrontEndConnect)
	log.Printf("Backend bind endpoint %s", wantBackEndBind)
	log.Printf("Backend connect endpoint %s", wantBackEndConnect)

	server.Run()
}
