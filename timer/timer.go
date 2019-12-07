package timer

import (
	"time"
)

type Timer struct {
	start time.Time
	end   time.Time
}

func NewTimer() *Timer { return &Timer{} }

func (t *Timer) Start() { t.start = time.Now() }

func (t *Timer) Stop() { t.end = time.Now() }

func (t Timer) Nanoseconds() int64 {
	return t.end.Sub(t.start).Nanoseconds()
}

func (t Timer) Milliseconds() int64 {
	duration := t.end.Sub(t.start)
	return duration.Nanoseconds() / 1000000
}
