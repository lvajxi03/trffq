package main

import (
	"github.com/hajimehoshi/ebiten/v2/text"
	"github.com/lvajxi03/primi"
	"github.com/lvajxi03/trffq/graphics"
)

type Menu struct {
	Parent         *Traffiq
	Option         int
	HeaderLeft     float64
	Rectangles     [MAX_MENUOPTIONS]*primi.Rect
	LangRectangles [MAX_LANGUAGES]*primi.Rect
}

func MenuInit(parent *Traffiq) *Menu {
	m := &Menu{
		Parent: parent,
	}
	m.LangRectangles[1] = primi.NewRect(ARENA_WIDTH-160, ARENA_HEIGHT-200, 80, 60)
	m.LangRectangles[0] = primi.NewRect(ARENA_WIDTH-80, ARENA_HEIGHT-200, 80, 60)
	m.RecalculateRectangles()
	return m
}

func (menu *Menu) RecalculateRectangles() {
	if menu != nil {
		if menu.Parent != nil {
			if menu.Parent.Lang >= 0 && menu.Parent.Lang < MAX_LANGUAGES {
				r := text.BoundString(Face_FredokaLarge, titles[menu.Parent.Lang][BOARD_MENU])
				menu.HeaderLeft = float64((ARENA_WIDTH - r.Max.X + r.Min.X) / 2)

				for index, title := range menustr[menu.Parent.Lang] {
					r1 := text.BoundString(Face_Fredoka, title)
					r2 := graphics.FontBounds2Rect(r1, 355, 355+80*index, 5)
					menu.Rectangles[index] = r2
				}
			}
		}
	}
}

func (menu *Menu) SelectOption(option int) {
	if menu != nil {
		if option >= 0 && option < MAX_MENUOPTIONS {
			menu.Option = option
		}
	}
}
