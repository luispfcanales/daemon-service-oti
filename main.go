package main

import (
	"github.com/luispfcanales/daemon-service-oti/entity"
)

func main() {
	c := entity.NewCommand()

	compSystem := entity.NewComputerSystem(c) //-> ready to load
	cpuSys := entity.NewCPUSystem(c)          //-> ready to load
	//disk := entity.NewPhysicalDisk(c)

	entity.NewSystemDescriptor().Run(compSystem, cpuSys)
	//entity.NewSystemDescriptor().Run(compSystem, cpuSys, disk)
	//entity.NewSystemDescriptor().Run(cpuSys)
}
