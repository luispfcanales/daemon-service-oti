package model

type ResponseApi struct {
	Status  int    `json:"status,omitempty"`
	Message string `json:"message,omitempty"`
	Code    string `json:"code,omitempty"`
}
type RequestComputer struct {
	PatrimonialCode string `json:"patrimonial_code,omitempty"`
	Serial          string `json:"serial,omitempty"`
	Name            string `json:"hostname,omitempty"`
	Maker           string `json:"facturer,omitempty"`
	Model           string `json:"model,omitempty"`
	Architecture    string `json:"architecture,omitempty"`
	Ram             string `json:"ram,omitempty"`
	Processor       string `json:"processor,omitempty"`
	Core            string `json:"core,omitempty"`
	LogicalCore     string `json:"logical_core,omitempty"`
	Disk            string `json:"disk,omitempty"`
	SizeDisk        string `json:"size_disk,omitempty"`
}
