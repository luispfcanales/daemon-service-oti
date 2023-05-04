package entity

import (
	"fmt"
	"log"
	"strconv"
	"strings"
	"sync"
)

// custom types
type (
	PD_MEDIA_TYPE string
	PD_SIZE_TYPE  string
)

const (
	PD_SIZE      string = "size"
	PD_MEDIATYPE string = "mediatype"
)

// GB is size in byte
const GB = 1073741824

type PhysicalDisk struct {
	MediaType PD_MEDIA_TYPE `json:"media_type,omitempty"`
	Size      PD_SIZE_TYPE  `json:"size,omitempty"`

	mailBox         chan interface{}
	STOP_ACTOR      chan struct{}
	PROCESS_END     chan struct{}
	process_workers int
	exc             Executor
}

// NewPhysicalDisk return instance of PhysicalDisk
func NewPhysicalDisk(executor Executor) *PhysicalDisk {
	return &PhysicalDisk{
		exc:         executor,
		mailBox:     make(chan interface{}, 10),
		STOP_ACTOR:  make(chan struct{}, 1),
		PROCESS_END: make(chan struct{}),
	}
}
func (p *PhysicalDisk) receiver(message interface{}) {
	switch valueType := message.(type) {
	case PD_SIZE_TYPE:
		p.Size = valueType
		p.PROCESS_END <- struct{}{}
	case PD_MEDIA_TYPE:
		p.MediaType = valueType
		p.PROCESS_END <- struct{}{}
	}
}

func (p *PhysicalDisk) start() {
actorLoop:
	for {
		select {
		case m := <-p.mailBox:
			p.receiver(m)
		case <-p.STOP_ACTOR:
			break actorLoop
		}
	}
}
func (p *PhysicalDisk) stop() {
	var counter int
loop:
	for {
		select {
		case <-p.PROCESS_END:
			counter++
			if p.process_workers == counter {
				p.STOP_ACTOR <- struct{}{}
				break loop
			}
		}
	}
}

func (p *PhysicalDisk) WorkerLoadInfo(wg *sync.WaitGroup) {
	defer wg.Done()

	loaders := []string{PD_MEDIATYPE, PD_SIZE}
	p.process_workers = len(loaders)

	go p.start()

	for _, value := range loaders {
		go p.Get(value)
	}

	p.stop()
}
func (p *PhysicalDisk) Prueba() {
	log.Println(p.getMediaTypeWithPSHELL())
}

// Get function send value typed to channel mailBox
func (p *PhysicalDisk) Get(option string) {
	switch option {
	case PD_SIZE:
		query_cmd := fmt.Sprintf("wmic diskdrive get %s /FORMAT:csv", option)
		values := p.exc.GetInfoCMD(query_cmd)

		s, _ := strconv.Atoi(values[1])
		sizeConvert := s / GB

		p.mailBox <- PD_SIZE_TYPE(fmt.Sprintf("%d GB", sizeConvert))
	case PD_MEDIATYPE:
		p.mailBox <- PD_MEDIA_TYPE(p.getMediaTypeWithPSHELL())
	}
}

func (p *PhysicalDisk) getMediaTypeWithPSHELL() string {
	body := p.exc.GetInfoPOWERSHELL("get-physicaldisk")
	rows := strings.SplitAfter(body, "\n")
	line := strings.Split(rows[3], "\n")
	value := strings.Fields(line[0])
	return value[4]
}
func (p *PhysicalDisk) PrintInfo() {
	fmt.Printf("MediaType\t: %s\n", p.MediaType)
	fmt.Printf("Size\t\t: %s\n", p.Size)
}
