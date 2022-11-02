package main

import (
	"math/rand"
	// "github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"image/color"
)

func TimerLoading(traffiq *Traffiq) {
	if traffiq != nil {
		if traffiq.WelcomeCounter > 30 {
						traffiq.StopTimer("timer-loading")
						traffiq.ChangeBoard(BOARD_INTRO)
		} else {
			for a := 0; a < ARENA_HEIGHT/40; a++ {
				Red := uint8(rand.Intn(255))
				Green := uint8(rand.Intn(255))
				Blue := uint8(rand.Intn(255))
				c := color.RGBA{R: Red, G: Green, B: Blue, A: 255}
				ebitenutil.DrawRect(traffiq.Canvas, 0, float64(a*40), ARENA_WIDTH, 40, c)
			}
			traffiq.WelcomeCounter++
		}
	}
}

func TimerIntro(traffiq *Traffiq) {
	if traffiq != nil {
		
	}
}
