package main

import (
	"golang.org/x/image/font"
	"github.com/hajimehoshi/ebiten/v2"
)

var (
	Faces = map[string]font.Face{
		"acknowledge": nil,
	}
	Face_Acknowledge    font.Face
	Face_TopazPlusA1200 font.Face
	Face_GemsbuckBlack  font.Face
	Face_GemsbuckBlackLarge font.Face
	Face_ImperialWeb    font.Face
        Face_ImperialWebLarge font.Face
		Face_ImperialWebLarge2 font.Face
	Face_Moder          font.Face
	Face_Mosoul         font.Face
	Face_Fredoka	font.Face
	Face_FredokaLarge font.Face
)

var (
	img_flag_en *ebiten.Image
	img_flag_pl *ebiten.Image
)
