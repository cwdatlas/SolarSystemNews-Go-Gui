package main

import (
	"gioui.org/app"
	system2 "gioui.org/io/system"
	"gioui.org/layout"
	"gioui.org/op"
	"gioui.org/text"
	"gioui.org/unit"
	"gioui.org/widget"
	"gioui.org/widget/material"
)

var buttonMargins = layout.Inset{
	Top:    unit.Dp(15),
	Bottom: unit.Dp(15),
	Right:  unit.Dp(35),
	Left:   unit.Dp(35),
}
var rows = layout.Flex{Axis: layout.Vertical, Spacing: layout.SpaceEvenly}

func getButtonMargins() layout.Inset {
	return buttonMargins
}

func getButtonClickable(btn *widget.Clickable, article string, name string) layout.FlexChild {
	th := material.NewTheme()
	articleButton := layout.Rigid(
		func(gtx layout.Context) layout.Dimensions {
			margins := getButtonMargins()
			return margins.Layout(gtx, func(gtx layout.Context) layout.Dimensions {
				for btn.Clicked(gtx) {
					createInfoWindow(article, name)
				}
				btn := material.Button(th, btn, name)
				return btn.Layout(gtx)
			})
		})
	return articleButton
}

func createInfoWindow(article string, name string) {
	go func() {
		w := app.NewWindow(
			app.Title(name),
			app.Size(unit.Dp(1350), unit.Dp(200)),
		)
		if err := infoPanelLoop(w, article); err != nil {
			panic(err)
		}
	}()

}
func infoPanelLoop(w *app.Window, article string) error {
	th := material.NewTheme()
	okButton := new(widget.Clickable)
	var ops op.Ops
	for {
		switch env := w.NextEvent().(type) {
		case app.FrameEvent:
			gtx := app.NewContext(&ops, env)
			layout.Flex{
				Axis:    layout.Vertical,
				Spacing: layout.SpaceStart,
			}.Layout(gtx,

				layout.Rigid(func(gtx layout.Context) layout.Dimensions {
					article := material.H6(th, article)
					article.Alignment = text.Middle
					return article.Layout(gtx)
				},
				),
				layout.Rigid(func(gtx layout.Context) layout.Dimensions {
					// Layout the "OK" button at the bottom
					for okButton.Clicked(gtx) {
						w.Perform(system2.ActionClose) // This will close the window
					}
					btn := material.Button(th, okButton, "Close")
					return btn.Layout(gtx)
				}),
			)
			env.Frame(gtx.Ops)
		case app.DestroyEvent:
			return env.Err

		}
	}
}
