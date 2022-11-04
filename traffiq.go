package main

// Game structure and related dependencies: game logic

import (
	"errors"
	"fmt"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

type Traffiq struct {
	Timers         map[string]*Timer
	Board          Board
	Mode           Mode
	Message        string
	IntroCounter   int
	WelcomeCounter int
	Canvas         *ebiten.Image
	Options        *ebiten.DrawImageOptions
	Menu           *Menu
	Lang           int
}

func NewTraffiq() *Traffiq {
	t := &Traffiq{
		Timers:         map[string]*Timer{},
		Board:          BOARD_LOADING,
		Mode:           MODE_INIT,
		WelcomeCounter: 0,
		Canvas:         ebiten.NewImage(ARENA_WIDTH, ARENA_HEIGHT),
	}
	t.Menu = MenuInit(t)
	return t
}

func (traffiq *Traffiq) Update() error {
	if traffiq != nil {
		traffiq.TimersTick()
		shift := ebiten.IsKeyPressed(ebiten.KeyShift)
		control := ebiten.IsKeyPressed(ebiten.KeyControl)
		alt := ebiten.IsKeyPressed(ebiten.KeyAlt)
		for _, value := range GameKeys {
			if inpututil.IsKeyJustPressed(value) {
				traffiq.KeyPress(value, shift, control, alt)
				return nil
			}
		}
		if inpututil.IsMouseButtonJustPressed(ebiten.MouseButtonLeft) {
			traffiq.LeftMouseClick()
		}
		if inpututil.IsMouseButtonJustPressed(ebiten.MouseButtonRight) {
			traffiq.RightMouseClick()
		}
		if inpututil.IsMouseButtonJustPressed(ebiten.MouseButtonMiddle) {
			traffiq.MiddleMouseClick()
		}
		if traffiq.Board == BOARD_QUIT {
			return errors.New("Quit")
		}
	}
	return nil
}

func (g *Traffiq) AddTimer(name string, handler func(g *Traffiq)) {
	timer := &Timer{}
	timer.Init(g, handler)
	g.Timers[name] = timer
}

func (g *Traffiq) RemoveTimer(name string) {
	delete(g.Timers, name)
}

func (traffiq *Traffiq) TimersTick() {
	for _, timer := range traffiq.Timers {
		if timer.State {
			timer.Tick()
		}
	}
}

func (traffiq *Traffiq) StartTimer(name string, timeout float64) {
	if timer, exists := traffiq.Timers[name]; exists {
		timer.Start(timeout)
	}
}

func (traffiq *Traffiq) StopTimer(name string) {
	if timer, exists := traffiq.Timers[name]; exists {
		timer.Stop()
	}
}

func (traffiq *Traffiq) ChangeBoard(newboard Board) {
	fmt.Println(newboard)
	if traffiq.Board != newboard {
		switch traffiq.Board {
		case BOARD_LOADING:
			traffiq.StopWelcome()

		case BOARD_MENU:
			traffiq.StopMenu()
		case BOARD_GAME:
			traffiq.StopGame()
		case BOARD_OPTIONS:
			traffiq.StopOptions()
		case BOARD_HISCORES:
			traffiq.StopHiscores()
		case BOARD_SETTINGS:
			traffiq.StopSettings()
		case BOARD_NEWSCORE:
			traffiq.StopNewscore()
		case BOARD_HELP:
			traffiq.StopHelp()
		case BOARD_ABOUT:
			traffiq.StopAbout()
		case BOARD_QUIT:
			traffiq.StopQuit()
		}
		traffiq.Board = newboard
		switch newboard {
		case BOARD_LOADING:
			traffiq.InitWelcome()
		case BOARD_MENU:
			traffiq.InitMenu()
		case BOARD_GAME:
			traffiq.InitGame()
		case BOARD_OPTIONS:
			traffiq.InitOptions()
		case BOARD_HISCORES:
			traffiq.InitHiscores()
		case BOARD_SETTINGS:
			traffiq.InitSettings()
		case BOARD_NEWSCORE:
			traffiq.InitNewscore()
		case BOARD_HELP:
			traffiq.InitHelp()
		case BOARD_ABOUT:
			traffiq.InitAbout()
		case BOARD_QUIT:
			traffiq.InitQuit()
		}
	}
}

func (traffiq *Traffiq) KeyPress(key ebiten.Key, shift bool, control bool, alt bool) {
	switch traffiq.Board {
	case BOARD_LOADING:
		traffiq.KeyPressLoading(key, shift, control, alt)
	case BOARD_INTRO:
		traffiq.KeyPressIntro(key, shift, control, alt)
	case BOARD_MENU:
		traffiq.KeyPressMenu(key, shift, control, alt)
	case BOARD_GAME:
		traffiq.KeyPressGame(key, shift, control, alt)
	case BOARD_OPTIONS:
		traffiq.KeyPressOptions(key, shift, control, alt)
	case BOARD_LOADGAME:
		traffiq.KeyPressLoad(key, shift, control, alt)
	case BOARD_HISCORES:
		traffiq.KeyPressHiscores(key, shift, control, alt)
	case BOARD_SETTINGS:
		traffiq.KeyPressSettings(key, shift, control, alt)
	case BOARD_NEWSCORE:
		traffiq.KeyPressNewscore(key, shift, control, alt)
	case BOARD_ABOUT:
		traffiq.KeyPressAbout(key, shift, control, alt)
	case BOARD_HELP:
		traffiq.KeyPressHelp(key, shift, control, alt)
	case BOARD_QUIT:
		traffiq.KeyPressQuit(key, shift, control, alt)
	}
}

func (traffiq *Traffiq) KeyPressLoading(key ebiten.Key, shift bool, control bool, alt bool) {
	traffiq.ChangeBoard(BOARD_INTRO)
}

func (traffiq *Traffiq) KeyPressIntro(key ebiten.Key, shift bool, control bool, alt bool) {
	traffiq.ChangeBoard(BOARD_MENU)
}

func (traffiq *Traffiq) KeyPressOptions(key ebiten.Key, shift bool, control bool, alt bool) {
	switch key {
	case ebiten.KeyEscape:
		traffiq.ChangeBoard(BOARD_MENU)
	case ebiten.KeyQ:
		traffiq.ChangeBoard(BOARD_MENU)
	case ebiten.KeyArrowLeft:
		traffiq.ChangeBoard(BOARD_MENU)
	}
}

func (traffiq *Traffiq) KeyPressLoad(key ebiten.Key, shift bool, control bool, alt bool) {
	switch key {
	case ebiten.KeyEscape:
		traffiq.ChangeBoard(BOARD_MENU)
	case ebiten.KeyQ:
		traffiq.ChangeBoard(BOARD_MENU)
	case ebiten.KeyArrowLeft:
		traffiq.ChangeBoard(BOARD_MENU)
	}
}

func (traffiq *Traffiq) KeyPressSettings(key ebiten.Key, shift bool, control bool, alt bool) {
	switch key {
	case ebiten.KeyEscape:
		traffiq.ChangeBoard(BOARD_MENU)
	case ebiten.KeyQ:
		traffiq.ChangeBoard(BOARD_MENU)
	case ebiten.KeyArrowLeft:
		traffiq.ChangeBoard(BOARD_MENU)
	}
}

func (traffiq *Traffiq) KeyPressMenu(key ebiten.Key, shift bool, control bool, alt bool) {
	switch key {
	case ebiten.KeyArrowDown:
		if traffiq.Menu.Option < MAX_MENUOPTIONS-1 {
			traffiq.Menu.Option++
		}
	case ebiten.KeyArrowUp:
		if traffiq.Menu.Option > 0 {
			traffiq.Menu.Option--
		}
	case ebiten.KeyEnter:
		traffiq.ChangeBoard(menu2board[traffiq.Menu.Option])
	case ebiten.KeyF1:
		traffiq.ChangeBoard(BOARD_HELP)
	case ebiten.KeyQ:
		traffiq.ChangeBoard(BOARD_QUIT)
	}
}

func (traffiq *Traffiq) KeyPressGame(key ebiten.Key, shift bool, control bool, alt bool) {
	switch key {
	case ebiten.KeyEscape:
		traffiq.ChangeBoard(BOARD_MENU)
	case ebiten.KeyQ:
		traffiq.ChangeBoard(BOARD_MENU)
	case ebiten.KeyArrowLeft:
		traffiq.ChangeBoard(BOARD_MENU)
	}
}

func (traffiq *Traffiq) KeyPressHiscores(key ebiten.Key, shift bool, control bool, alt bool) {
	switch key {
	case ebiten.KeyEscape:
		traffiq.ChangeBoard(BOARD_MENU)
	case ebiten.KeyQ:
		traffiq.ChangeBoard(BOARD_MENU)
	case ebiten.KeyArrowLeft:
		traffiq.ChangeBoard(BOARD_MENU)
	}
}

func (traffiq *Traffiq) KeyPressNewscore(key ebiten.Key, shift bool, control bool, alt bool) {
	switch key {
	case ebiten.KeyEscape:
		traffiq.ChangeBoard(BOARD_MENU)
	case ebiten.KeyQ:
		traffiq.ChangeBoard(BOARD_MENU)
	case ebiten.KeyArrowLeft:
		traffiq.ChangeBoard(BOARD_MENU)
	}
}

func (traffiq *Traffiq) KeyPressQuit(key ebiten.Key, shift bool, control bool, alt bool) {

}

func (traffiq *Traffiq) KeyPressAbout(key ebiten.Key, shift bool, control bool, alt bool) {
	switch key {
	case ebiten.KeyEscape:
		traffiq.ChangeBoard(BOARD_MENU)
	case ebiten.KeyQ:
		traffiq.ChangeBoard(BOARD_MENU)
	case ebiten.KeyArrowLeft:
		traffiq.ChangeBoard(BOARD_MENU)
	}
}

func (traffiq *Traffiq) KeyPressHelp(key ebiten.Key, shift bool, control bool, alt bool) {
	switch key {
	case ebiten.KeyEscape:
		traffiq.ChangeBoard(BOARD_MENU)
	case ebiten.KeyQ:
		traffiq.ChangeBoard(BOARD_MENU)
	case ebiten.KeyArrowLeft:
		traffiq.ChangeBoard(BOARD_MENU)
	}
}

func (traffiq *Traffiq) InitWelcome() {
}

func (traffiq *Traffiq) InitMenu() {
	if traffiq != nil && traffiq.Menu != nil {
		traffiq.Menu.RecalculateRectangles()
	}
}

func (traffiq *Traffiq) InitGame() {
}

func (traffiq *Traffiq) InitOptions() {
}

func (traffiq *Traffiq) InitHiscores() {
}

func (traffiq *Traffiq) InitSettings() {
}

func (traffiq *Traffiq) InitNewscore() {
}

func (traffiq *Traffiq) InitHelp() {
}

func (traffiq *Traffiq) InitAbout() {
}

func (traffiq *Traffiq) InitQuit() {

}

func (traffiq *Traffiq) StopWelcome() {
}

func (traffiq *Traffiq) StopMenu() {
	traffiq.StopTimer("menu")
}

func (traffiq *Traffiq) StopGame() {
}

func (traffiq *Traffiq) StopOptions() {
}

func (traffiq *Traffiq) StopHiscores() {
}

func (traffiq *Traffiq) StopSettings() {
}

func (traffiq *Traffiq) StopNewscore() {
}

func (traffiq *Traffiq) StopHelp() {
}

func (traffiq *Traffiq) StopAbout() {
}

func (traffiq *Traffiq) StopQuit() {
}

func (traffiq *Traffiq) ChangeLang(lang int) {
	if lang >= 0 && lang < MAX_LANGUAGES && traffiq != nil {
		traffiq.Lang = lang
		if traffiq.Menu != nil {
			traffiq.Menu.RecalculateRectangles()
		}
	}
}
