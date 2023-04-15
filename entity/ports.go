package entity

import "sync"

type Executor interface {
	GetInfoCMD(string) []string
	GetInfoPOWERSHELL(string) string
}

type Loader interface {
	WorkerLoadInfo(*sync.WaitGroup)
	PrintInfo()
}
