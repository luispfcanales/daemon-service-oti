package entity

import (
	"fmt"
	"sync"
)

// declare custom types
type (
	CPU_NAME_TYPE               string
	CPU_NUMBER_CORES_TYPE       string
	CPU_LOGICAL_PROCESSORS_TYPE string
)

const (
	CPU_NAME               = "name"
	CPU_NUMBER_CORES       = "NumberOfCores"
	CPU_LOGICAL_PROCESSORS = "NumberOfLogicalProcessors"
)

// CPUSystem is model to load information of CPU
type CPUSystem struct {
	Name              CPU_NAME_TYPE               `json:"name,omitempty"`
	Cores             CPU_NUMBER_CORES_TYPE       `json:"cores,omitempty"`
	LogicalProcessors CPU_LOGICAL_PROCESSORS_TYPE `json:"logical_processors,omitempty"`

	mailBox        chan interface{} //
	STOP_ACTOR     chan struct{}    // stop actor
	PROCESS_END    chan struct{}    // stop load information process
	numberToFinish int
	exc            Executor
}

// NewCPUSystem return instance of CPUSystem
func NewCPUSystem(executor Executor) *CPUSystem {
	CPU := &CPUSystem{
		exc:         executor,
		mailBox:     make(chan interface{}, 10),
		STOP_ACTOR:  make(chan struct{}, 1),
		PROCESS_END: make(chan struct{}),
	}
	return CPU
}

func (cs *CPUSystem) start() {
ActorLoop:
	for {
		select {
		case valueType := <-cs.mailBox:
			switch value := valueType.(type) {
			case CPU_NAME_TYPE:
				cs.Name = value
				cs.PROCESS_END <- struct{}{}
			case CPU_NUMBER_CORES_TYPE:
				cs.Cores = value
				cs.PROCESS_END <- struct{}{}
			case CPU_LOGICAL_PROCESSORS_TYPE:
				cs.LogicalProcessors = value
				cs.PROCESS_END <- struct{}{}
			}
		case <-cs.STOP_ACTOR:
			break ActorLoop
		}
	}
}

// stop load process and send signal to STOP ACTORLOOP
func (cs *CPUSystem) stop() {
	var counter int
ActorStop:
	for {
		select {
		case <-cs.PROCESS_END:
			counter++
			if counter == cs.numberToFinish {
				cs.STOP_ACTOR <- struct{}{}
				close(cs.PROCESS_END)
				close(cs.STOP_ACTOR)
				break ActorStop
			}
		}
	}
}

func (cs *CPUSystem) WorkerLoadInfo(wg *sync.WaitGroup) {
	defer wg.Done()

	loaders := []string{CPU_NAME, CPU_NUMBER_CORES, CPU_LOGICAL_PROCESSORS}
	cs.numberToFinish = len(loaders)

	go cs.start()

	for _, value := range loaders {
		go cs.GetINfo(value)
	}

	cs.stop()
}

func (cs *CPUSystem) GetINfo(option string) {
	queryCMD := fmt.Sprintf("wmic cpu get %s /FORMAT:csv", option)
	values := cs.exc.GetInfoCMD(queryCMD)

	switch option {
	case CPU_NAME:
		cs.mailBox <- CPU_NAME_TYPE(values[1])
	case CPU_NUMBER_CORES:
		cs.mailBox <- CPU_NUMBER_CORES_TYPE(values[1])
	case CPU_LOGICAL_PROCESSORS:
		cs.mailBox <- CPU_LOGICAL_PROCESSORS_TYPE(values[1])
	}
}

func (cs *CPUSystem) PrintInfo() {
	fmt.Printf("Chip\t\t: %s\n", cs.Name)
	fmt.Printf("NumberCore\t: %s\n", cs.Cores)
	fmt.Printf("LogicalCore\t: %s\n", cs.LogicalProcessors)
}
