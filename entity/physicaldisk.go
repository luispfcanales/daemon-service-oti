package entity

import (
	"fmt"
	"strings"
	"sync"
)

type PhysicalDisk struct {
	MediaType string `json:"media_type,omitempty"`
	Size      string `json:"size,omitempty"`
	exc       Executor
}

func NewPhysicalDisk(executor Executor) *PhysicalDisk {
	return &PhysicalDisk{
		exc: executor,
	}
}

func (p *PhysicalDisk) WorkerLoadInfo(wg *sync.WaitGroup) {
	body := p.exc.GetInfoPOWERSHELL("get-physicaldisk")
	rows := strings.SplitAfter(body, "\n")
	line := strings.Split(rows[3], "\n")
	value := strings.Fields(line[0])

	p.MediaType = value[4]
	p.Size = value[9] + value[10]
	wg.Done()
}
func (p *PhysicalDisk) PrintInfo() {
	fmt.Printf("MediaType\t: %s\n", p.MediaType)
	fmt.Printf("Size\t\t: %s\n", p.Size)
}
