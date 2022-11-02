package graphics

import (
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"image"
	"image/color"	
	"github.com/hajimehoshi/ebiten/v2"	
	"github.com/lvajxi03/primi"
)

func FontBounds2Frame(rect image.Rectangle, newx, newy, margin int) (x1, y1, x2, y2 float64) {
	w := rect.Max.X - rect.Min.X + margin
	h := rect.Max.Y - rect.Min.Y + margin
	y := newy - h
	return float64(newx), float64(y), float64(newx + w), float64(y + h)
}

func FontBounds2Rect(rect image.Rectangle, newx, newy, margin int) *primi.Rect {
	w := float64(rect.Max.X - rect.Min.X + margin)
	h := float64(rect.Max.Y - rect.Min.Y + margin)
	y := float64(newy) - h
	r := &primi.Rect{float64(newx), y, w, h}
	return r
}

func DrawFrame(surface *ebiten.Image, x1, y1, x2, y2 float64, clr color.Color) {
	if surface != nil {
		ebitenutil.DrawLine(surface, x1, y1, x2, y1, clr)
		ebitenutil.DrawLine(surface, x2, y1, x2, y2, clr)
		ebitenutil.DrawLine(surface, x2, y2, x1, y2, clr)
		ebitenutil.DrawLine(surface, x1, y2, x1, y1, clr)
	}
}

func DrawFrameRect(surface *ebiten.Image, rect *primi.Rect, clr color.Color) {
	if surface != nil && rect != nil {
		x2 := rect.X + rect.W
		y2 := rect.Y + rect.H
		ebitenutil.DrawLine(surface, rect.X, rect.Y, x2, rect.Y, clr)
		ebitenutil.DrawLine(surface, x2, rect.Y, x2, y2, clr)
		ebitenutil.DrawLine(surface, x2, y2, rect.X, y2, clr)
		ebitenutil.DrawLine(surface, rect.X, y2, rect.X, rect.Y, clr)
	}
}
