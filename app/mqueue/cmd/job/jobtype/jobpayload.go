package jobtype

import "time"

// TestQueuePayload  test queue payload
type TestQueuePayload struct {
	Name string
	Time time.Time
}

// DeferCloseHomestayOrderPayload defer close homestay order
type DeferCloseHomestayOrderPayload struct {
	Sn string
}
