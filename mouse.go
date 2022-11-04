package main

import (
	"github.com/hajimehoshi/ebiten/v2"
)

func (traffiq *Traffiq) LeftMouseClick() {
	if traffiq != nil {
		switch traffiq.Board {
		case BOARD_LOADING:
			traffiq.LeftMouseClickLoading()
		case BOARD_INTRO:
			traffiq.LeftMouseClickIntro()
		case BOARD_MENU:
			traffiq.LeftMouseClickMenu()
		case BOARD_GAME:
			traffiq.LeftMouseClickGame()
		case BOARD_OPTIONS:
			traffiq.LeftMouseClickOptions()
		case BOARD_LOADGAME:
			traffiq.LeftMouseClickLoadgame()
		case BOARD_HISCORES:
			traffiq.LeftMouseClickHiscores()
		case BOARD_SETTINGS:
			traffiq.LeftMouseClickSettings()
		case BOARD_NEWSCORE:
			traffiq.LeftMouseClickNewscore()
		case BOARD_HELP:
			traffiq.LeftMouseClickHelp()
		case BOARD_ABOUT:
			traffiq.LeftMouseClickAbout()
		}
	}
}

func (traffiq *Traffiq) RightMouseClick() {
	if traffiq != nil {
		switch traffiq.Board {
		case BOARD_LOADING:
			traffiq.RightMouseClickLoading()
		case BOARD_INTRO:
			traffiq.RightMouseClickIntro()
		case BOARD_MENU:
			traffiq.RightMouseClickMenu()
		case BOARD_GAME:
			traffiq.RightMouseClickGame()
		case BOARD_OPTIONS:
			traffiq.RightMouseClickOptions()
		case BOARD_LOADGAME:
			traffiq.RightMouseClickLoadgame()
		case BOARD_HISCORES:
			traffiq.RightMouseClickHiscores()
		case BOARD_SETTINGS:
			traffiq.RightMouseClickSettings()
		case BOARD_NEWSCORE:
			traffiq.RightMouseClickNewscore()
		case BOARD_HELP:
			traffiq.RightMouseClickHelp()
		case BOARD_ABOUT:
			traffiq.RightMouseClickAbout()
		}
	}
}

func (traffiq *Traffiq) MiddleMouseClick() {
	if traffiq != nil {
		switch traffiq.Board {
		case BOARD_LOADING:
			traffiq.MiddleMouseClickLoading()
		case BOARD_INTRO:
			traffiq.MiddleMouseClickLoading()
		case BOARD_MENU:
			traffiq.MiddleMouseClickMenu()
		case BOARD_GAME:
			traffiq.MiddleMouseClickGame()
		case BOARD_OPTIONS:
			traffiq.MiddleMouseClickOptions()
		case BOARD_LOADGAME:
			traffiq.MiddleMouseClickLoadgame()
		case BOARD_HISCORES:
			traffiq.MiddleMouseClickHiscores()
		case BOARD_SETTINGS:
			traffiq.MiddleMouseClickSettings()
		case BOARD_NEWSCORE:
			traffiq.MiddleMouseClickNewscore()
		case BOARD_HELP:
			traffiq.MiddleMouseClickHelp()
		case BOARD_ABOUT:
			traffiq.MiddleMouseClickAbout()
		}
	}
}

func (traffiq *Traffiq) LeftMouseClickLoading() {
	if traffiq != nil {
		traffiq.ChangeBoard(BOARD_INTRO)
	}
}

func (traffiq *Traffiq) LeftMouseClickIntro() {
	if traffiq != nil {
		traffiq.ChangeBoard(BOARD_MENU)
	}
}

func (traffiq *Traffiq) LeftMouseClickMenu() {
	if traffiq != nil {
		x, y := ebiten.CursorPosition()
		for index, rect := range traffiq.Menu.Rectangles {
			if rect.Contains(float64(x), float64(y)) {
				traffiq.ChangeBoard(menu2board[index])
			}
		}
		for index, rect := range traffiq.Menu.LangRectangles {

			if rect.Contains(float64(x), float64(y)) {
				traffiq.ChangeLang(index)
			}
		}
	}
}

func (traffiq *Traffiq) Temporary_Placeholder() {
	// TODO: remove me
}

func (traffiq *Traffiq) LeftMouseClickGame() {
	if traffiq != nil {
		switch Traffiq.Mode {
		case MODE_INIT:
			traffiq.Temporary_Placeholder()
		case MODE_PLAY:
			traffiq.Temporary_Placeholder()
		case MODE_PAUSED:
			traffiq.Temporary_Placeholder()
		}
	}
}

func (traffiq *Traffiq) LeftMouseClickOptions() {
	if traffiq != nil {
		traffiq.Temporary_Placeholder()
	}
}

func (traffiq *Traffiq) LeftMouseClickLoadgame() {
	if traffiq != nil {
		traffiq.Temporary_Placeholder()

	}
}

func (traffiq *Traffiq) LeftMouseClickHiscores() {
	if traffiq != nil {
		x, y := ebiten.CursorPosition()
		for index, rect := range traffiq.Menu.LangRectangles {

			if rect.Contains(float64(x), float64(y)) {
				traffiq.ChangeLang(index)
				return
			}
		}
		traffiq.ChangeBoard(BOARD_MENU)
	}
}

func (traffiq *Traffiq) LeftMouseClickSettings() {
	if traffiq != nil {
		traffiq.Temporary_Placeholder()

	}
}

func (traffiq *Traffiq) LeftMouseClickNewscore() {
	if traffiq != nil {
		traffiq.Temporary_Placeholder()

	}
}

func (traffiq *Traffiq) LeftMouseClickHelp() {
	if traffiq != nil {
		x, y := ebiten.CursorPosition()
		for index, rect := range traffiq.Menu.LangRectangles {

			if rect.Contains(float64(x), float64(y)) {
				traffiq.ChangeLang(index)
				return
			}
		}
		traffiq.ChangeBoard(BOARD_MENU)
	}
}

func (traffiq *Traffiq) LeftMouseClickAbout() {
	if traffiq != nil {
		x, y := ebiten.CursorPosition()
		for index, rect := range traffiq.Menu.LangRectangles {

			if rect.Contains(float64(x), float64(y)) {
				traffiq.ChangeLang(index)
				return
			}
		}
		traffiq.ChangeBoard(BOARD_MENU)
	}
}

func (traffiq *Traffiq) RightMouseClickLoading() {
	if traffiq != nil {
		traffiq.ChangeBoard(BOARD_INTRO)
	}
}

func (traffiq *Traffiq) RightMouseClickIntro() {
	if traffiq != nil {
		traffiq.ChangeBoard(BOARD_MENU)
	}
}

func (traffiq *Traffiq) RightMouseClickMenu() {
	if traffiq != nil {
		traffiq.Temporary_Placeholder()

	}
}

func (traffiq *Traffiq) RightMouseClickGame() {
	if traffiq != nil {
		traffiq.Temporary_Placeholder()

	}
}

func (traffiq *Traffiq) RightMouseClickOptions() {
	if traffiq != nil {
		traffiq.Temporary_Placeholder()

	}
}

func (traffiq *Traffiq) RightMouseClickLoadgame() {
	if traffiq != nil {
		traffiq.Temporary_Placeholder()

	}
}

func (traffiq *Traffiq) RightMouseClickHiscores() {
	if traffiq != nil {
		traffiq.Temporary_Placeholder()

	}
}

func (traffiq *Traffiq) RightMouseClickSettings() {
	if traffiq != nil {
		traffiq.Temporary_Placeholder()

	}
}

func (traffiq *Traffiq) RightMouseClickNewscore() {
	if traffiq != nil {
		traffiq.Temporary_Placeholder()

	}
}

func (traffiq *Traffiq) RightMouseClickHelp() {
	if traffiq != nil {
		traffiq.Temporary_Placeholder()

	}
}

func (traffiq *Traffiq) RightMouseClickAbout() {
	if traffiq != nil {
		traffiq.Temporary_Placeholder()

	}
}

func (traffiq *Traffiq) MiddleMouseClickLoading() {
	if traffiq != nil {
		traffiq.Temporary_Placeholder()

	}
}

func (traffiq *Traffiq) MiddleMouseClickIntro() {
	if traffiq != nil {
		traffiq.Temporary_Placeholder()

	}
}

func (traffiq *Traffiq) MiddleMouseClickMenu() {
	if traffiq != nil {
		traffiq.Temporary_Placeholder()

	}
}

func (traffiq *Traffiq) MiddleMouseClickGame() {
	if traffiq != nil {
		traffiq.Temporary_Placeholder()

	}
}

func (traffiq *Traffiq) MiddleMouseClickOptions() {
	if traffiq != nil {
		traffiq.Temporary_Placeholder()

	}
}

func (traffiq *Traffiq) MiddleMouseClickLoadgame() {
	if traffiq != nil {
		traffiq.Temporary_Placeholder()

	}
}

func (traffiq *Traffiq) MiddleMouseClickHiscores() {
	if traffiq != nil {
		traffiq.Temporary_Placeholder()

	}
}

func (traffiq *Traffiq) MiddleMouseClickSettings() {
	if traffiq != nil {
		traffiq.Temporary_Placeholder()

	}
}

func (traffiq *Traffiq) MiddleMouseClickNewscore() {
	if traffiq != nil {
		traffiq.Temporary_Placeholder()

	}
}

func (traffiq *Traffiq) MiddleMouseClickHelp() {
	if traffiq != nil {
		traffiq.Temporary_Placeholder()

	}
}

func (traffiq *Traffiq) MiddleMouseClickAbout() {
	if traffiq != nil {
		traffiq.Temporary_Placeholder()

	}
}
