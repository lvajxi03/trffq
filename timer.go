package main

import (
	"github.com/hajimehoshi/ebiten/v2"
)

type Timer struct {
	State   bool
	Ticks   uint
	Timeout uint
	Parent  *Traffiq
	Handler func(g *Traffiq)
}

func (t *Timer) Tick() {
	if t.State {
		if t.Ticks < t.Timeout {
			t.Ticks++
		} else {
			t.Ticks = 0
			if t.Handler != nil {
				t.Handler(t.Parent)
			}
		}
	}
}

func (t *Timer) Init(parent *Traffiq, handler func(g *Traffiq)) {
	t.Parent = parent
	t.Handler = handler
	t.State = false

}

func (t *Timer) Start(timeout float64) {
	t.Timeout = uint(float64(ebiten.TPS()) * timeout)
	t.Ticks = 0
	t.State = true
}

func (t *Timer) IsActive() bool {
	return t.State
}

func (t *Timer) Stop() {
	t.Ticks = 0
	t.State = false
}
