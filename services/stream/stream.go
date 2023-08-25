package stream

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net"
	"os"
	"path/filepath"
	"time"

	"github.com/go-toast/toast"
	"github.com/google/uuid"
	"github.com/luispfcanales/daemon-service-oti/entity"
	"github.com/luispfcanales/daemon-service-oti/model"
	"golang.org/x/net/websocket"
)

type StreamSrv struct {
	ID            string
	urlWS         string
	origin        string
	ws            *websocket.Conn
	tryConnection *time.Ticker
	mailBox       chan []byte
	signalConnect chan struct{}
	errorChan     chan error
}

var (
	EVENT_NOTIFICATION string = "notify"
	EVENT_LOAD_INFO    string = "load-info-system"

	EVENT_LOADED string = "loaded-info"
)

func NewStreamWS(UrlWebsocket, UrlOrigin string) *StreamSrv {
	srv := &StreamSrv{
		ID:            uuid.New().String(),
		urlWS:         UrlWebsocket,
		origin:        UrlOrigin,
		tryConnection: time.NewTicker(time.Second * 7),
		mailBox:       make(chan []byte, 10),
		signalConnect: make(chan struct{}, 1),
		errorChan:     make(chan error, 1),
	}
	go srv.Start()
	return srv
}

func (st *StreamSrv) Start() {
	for {
		select {
		case <-st.tryConnection.C:
			st.connect()

		case err := <-st.errorChan:
			log.Println("chan error: ", err)

		case msg := <-st.mailBox:
			st.receiverEvent(msg)

		case <-st.signalConnect:
			st.tryConnection.Stop()

			go st.ReadLoop()
			st.Broadcast(&model.StreamEvent{
				ID:     st.ID,
				Status: "online",
				Event:  EVENT_NOTIFICATION,
				Role:   "desktop",
			})
			st.emitNotification()
		}
	}
}
func (st *StreamSrv) emitNotification() {
	dir, _ := os.Getwd()
	fullpath := filepath.Join(dir, "logo-unamad.png")
	nt := toast.Notification{
		AppID:    "UNAMAD",
		Title:    "Notification",
		Icon:     fullpath,
		Message:  "Se conecto al servicio de Comandos UNAMAD",
		Duration: toast.Long,
	}

	nt.Push()
}

func (st *StreamSrv) receiverEvent(value []byte) {
	var event model.StreamEvent
	json.Unmarshal(value, &event)
	st.Broadcast(&event)
}

func (st *StreamSrv) Broadcast(msg *model.StreamEvent) {
	switch msg.Event {
	case EVENT_LOAD_INFO:
		st.loadInfoSystem(msg)
	case EVENT_NOTIFICATION:
		st.notify(msg)
	}
}

func (st *StreamSrv) loadInfoSystem(msg *model.StreamEvent) {
	cmd := entity.NewCommand()
	pcSys := entity.NewComputerSystem(cmd)
	cpuSys := entity.NewCPUSystem(cmd)
	disk := entity.NewPhysicalDisk(cmd)
	bios := entity.NewBios(cmd)

	entity.NewSystemDescriptor().Run(cpuSys, pcSys, disk, bios)

	msg.Payload.Hostname = string(pcSys.Hostname)
	msg.Payload.Manufacturer = string(pcSys.Manufacturer)
	msg.Payload.Model = string(pcSys.Model)
	msg.Payload.System = string(pcSys.System)
	msg.Payload.TotalPhysicalMemory = float64(pcSys.TotalPhysicalMemory)

	msg.Payload.Name = string(cpuSys.Name)
	msg.Payload.Core = string(cpuSys.Cores)
	msg.Payload.LogicalProcessor = string(cpuSys.LogicalProcessors)

	msg.Payload.MediaType = string(disk.MediaType)
	msg.Payload.Size = string(disk.Size)

	msg.Payload.SerialNumber = string(bios.SerialNumber)

	msg.Event = EVENT_LOADED
	msg.Role = "desktop"

	buf, err := json.Marshal(msg)
	if err != nil {
		log.Println(err)
		return
	}
	st.ws.Write(buf)
}
func (st *StreamSrv) notify(msg *model.StreamEvent) {
	buf, err := json.Marshal(msg)
	if err != nil {
		log.Println(err)
		return
	}
	st.ws.Write(buf)
}

const ErrorForciblyClosed = "wsarecv: Se ha forzado la interrupción de una conexión existente por el host remoto."

func (st *StreamSrv) ReadLoop() {
	buf := make([]byte, 1024)
	for {
		n, err := st.ws.Read(buf)
		if err != nil {
			if opErr, _ := err.(*net.OpError); opErr.Err.Error() == ErrorForciblyClosed {
				st.tryConnection.Reset(time.Second * 7)
				break
			}
			if err == io.EOF {
				break
			}
			continue
		}
		msg := buf[:n]
		st.mailBox <- msg
	}
}

func (st *StreamSrv) connect() {
	ws, err := websocket.Dial(fmt.Sprintf(st.urlWS, st.ID), "", st.origin)
	if err != nil {
		st.errorChan <- err
		return
	}
	st.ws = ws
	st.signalConnect <- struct{}{}
}
