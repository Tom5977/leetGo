package tool

import "time"

type Timer struct {
	start   time.Time
	stop    time.Time
	last    time.Duration
	counter []time.Duration
}

func (this *Timer) Start() {
	this.start = time.Now()
}

func (this *Timer) Stop() {
	this.stop = time.Now()
	this.last = this.stop.Sub(this.start)
	this.counter = append(this.counter, this.last)
}

func (this Timer) Duration() time.Duration {
	return this.last
}
