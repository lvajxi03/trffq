package main

import (
	"image/color"
)

const (
	ARENA_WIDTH  = 1920
	ARENA_HEIGHT = 1080
)

const (
	APPLICATION_TITLE = "Trffq"
)

type Board int

const (
	BOARD_LOADING Board = iota
	BOARD_INTRO
	BOARD_MENU
	BOARD_GAME
	BOARD_OPTIONS
	BOARD_LOADGAME
	BOARD_HISCORES
	BOARD_SETTINGS
	BOARD_NEWSCORE
	BOARD_HELP
	BOARD_ABOUT
	BOARD_QUIT
)

type Mode int

const (
	MODE_INIT Mode = iota
	MODE_PLAY
	MODE_PAUSED
	MODE_KILLED
	MODE_GAMEOVER
)

type Language int

const (
	LANG_POLISH Language = iota
	LANG_ENGLISH
)

const MAX_LANGUAGES = 2
const MAX_MENUOPTIONS = 8

const MENU_LETTERS_DELTA = 15

const SHADOW_MARGIN = 5

var colors = map[string]color.Color {
	"default-font-shadow": color.RGBA{30, 30, 30, 128},
	"default-font-inactive": color.RGBA{255, 193, 0, 224},
	"default-font-active": color.RGBA{180, 46, 178, 224},
}

var menu2board = map[int]Board {
	0: BOARD_GAME,
	1: BOARD_OPTIONS,
	2: BOARD_LOADGAME,
	3: BOARD_HISCORES,
	4: BOARD_SETTINGS,	
	5: BOARD_HELP,
	6: BOARD_ABOUT,
	7: BOARD_QUIT,
}
