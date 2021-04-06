package ui

import "github.com/gdamore/tcell/v2"

type mycolors struct {
	bg0_h   tcell.Color
	bg0     tcell.Color
	bg1     tcell.Color
	bg2     tcell.Color
	bg3     tcell.Color
	bg4     tcell.Color
	fg      tcell.Color
	fg0     tcell.Color
	fg1     tcell.Color
	fg2     tcell.Color
	fg3     tcell.Color
	fg4     tcell.Color
	red     tcell.Color
	green   tcell.Color
	yellow  tcell.Color
	blue    tcell.Color
	purple  tcell.Color
	aqua    tcell.Color
	orange  tcell.Color
	red1    tcell.Color
	green1  tcell.Color
	yellow1 tcell.Color
	blue1   tcell.Color
	purple1 tcell.Color
	aqua1   tcell.Color
	orange1 tcell.Color
	gray1   tcell.Color
	gray2   tcell.Color
}

func newcolors() *mycolors {
	colors := &mycolors{
		bg0_h:   tcell.NewHexColor(0x1d2021),
		bg0:     tcell.NewHexColor(0x282828),
		bg1:     tcell.NewHexColor(0x3c3836),
		bg2:     tcell.NewHexColor(0x504945),
		bg3:     tcell.NewHexColor(0x665c54),
		bg4:     tcell.NewHexColor(0x7c6f64),
		red:     tcell.NewHexColor(0xcc241d),
		green:   tcell.NewHexColor(0x98971a),
		yellow:  tcell.NewHexColor(0xd79921),
		blue:    tcell.NewHexColor(0x458588),
		purple:  tcell.NewHexColor(0xb16286),
		aqua:    tcell.NewHexColor(0x689d6a),
		orange:  tcell.NewHexColor(0xd65d0e),
		red1:    tcell.NewHexColor(0xfb4934),
		green1:  tcell.NewHexColor(0xb8bb26),
		yellow1: tcell.NewHexColor(0xfabd2f),
		blue1:   tcell.NewHexColor(0x83a598),
		purple1: tcell.NewHexColor(0xd3869b),
		aqua1:   tcell.NewHexColor(0x8ec07c),
		orange1: tcell.NewHexColor(0xfe8019),
		gray1:   tcell.NewHexColor(0xa89984),
		gray2:   tcell.NewHexColor(0x928374),
		fg:      tcell.NewHexColor(0xebdbb2),
		fg0:     tcell.NewHexColor(0xfbf1c7),
		fg1:     tcell.NewHexColor(0xedbbb2),
		fg2:     tcell.NewHexColor(0xd5c4a1),
		fg3:     tcell.NewHexColor(0xbdae93),
		fg4:     tcell.NewHexColor(0xa89984),
	}

	return colors
}
