package entity

import (
	"fmt"
	"log"
	"math"
	"strconv"
	"sync"
)

type (
	HostnameType            string
	ManufacturerType        string
	ModelType               string
	SystemType              string
	TotalPhysicalMemoryType float64
)

type ComputerSystem struct {
	Hostname             HostnameType            `json:"hostname,omitempty"`
	Manufacturer         ManufacturerType        `json:"manufacturer,omitempty"`
	Model                ModelType               `json:"model,omitempty"`
	System               SystemType              `json:"system,omitempty"`
	TotalPhysicalMemory  TotalPhysicalMemoryType `json:"total_physical_memory,omitempty"`
	infoSystemChan       chan interface{}
	signalChan           chan struct{}
	signalNumberToFinish int
	exc                  Executor
}

const MB int = 1073741824

func NewComputerSystem(executor Executor) *ComputerSystem {
	return &ComputerSystem{
		infoSystemChan: make(chan interface{}),
		signalChan:     make(chan struct{}, 1),
		exc:            executor,
	}
}

func (pc *ComputerSystem) WorkerLoadInfo(wg *sync.WaitGroup) {
	defer wg.Done()
	loaders := []string{"Manufacturer", "model", "systemtype", "TotalPhysicalMemory"}

	sizeLoaders := len(loaders)

	for _, value := range loaders {
		go pc.GetInfo(value, pc.infoSystemChan)
	}

workerSystem:
	for {
		select {
		case <-pc.signalChan:
			break workerSystem
		case valueType := <-pc.infoSystemChan:
			switch value := valueType.(type) {
			case ManufacturerType:
				pc.Manufacturer = value
				pc.sendSignalFinish(sizeLoaders)
			case ModelType:
				pc.Model = value
				pc.sendSignalFinish(sizeLoaders)
			case SystemType:
				pc.System = value
				pc.sendSignalFinish(sizeLoaders)
			case TotalPhysicalMemoryType:
				pc.TotalPhysicalMemory = value
				pc.sendSignalFinish(sizeLoaders)
			}
		}
	}

}
func (pc *ComputerSystem) sendSignalFinish(sizeEnd int) {
	pc.signalNumberToFinish++
	if pc.signalNumberToFinish == sizeEnd {
		pc.signalChan <- struct{}{}
		close(pc.infoSystemChan)
		close(pc.signalChan)
	}
}

func (pc *ComputerSystem) GetInfo(option string, sender chan<- interface{}) {
	queryCMD := fmt.Sprintf("wmic computersystem get %s /FORMAT:csv", option)
	values := pc.exc.GetInfoCMD(queryCMD)

	switch option {
	case "Manufacturer":
		sender <- ManufacturerType(values[1])
	case "model":
		sender <- ModelType(values[1])
	case "systemtype":
		sender <- SystemType(values[1])
	case "TotalPhysicalMemory":
		sender <- TotalPhysicalMemoryType(pc.getSizeMemoryRAM(values[1]))
	}
}

func (pc *ComputerSystem) getSizeMemoryRAM(size string) float64 {
	sizeMB, err := strconv.ParseFloat(size, 32)
	if err != nil {
		log.Println(err)
		return 0
	}
	sizeMB = sizeMB / float64(MB)
	sizeMB = math.Round(sizeMB)
	return sizeMB
}

func (pc *ComputerSystem) PrintInfo() {
	fmt.Printf("Hostname\t: %s\n", pc.Hostname)
	fmt.Printf("Manufacturer\t: %s\n", pc.Manufacturer)
	fmt.Printf("Model\t\t: %s\n", pc.Model)
	fmt.Printf("Architecture\t: %s\n", pc.System)
	fmt.Printf("RamGB\t\t: %f\n", pc.TotalPhysicalMemory)
}
