package main

import (
	"bytes"
	"github.com/hajimehoshi/ebiten/v2"
//	"github.com/hajimehoshi/ebiten/v2/text"
	"golang.org/x/image/font"
	"golang.org/x/image/font/opentype"
	"image"
	_ "image/png"
	"log"
	"math/rand"
	"time"
	"github.com/lvajxi03/trffq/assets"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	tt, _ := opentype.Parse(assets.Font_Fredoka)
	Face_Fredoka, _ = opentype.NewFace(tt, &opentype.FaceOptions{
		Size: 72,
		DPI: 72,
		Hinting: font.HintingFull,
	})
	Face_FredokaLarge, _ = opentype.NewFace(tt, &opentype.FaceOptions{
		Size: 160,
		DPI: 72,
		Hinting: font.HintingFull,
	})
	ebiten.SetWindowSize(ARENA_WIDTH, ARENA_HEIGHT)
	ebiten.SetWindowTitle(APPLICATION_TITLE)
	ebiten.SetWindowDecorated(false)
	ebiten.SetFullscreen(true)

	img, _, _ := image.Decode(bytes.NewReader(assets.Embedded_Background_Default))
	assets.Background_Default = ebiten.NewImageFromImage(img)
	img, _, _ = image.Decode(bytes.NewReader(assets.Flag_EN))
	img_flag_en = ebiten.NewImageFromImage(img)
	img, _, _ = image.Decode(bytes.NewReader(assets.Flag_PL))
	img_flag_pl = ebiten.NewImageFromImage(img)
	traffiq := NewTraffiq()
	traffiq.AddTimer("timer-loading", TimerLoading)
	traffiq.StartTimer("timer-loading", 0.02)
	if err := ebiten.RunGame(traffiq); err != nil {
		log.Fatal(err)
	}
}
