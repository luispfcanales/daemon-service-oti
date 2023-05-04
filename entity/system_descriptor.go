package entity

import "sync"

type SystemDescriptor struct{}

func NewSystemDescriptor() *SystemDescriptor {
	return &SystemDescriptor{}
}

func (sd *SystemDescriptor) Run(loaders ...Loader) {
	var wg sync.WaitGroup

	sizeWorkers := len(loaders)
	wg.Add(sizeWorkers)

	for _, load := range loaders {
		go load.WorkerLoadInfo(&wg)
	}
	wg.Wait()

	//for _, load := range loaders {
	//	load.PrintInfo()
	//}
}
