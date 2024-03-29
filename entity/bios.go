package entity

import (
	"fmt"
	"sync"
)

const (
	BIOS_SERIAL = "serialnumber"
)

type SERIAL_NUMBER_TYPE string

type Bios struct {
	SerialNumber SERIAL_NUMBER_TYPE `json:"serial_number,omitempty"`

	mailBox         chan interface{}
	stop_actor      chan struct{}
	process_end     chan struct{}
	process_workers int
	exc             Executor
}

func NewBios(executor Executor) *Bios {
	return &Bios{
		exc:         executor,
		mailBox:     make(chan interface{}, 10),
		stop_actor:  make(chan struct{}, 1),
		process_end: make(chan struct{}),
	}
}
func (b *Bios) receiver(message interface{}) {
	switch valueType := message.(type) {
	case SERIAL_NUMBER_TYPE:
		b.SerialNumber = valueType
		b.process_end <- struct{}{}
	}
}

func (p *Bios) start() {
actorLoop:
	for {
		select {
		case m := <-p.mailBox:
			p.receiver(m)
		case <-p.stop_actor:
			break actorLoop
		}
	}
}
func (b *Bios) stop() {
	var counter int
loop:
	for {
		select {
		case <-b.process_end:
			counter++
			if b.process_workers == counter {
				b.stop_actor <- struct{}{}
				break loop
			}
		}
	}
}
func (b *Bios) WorkerLoadInfo(wg *sync.WaitGroup) {
	defer wg.Done()

	loaders := []string{BIOS_SERIAL}
	b.process_workers = len(loaders)

	go b.start()

	for _, value := range loaders {
		go b.Get(value)
	}

	b.stop()
}
func (b *Bios) Get(option string) {
	switch option {
	case BIOS_SERIAL:
		b.mailBox <- SERIAL_NUMBER_TYPE(b.getSerialNumber())
	}
}

func (b *Bios) getSerialNumber() string {
	queryCMD := fmt.Sprintf("wmic bios get %s /format:csv", BIOS_SERIAL)
	values := b.exc.GetInfoCMD(queryCMD)
	return values[1]
}
func (b *Bios) PrintInfo() {
	fmt.Printf("Serial\t\t: %s\n", b.SerialNumber)
}
