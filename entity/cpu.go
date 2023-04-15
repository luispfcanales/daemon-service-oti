package entity

import (
	"fmt"
	"sync"
)

type CPUSystem struct {
	Name              string `json:"name,omitempty"`
	Cores             string `json:"cores,omitempty"`
	LogicalProcessors string `json:"logical_processors,omitempty"`
	exc               Executor
}

func NewCPUSystem(executor Executor) *CPUSystem {
	return &CPUSystem{
		exc: executor,
	}
}

func (cs *CPUSystem) WorkerLoadInfo(wg *sync.WaitGroup) {
	values := cs.exc.GetInfoCMD("wmic cpu get name,NumberOfCores,NumberOfLogicalProcessors /FORMAT:csv")
	cs.Name = values[1]
	cs.Cores = values[2]
	cs.LogicalProcessors = values[3]
	wg.Done()
}
func (cs *CPUSystem) PrintInfo() {
	fmt.Printf("Chip\t\t: %s\n", cs.Name)
	fmt.Printf("NumberCore\t: %s\n", cs.Cores)
	fmt.Printf("LogicalCore\t: %s\n", cs.LogicalProcessors)
}
