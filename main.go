package main

import (
	"github.com/luispfcanales/daemon-service-oti/gui"
	"github.com/luispfcanales/daemon-service-oti/services/stream"
)

func main() {
	streaming := stream.NewStreamWS("ws://40.0.2.2:3000/stream/desktop/computer/%s", "http://40.0.2.2")
	gui.Run(streaming)
}
