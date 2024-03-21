package main

import (
	"fmt"
	"gioui.org/app"
	"gioui.org/layout"
	"gioui.org/op"
	"gioui.org/unit"
	"gioui.org/widget"
	"gioui.org/widget/material"
)

const (
	system = iota
	local
)

func draw(w *app.Window) error {
	// logic operators
	newsWindow := system
	// Window variables
	var ops op.Ops
	var localButton widget.Clickable
	var systemButton widget.Clickable
	th := material.NewTheme()
	for {
		// detect what type of event
		switch env := w.NextEvent().(type) {

		// this is sent when the application should re-render.
		case app.FrameEvent:
			gtx := app.NewContext(&ops, env)
			layout.Flex{
				Axis:    layout.Vertical,
				Spacing: layout.SpaceStart,
			}.Layout(gtx,
				layout.Rigid(
					func(gtx layout.Context) layout.Dimensions {
						// ONE: First define margins around the button using layout.Inset ...
						margins := getButtonMargins()
						// TWO: ... then we lay out those margins ...
						return margins.Layout(gtx,
							func(gtx layout.Context) layout.Dimensions {
								for localButton.Clicked(gtx) {
									newsWindow = local
								}
								btn := material.Button(th, &localButton, "Local news")
								return btn.Layout(gtx)
							},
						)
					},
				),
				layout.Rigid(
					func(gtx layout.Context) layout.Dimensions {
						// ONE: First define margins around the button using layout.Inset ...
						margins := getButtonMargins()
						// TWO: ... then we lay out those margins ...
						return margins.Layout(gtx,
							func(gtx layout.Context) layout.Dimensions {
								for systemButton.Clicked(gtx) {
									newsWindow = system
								}
								btn := material.Button(th, &systemButton, "System news")
								return btn.Layout(gtx)
							},
						)
					},
				),
				layout.Rigid(
					layout.Spacer{Height: unit.Dp(25)}.Layout,
				),
			)
			// ...
			fmt.Println(newsWindow)
			env.Frame(gtx.Ops)
		// this is sent when the application is closed
		case app.DestroyEvent:
			return env.Err
		}
	}
}
