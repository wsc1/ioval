package ioval

import (
	"time"

	"zikichombo.org/sound/freq"
)

type Bus interface {
	SampleRate() freq.T
	SampleTime() int64
	Client(name string) (Client, error)
	Clients() []string
}

func NewBus() Bus {
	return nil
}

type Client interface {
	SampleTime() int64
	StartTime() time.Time
	Get() float64
}

type Message struct {
	Recipient   string
	Float64     float64
	StartSample int64
	EndSample   int64
}
