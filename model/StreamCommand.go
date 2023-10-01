package model

type InformationSystem struct {
	Hostname            string  `json:"hostname,omitempty"`
	Manufacturer        string  `json:"manufacturer,omitempty"`
	Model               string  `json:"model,omitempty"`
	System              string  `json:"system,omitempty"`
	TotalPhysicalMemory float64 `json:"total_physical_memory,omitempty"`
	Name                string  `json:"name,omitempty"`
	Core                string  `json:"core,omitempty"`
	LogicalProcessor    string  `json:"logical_processor,omitempty"`
	MediaType           string  `json:"media_type,omitempty"`
	Size                string  `json:"size,omitempty"`
	SerialNumber        string  `json:"serial_number,omitempty"`
}

type StreamEvent struct {
	ID            string            `json:"id,omitempty"`
	Status        string            `json:"status,omitempty"`
	Event         uint8             `json:"event,omitempty"`
	Role          string            `json:"role,omitempty"`
	EventEmisorID string            `json:"event_emisor_id,omitempty"`
	Payload       InformationSystem `json:"payload,omitempty"`
}
