package main

import (
	"github.com/luispfcanales/daemon-service-oti/entity"
)

func main() {
	c := entity.NewCommand()

	compSystem := entity.NewComputerSystem(c)
	cpuSys := entity.NewCPUSystem(c)
	disk := entity.NewPhysicalDisk(c)

	entity.NewSystemDescriptor().Run(compSystem, cpuSys, disk)
}
