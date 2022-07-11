package signer

import "fmt"

const Key = -15
const Type = "signal"

type AppSignal struct {
	ch chan bool
}

type Signal interface {
	String() string
	GetKey() int
	Notify()
	C() <-chan bool
}

// New ... create a signal
func New() *AppSignal {
	return &AppSignal{
		ch: make(chan bool, 1),
	}
}

// Notify ... notify chan event
func (r *AppSignal) Notify() {
	r.ch <- true
}

// C ... handle chan
func (r *AppSignal) C() <-chan bool {
	return r.ch
}

// GetKey ... get signal key
func (r *AppSignal) GetKey() int {
	return Key
}

// String ... to keyword
func (r *AppSignal) String() string {
	return fmt.Sprintf("%s:%d", Type, Key)
}
