package entity

import (
	"fmt"
	"log"
	"math"
	"strconv"
	"sync"
)

type ComputerSystem struct {
	Hostname            string  `json:"hostname,omitempty"`
	Manufacturer        string  `json:"manufacturer,omitempty"`
	Model               string  `json:"model,omitempty"`
	SystemType          string  `json:"system_type,omitempty"`
	TotalPhysicalMemory float64 `json:"total_physical_memory,omitempty"`
	exc                 Executor
}

const MB int = 1073741824

func NewComputerSystem(executor Executor) *ComputerSystem {
	return &ComputerSystem{
		exc: executor,
	}
}

func (pc *ComputerSystem) WorkerLoadInfo(wg *sync.WaitGroup) {
	values := pc.exc.GetInfoCMD("wmic computersystem get model,Manufacturer,systemtype,TotalPhysicalMemory /FORMAT:csv")
	sizeMB, err := strconv.ParseFloat(values[4], 32)
	if err != nil {
		log.Println(err)
		return
	}
	sizeMB = sizeMB / float64(MB)
	sizeMB = math.Round(sizeMB)

	pc.Hostname = values[0]
	pc.Manufacturer = values[1]
	pc.Model = values[2]
	pc.SystemType = values[3]
	pc.TotalPhysicalMemory = sizeMB
	wg.Done()
}
func (pc *ComputerSystem) PrintInfo() {
	fmt.Printf("Hostname\t: %s\n", pc.Hostname)
	fmt.Printf("Manufacturer\t: %s\n", pc.Manufacturer)
	fmt.Printf("Model\t\t: %s\n", pc.Model)
	fmt.Printf("Architecture\t: %s\n", pc.SystemType)
	fmt.Printf("RamGB\t\t: %f\n", pc.TotalPhysicalMemory)
}
