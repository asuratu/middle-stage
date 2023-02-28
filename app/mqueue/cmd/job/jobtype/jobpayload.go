package jobtype

import "time"

// TestQueuePayload  test queue payload
type TestQueuePayload struct {
	Name string
	Time time.Time
}

// RegisterUserPayload  register userinfo to es
type RegisterUserPayload struct {
	Name         string `json:"name"`
	City         string `json:"city"`
	Introduction string `json:"introduction"`
}

// DeferCloseHomestayOrderPayload defer close homestay order
type DeferCloseHomestayOrderPayload struct {
	Sn string
}
