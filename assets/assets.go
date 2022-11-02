package assets

import (
	_ "embed"
	"github.com/hajimehoshi/ebiten/v2"
)

var (
	// Plain images

	//go:embed background_default.png
	Embedded_Background_Default []byte
	//go:embed flag_en.png
	Flag_EN []byte
	//go:embed flag_pl.png
	Flag_PL []byte

	// TrueType Fonts
	//go:embed acknowtt.ttf
	Font_Acknowledge []byte
	//go:embed fredoka_one.ttf
	Font_Fredoka []byte
)

var (
	Background_Default *ebiten.Image
)
