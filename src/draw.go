package main

import (
	"gioui.org/app"
	"gioui.org/layout"
	"gioui.org/op"
	"gioui.org/text"
	"gioui.org/widget"
	"gioui.org/widget/material"
)

/*
	Author: Aidan Scott
	This file is designated to managing the logic for the visuals across the application.
	Some tasks are delegated to utilities.go for convenience
*/
// Const is similar to using an enum. in this case I use it to keep track of which buttons are displayed
const (
	//f ull solar system news
	system = iota
	// only local news
	local
)

/*
draw is the heavy lifting function of this file. It adds all of the elements to the gui, and manages the events
so a button press actually does something.
Some usage is delegated to the utilities file
*/
func draw(w *app.Window) error {
	// Window variables, variables that the window must have to run
	var ops op.Ops
	th := material.NewTheme() // default theme used for Gio

	// logic operators
	var newsWindow = system                     // state tracking
	var localButton widget.Clickable            // button for changing state to local
	var systemButton widget.Clickable           // button for changing state to system
	sysBtns := map[string]widget.Clickable{}    // map of clickable widgets that will be used to create buttons
	sysArtBtns := map[string]layout.FlexChild{} // map of buttons with event handling included
	// loop through system news and print each map's key
	for k := range getSystemNews() {
		sysBtns[k] = widget.Clickable{}
		btn := sysBtns[k] // I had issues casting the value of sysBtns[k] to a pointer in the next function so I moved it before
		// creating each button for only articles then add them to the map
		sysArtBtns[k] = getButtonClickable(&btn, getSystemNews()[k][1], getSystemNews()[k][0])
	}
	// look above at comments to understand this block
	locBtns := map[string]widget.Clickable{}
	locArtBtns := map[string]layout.FlexChild{}
	for k := range getLocalNews() {
		locBtns[k] = widget.Clickable{}
		btn := locBtns[k]
		locArtBtns[k] = getButtonClickable(&btn, getLocalNews()[k][1], getLocalNews()[k][0])
	}
	// To give the user a sense of the world, I have included a starting info panel for the users past
	// This uses the article window generator to quickly spin up a info panel
	dispalyStartingInfoWindow()
	// start the loop that refreshes the page
	for {
		// detect what type of event
		switch env := w.NextEvent().(type) {

		// this is sent when the application should re-render.
		case app.FrameEvent:
			gtx := app.NewContext(&ops, env) // creating new context to be used in child functions
			// setting values for relevant layout options
			layout.Flex{
				Axis:    layout.Vertical,
				Spacing: layout.SpaceStart,
			}.Layout(gtx,
				// Rigid to display the title of the fictional news provider
				layout.Rigid(func(gtx layout.Context) layout.Dimensions {
					return getButtonMargins().Layout(gtx,
						func(gtx layout.Context) layout.Dimensions {
							// create H4 and align it to middle then return the dimensions placing it where it needs to go
							info := material.H4(th, "Welcome to Comet News")
							info.Alignment = text.Middle
							return info.Layout(gtx)
						})
				}),
				// Rigid for local button
				layout.Rigid(
					func(gtx layout.Context) layout.Dimensions {
						// margins to create spacing around button
						// function from utilities
						margins := getButtonMargins()
						return margins.Layout(gtx,
							func(gtx layout.Context) layout.Dimensions {
								// if localButton is clicked, then newsWindow will be set to local
								for localButton.Clicked(gtx) {
									newsWindow = local
								}
								// create physical button, then return dimensions
								btn := material.Button(th, &localButton, "Local news")
								return btn.Layout(gtx)
							},
						)
					},
				),
				// Rigid for system button
				layout.Rigid(
					func(gtx layout.Context) layout.Dimensions {
						// margins to create spacing around button
						// function from utilities
						margins := getButtonMargins()
						return margins.Layout(gtx,
							func(gtx layout.Context) layout.Dimensions {
								// if systemButton is clicked, then newsWindow will be set to local
								for systemButton.Clicked(gtx) {
									newsWindow = system
								}
								// create physical button, then return dimensions
								btn := material.Button(th, &systemButton, "System news")
								return btn.Layout(gtx)
							},
						)
					},
				),
				// Rigid for H6 stating to the user that next buttons are articles
				layout.Rigid(func(gtx layout.Context) layout.Dimensions {
					return getButtonMargins().Layout(gtx,
						func(gtx layout.Context) layout.Dimensions {
							// Create H6, align it and return its dimensions
							info := material.H6(th, "News Articles")
							info.Alignment = text.Middle
							return info.Layout(gtx)
						})
				}),
				// create button depending on state
				layout.Rigid(
					func(gtx layout.Context) layout.Dimensions {
						// set dem depending on state, changing what is displayed to the gui
						dem := layout.Dimensions{}
						if newsWindow == system {
							// large part of this code is delegated to utilities
							dem = rows.Layout(gtx, sysArtBtns["venus"])
						}
						if newsWindow == local {
							// large part of this code is delegated to utilities
							dem = rows.Layout(gtx, locArtBtns["updates"])
						}
						return dem
					},
				),
				layout.Rigid(
					func(gtx layout.Context) layout.Dimensions {
						// set dem depending on state, changing what is displayed to the gui
						dem := layout.Dimensions{}
						if newsWindow == system {
							// large part of this code is delegated to utilities
							dem = rows.Layout(gtx, sysArtBtns["mirror"])
						}
						if newsWindow == local {
							// large part of this code is delegated to utilities
							dem = rows.Layout(gtx, locArtBtns["comets"])
						}
						return dem
					},
				),
				layout.Rigid(
					func(gtx layout.Context) layout.Dimensions {
						// set dem depending on state, changing what is displayed to the gui
						dem := layout.Dimensions{}
						if newsWindow == system {
							// large part of this code is delegated to utilities
							dem = rows.Layout(gtx, sysArtBtns["work"])
						}
						if newsWindow == local {
							// large part of this code is delegated to utilities
							dem = rows.Layout(gtx, locArtBtns["carriers"])
						}
						return dem
					},
				),
			)
			// display configured frame. this is what visualizes everything
			env.Frame(gtx.Ops)
		// this is sent when the application is closed
		case app.DestroyEvent:
			return env.Err
		}
	}
}
