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

	log.Print("Redirected standard library")
	log.Print("Starting server")
	wantFrontEndBind := fmt.Sprintf("tcp://*:%q", frontendPort)
	wantFrontEndConnect := fmt.Sprintf("tcp://%q:%q", frontendHost, frontendPort)
	wantBackEndBind := fmt.Sprintf("tcp://*:%q", backendPort)
	wantBackEndConnect := fmt.Sprintf("tcp://%q:%q", backendHost, backendPort)
	server := efogIotEdgeHubServer.NewServer(&wantBackEndBind, &wantBackEndConnect, &wantFrontEndBind, &wantFrontEndConnect)

	log.Printf("Frontend endpoint %q", wantFrontEndBind)
	log.Printf("Frontend endpoint %q", wantFrontEndConnect)
	log.Printf("Backend endpoint %q", wantBackEndBind)
	log.Printf("Backend endpoint %q", wantBackEndConnect)

	server.Run()
}
