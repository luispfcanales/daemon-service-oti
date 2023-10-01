package entity

import (
	"strings"
	"sync"
)

type IPV4_TYPE string

const (
	CMD_IP = "ipconfig | findstr /i ipv4"
)

type IpHost struct {
	Ipv4 IPV4_TYPE `json:"ipv4,omitempty"`

	mailBox     chan interface{}
	stop_actor  chan struct{}
	end_process chan struct{}
	workers     int
	cmd         *Command
}

func NewIpHostSystem(cmd *Command) *IpHost {
	actor := &IpHost{
		cmd:         cmd,
		mailBox:     make(chan any),
		stop_actor:  make(chan struct{}),
		end_process: make(chan struct{}),
	}
	go actor.start()
	return actor
}

func (a *IpHost) receiver(msg any) {
	switch v := msg.(type) {
	case IPV4_TYPE:
		a.Ipv4 = v
		a.end_process <- struct{}{}
	}
}

func (a *IpHost) start() {
loop:
	for {
		select {
		case msg := <-a.mailBox:
			a.receiver(msg)
		case <-a.stop_actor:
			break loop
		}
	}
}
func (a *IpHost) stop() {
	for i := 0; i < a.workers; i++ {
		<-a.end_process
	}
	a.stop_actor <- struct{}{}
}

func (a *IpHost) WorkerLoadInfo(wg ...*sync.WaitGroup) {
	workeds_size := len(wg)

	for _, v := range wg {
		defer v.Done()
	}

	a.workers = workeds_size

	if workeds_size == 0 {
		a.workers = 1
	}

	go a.load()

	a.stop()
}

func (a *IpHost) load() {
	b := a.cmd.MustExecCmd(CMD_IP)
	arr := strings.Split(string(b), ":")
	v := strings.Trim(arr[1], "\n")
	v = strings.TrimSpace(v)
	a.mailBox <- IPV4_TYPE(v)
}
