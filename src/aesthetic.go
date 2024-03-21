package main

import (
	"gioui.org/layout"
	"gioui.org/unit"
)

var buttonMargins = layout.Inset{
	Top:    unit.Dp(25),
	Bottom: unit.Dp(25),
	Right:  unit.Dp(35),
	Left:   unit.Dp(35),
}

func getButtonMargins() layout.Inset {
	return buttonMargins
}
