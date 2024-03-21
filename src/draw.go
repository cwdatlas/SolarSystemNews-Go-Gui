package main

import (
	"gioui.org/app"
	"gioui.org/layout"
	"gioui.org/op"
	"gioui.org/text"
	"gioui.org/widget"
	"gioui.org/widget/material"
)

const (
	system = iota
	local
)

func draw(w *app.Window) error {
	// Window variables
	var ops op.Ops
	var localButton widget.Clickable
	var systemButton widget.Clickable
	th := material.NewTheme()

	// logic operators
	var newsWindow = system
	sysBtns := map[string]widget.Clickable{}
	sysArtBtns := map[string]layout.FlexChild{}
	for k := range getSystemNews() {
		sysBtns[k] = widget.Clickable{}
		btn := sysBtns[k]
		sysArtBtns[k] = getButtonClickable(&btn, getSystemNews()[k][1], getSystemNews()[k][0])
	}

	locBtns := map[string]widget.Clickable{}
	locArtBtns := map[string]layout.FlexChild{}
	for k := range getLocalNews() {
		locBtns[k] = widget.Clickable{}
		btn := locBtns[k]
		locArtBtns[k] = getButtonClickable(&btn, getLocalNews()[k][1], getLocalNews()[k][0])
	}

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
				layout.Rigid(func(gtx layout.Context) layout.Dimensions {
					return getButtonMargins().Layout(gtx,
						func(gtx layout.Context) layout.Dimensions {
							info := material.H4(th, "Welcome to Comet News")
							info.Alignment = text.Middle
							return info.Layout(gtx)
						})
				}),
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
				layout.Rigid(func(gtx layout.Context) layout.Dimensions {
					return getButtonMargins().Layout(gtx,
						func(gtx layout.Context) layout.Dimensions {
							info := material.H6(th, "News Articles From the Comet")
							info.Alignment = text.Middle
							return info.Layout(gtx)
						})
				}),
				layout.Rigid(
					func(gtx layout.Context) layout.Dimensions {
						dem := layout.Dimensions{}
						if newsWindow == system {
							dem = rows.Layout(gtx, sysArtBtns["venus"])
						}
						if newsWindow == local {
							dem = rows.Layout(gtx, locArtBtns["updates"])
						}
						return dem
					},
				),
				layout.Rigid(
					func(gtx layout.Context) layout.Dimensions {
						dem := layout.Dimensions{}
						if newsWindow == system {
							dem = rows.Layout(gtx, sysArtBtns["mirror"])
						}
						if newsWindow == local {
							dem = rows.Layout(gtx, locArtBtns["comets"])
						}
						return dem
					},
				),
				layout.Rigid(
					func(gtx layout.Context) layout.Dimensions {
						dem := layout.Dimensions{}
						if newsWindow == system {
							dem = rows.Layout(gtx, sysArtBtns["work"])
						}
						if newsWindow == local {
							dem = rows.Layout(gtx, locArtBtns["carriers"])
						}
						return dem
					},
				),
			)
			env.Frame(gtx.Ops)
		// this is sent when the application is closed
		case app.DestroyEvent:
			return env.Err
		}
	}
}
