package main

// traffiq structure: render game into a window and handle user input

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/text"
	"image/color"
"github.com/lvajxi03/trffq/assets"
)

var GameKeys [46]ebiten.Key = [46]ebiten.Key{
	ebiten.KeyA, ebiten.KeyB, ebiten.KeyC, ebiten.KeyD, ebiten.KeyE, ebiten.KeyF, ebiten.KeyG, ebiten.KeyH,
	ebiten.KeyI, ebiten.KeyJ, ebiten.KeyK, ebiten.KeyL, ebiten.KeyM, ebiten.KeyN, ebiten.KeyO, ebiten.KeyP,
	ebiten.KeyQ, ebiten.KeyR, ebiten.KeyS, ebiten.KeyT, ebiten.KeyU, ebiten.KeyV, ebiten.KeyW, ebiten.KeyX,
	ebiten.KeyY, ebiten.KeyZ, ebiten.KeyDigit0, ebiten.KeyDigit1, ebiten.KeyDigit2, ebiten.KeyDigit3,
	ebiten.KeyDigit4, ebiten.KeyDigit5, ebiten.KeyDigit6, ebiten.KeyDigit7, ebiten.KeyDigit8, ebiten.KeyDigit9,
	ebiten.KeySpace, ebiten.KeyTab, ebiten.KeyF1, ebiten.KeyArrowLeft, ebiten.KeyArrowRight, ebiten.KeyArrowUp,
	ebiten.KeyArrowDown, ebiten.KeyEnter, ebiten.KeyEscape}

func (traffiq *Traffiq) Draw(screen *ebiten.Image) {
	if traffiq != nil {
		switch traffiq.Board {
		case BOARD_LOADING:
			traffiq.PaintLoading(screen)
		case BOARD_INTRO:
			traffiq.PaintIntro(screen)
		case BOARD_MENU:
			traffiq.PaintMenu(screen)
		case BOARD_LOADGAME:
			traffiq.PaintLoad(screen)
		case BOARD_GAME:
			traffiq.PaintGame(screen)
		case BOARD_OPTIONS:
			traffiq.PaintOptions(screen)
		case BOARD_HISCORES:
			traffiq.PaintHiscores(screen)
		case BOARD_SETTINGS:
			traffiq.PaintSettings(screen)
		case BOARD_NEWSCORE:
			traffiq.PaintNewscore(screen)
		case BOARD_HELP:
			traffiq.PaintHelp(screen)
		case BOARD_ABOUT:
			traffiq.PaintAbout(screen)
		}
	}
}

func (traffiq *Traffiq) PaintLoad (screen *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}
	traffiq.Canvas.DrawImage(assets.Background_Default, op)
	c := color.RGBA{R: 0, G: 0, B: 0, A: 128}
	ebitenutil.DrawRect(traffiq.Canvas, 0, 960, ARENA_WIDTH, 120, c)
	
	text.Draw(traffiq.Canvas, titles[traffiq.Lang][traffiq.Board], Face_FredokaLarge, int(traffiq.Menu.HeaderLeft + SHADOW_MARGIN), 165, colors["default-font-shadow"])
	text.Draw(traffiq.Canvas, titles[traffiq.Lang][traffiq.Board], Face_FredokaLarge, int(traffiq.Menu.HeaderLeft), 160, colors["default-font-inactive"])
	op.GeoM.Translate(float64(ARENA_WIDTH - 160), float64(ARENA_HEIGHT - 200))
	traffiq.Canvas.DrawImage(img_flag_en, op)
	op.GeoM.Translate(float64(80), 0)
	traffiq.Canvas.DrawImage(img_flag_pl, op)
	op = &ebiten.DrawImageOptions{}
	screen.DrawImage(traffiq.Canvas, op)
}

func (traffiq *Traffiq) Layout(w, h int) (screenWidth, screenHeight int) {
	return ARENA_WIDTH, ARENA_HEIGHT
}

func (traffiq *Traffiq) PaintLoading(screen *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}
	screen.DrawImage(traffiq.Canvas, op)
}

func (traffiq *Traffiq) PaintIntro(screen *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}
	screen.DrawImage(traffiq.Canvas, op)
}

func (traffiq *Traffiq) PaintMenu(screen *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}
	traffiq.Canvas.DrawImage(assets.Background_Default, op)
	c := color.RGBA{R: 0, G: 0, B: 0, A: 128}
	ebitenutil.DrawRect(traffiq.Canvas, 0, 960, ARENA_WIDTH, 120, c)
	
	text.Draw(traffiq.Canvas, "TRAFFIQ!", Face_FredokaLarge, int(traffiq.Menu.HeaderLeft + SHADOW_MARGIN), 165, colors["default-font-shadow"])
	text.Draw(traffiq.Canvas, "TRAFFIQ!", Face_FredokaLarge, int(traffiq.Menu.HeaderLeft), 160, colors["default-font-inactive"])
	i := 0
	for index, value := range menustr[traffiq.Lang] {
		text.Draw(traffiq.Canvas, value, Face_Fredoka, 355, 355 + i, colors["default-font-shadow"])
		if index == traffiq.Menu.Option {
			text.Draw(traffiq.Canvas, value, Face_Fredoka, 350, 350 + i, colors["default-font-active"])
		} else {
			text.Draw(traffiq.Canvas, value, Face_Fredoka, 350, 350 + i, colors["default-font-inactive"])
		}
		
		
		i += 80
	}
	op.GeoM.Translate(float64(ARENA_WIDTH - 160), float64(ARENA_HEIGHT - 200))
	traffiq.Canvas.DrawImage(img_flag_en, op)
	op.GeoM.Translate(float64(80), 0)
	traffiq.Canvas.DrawImage(img_flag_pl, op)
	op = &ebiten.DrawImageOptions{}
	screen.DrawImage(traffiq.Canvas, op)
}

func (traffiq *Traffiq) PaintGame(screen *ebiten.Image) {
	if traffiq != nil {
		switch traffiq.Mode {
		case MODE_INIT:
			traffiq.PaintInit(screen)
		case MODE_PLAY:
			traffiq.PaintPlay(screen)
		case MODE_PAUSED:
			traffiq.PaintPaused(screen)
		case MODE_KILLED:
			traffiq.PaintKilled(screen)
		case MODE_GAMEOVER:
			traffiq.PaintGameover(screen)
		}
	}
}

func (traffiq *Traffiq) PaintHiscores(screen *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}
	traffiq.Canvas.DrawImage(assets.Background_Default, op)
	c := color.RGBA{R: 0, G: 0, B: 0, A: 128}
	ebitenutil.DrawRect(traffiq.Canvas, 0, 960, ARENA_WIDTH, 120, c)
	
	text.Draw(traffiq.Canvas, titles[traffiq.Lang][traffiq.Board], Face_FredokaLarge, int(traffiq.Menu.HeaderLeft + SHADOW_MARGIN), 165, colors["default-font-shadow"])
	text.Draw(traffiq.Canvas, titles[traffiq.Lang][traffiq.Board], Face_FredokaLarge, int(traffiq.Menu.HeaderLeft), 160, colors["default-font-inactive"])
	op.GeoM.Translate(float64(ARENA_WIDTH - 160), float64(ARENA_HEIGHT - 200))
	traffiq.Canvas.DrawImage(img_flag_en, op)
	op.GeoM.Translate(float64(80), 0)
	traffiq.Canvas.DrawImage(img_flag_pl, op)
	op = &ebiten.DrawImageOptions{}
	screen.DrawImage(traffiq.Canvas, op)
}

func (traffiq *Traffiq) PaintNewscore(screen *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}
	traffiq.Canvas.DrawImage(assets.Background_Default, op)
	c := color.RGBA{R: 0, G: 0, B: 0, A: 128}
	ebitenutil.DrawRect(traffiq.Canvas, 0, 960, ARENA_WIDTH, 120, c)
	
	text.Draw(traffiq.Canvas, titles[traffiq.Lang][traffiq.Board], Face_FredokaLarge, int(traffiq.Menu.HeaderLeft + SHADOW_MARGIN), 165, colors["default-font-shadow"])
	text.Draw(traffiq.Canvas, titles[traffiq.Lang][traffiq.Board], Face_FredokaLarge, int(traffiq.Menu.HeaderLeft), 160, colors["default-font-inactive"])
	op.GeoM.Translate(float64(ARENA_WIDTH - 160), float64(ARENA_HEIGHT - 200))
	traffiq.Canvas.DrawImage(img_flag_en, op)
	op.GeoM.Translate(float64(80), 0)
	traffiq.Canvas.DrawImage(img_flag_pl, op)
	op = &ebiten.DrawImageOptions{}
	screen.DrawImage(traffiq.Canvas, op)
}

func (traffiq *Traffiq) PaintHelp(screen *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}
	traffiq.Canvas.DrawImage(assets.Background_Default, op)
	c := color.RGBA{R: 0, G: 0, B: 0, A: 128}
	ebitenutil.DrawRect(traffiq.Canvas, 0, 960, ARENA_WIDTH, 120, c)
	
	text.Draw(traffiq.Canvas, titles[traffiq.Lang][traffiq.Board], Face_FredokaLarge, int(traffiq.Menu.HeaderLeft + SHADOW_MARGIN), 165, colors["default-font-shadow"])
	text.Draw(traffiq.Canvas, titles[traffiq.Lang][traffiq.Board], Face_FredokaLarge, int(traffiq.Menu.HeaderLeft), 160, colors["default-font-inactive"])
	op.GeoM.Translate(float64(ARENA_WIDTH - 160), float64(ARENA_HEIGHT - 200))
	traffiq.Canvas.DrawImage(img_flag_en, op)
	op.GeoM.Translate(float64(80), 0)
	traffiq.Canvas.DrawImage(img_flag_pl, op)
	op = &ebiten.DrawImageOptions{}
	screen.DrawImage(traffiq.Canvas, op)
}

func (traffiq *Traffiq) PaintSettings(screen *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}
	traffiq.Canvas.DrawImage(assets.Background_Default, op)
	c := color.RGBA{R: 0, G: 0, B: 0, A: 128}
	ebitenutil.DrawRect(traffiq.Canvas, 0, 960, ARENA_WIDTH, 120, c)
	
	text.Draw(traffiq.Canvas, titles[traffiq.Lang][traffiq.Board], Face_FredokaLarge, int(traffiq.Menu.HeaderLeft + SHADOW_MARGIN), 165, colors["default-font-shadow"])
	text.Draw(traffiq.Canvas, titles[traffiq.Lang][traffiq.Board], Face_FredokaLarge, int(traffiq.Menu.HeaderLeft), 160, colors["default-font-inactive"])
	op.GeoM.Translate(float64(ARENA_WIDTH - 160), float64(ARENA_HEIGHT - 200))
	traffiq.Canvas.DrawImage(img_flag_en, op)
	op.GeoM.Translate(float64(80), 0)
	traffiq.Canvas.DrawImage(img_flag_pl, op)
	op = &ebiten.DrawImageOptions{}
	screen.DrawImage(traffiq.Canvas, op)
}
func (traffiq *Traffiq) PaintAbout(screen *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}
	traffiq.Canvas.DrawImage(assets.Background_Default, op)
	c := color.RGBA{R: 0, G: 0, B: 0, A: 128}
	ebitenutil.DrawRect(traffiq.Canvas, 0, 960, ARENA_WIDTH, 120, c)
	
	text.Draw(traffiq.Canvas, titles[traffiq.Lang][traffiq.Board], Face_FredokaLarge, int(traffiq.Menu.HeaderLeft + SHADOW_MARGIN), 165, colors["default-font-shadow"])
	text.Draw(traffiq.Canvas, titles[traffiq.Lang][traffiq.Board], Face_FredokaLarge, int(traffiq.Menu.HeaderLeft), 160, colors["default-font-inactive"])
	op.GeoM.Translate(float64(ARENA_WIDTH - 160), float64(ARENA_HEIGHT - 200))
	traffiq.Canvas.DrawImage(img_flag_en, op)
	op.GeoM.Translate(float64(80), 0)
	traffiq.Canvas.DrawImage(img_flag_pl, op)
	op = &ebiten.DrawImageOptions{}
	screen.DrawImage(traffiq.Canvas, op)
}

func (traffiq *Traffiq) PaintInit(screen *ebiten.Image) {

}

func (traffiq *Traffiq) PaintPlay(screen *ebiten.Image) {

}

func (traffiq *Traffiq) PaintPaused(screen *ebiten.Image) {

}

func (traffiq *Traffiq) PaintKilled(screen *ebiten.Image) {

}

func (traffiq *Traffiq) PaintGameover(screen *ebiten.Image) {

}

func (traffiq *Traffiq) PaintOptions(screen *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}
	traffiq.Canvas.DrawImage(assets.Background_Default, op)
	c := color.RGBA{R: 0, G: 0, B: 0, A: 128}
	ebitenutil.DrawRect(traffiq.Canvas, 0, 960, ARENA_WIDTH, 120, c)
	
	text.Draw(traffiq.Canvas, titles[traffiq.Lang][traffiq.Board], Face_FredokaLarge, int(traffiq.Menu.HeaderLeft + SHADOW_MARGIN), 165, colors["default-font-shadow"])
	text.Draw(traffiq.Canvas, titles[traffiq.Lang][traffiq.Board], Face_FredokaLarge, int(traffiq.Menu.HeaderLeft), 160, colors["default-font-inactive"])
	op.GeoM.Translate(float64(ARENA_WIDTH - 160), float64(ARENA_HEIGHT - 200))
	traffiq.Canvas.DrawImage(img_flag_en, op)
	op.GeoM.Translate(float64(80), 0)
	traffiq.Canvas.DrawImage(img_flag_pl, op)
	op = &ebiten.DrawImageOptions{}
	screen.DrawImage(traffiq.Canvas, op)
}
