package main

import (
	"github.com/luispfcanales/daemon-service-oti/gui"
	"github.com/luispfcanales/daemon-service-oti/services/stream"
)

const (
	IP      string = "192.168.0.10"
	PORT_IP string = ":3000"
)

func main() {
	streaming := stream.NewStreamWS("ws://"+IP+PORT_IP+"/stream/desktop/computer/%s", "http://"+IP)
	gui.Run(streaming)
}
