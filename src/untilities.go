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

/*
	Author Aidan Scott
	Utilities holds many nice to have separate functions so config can be called once in the name of DRY
*/
// config for button margins
var buttonMargins = layout.Inset{
	Top:    unit.Dp(15),
	Bottom: unit.Dp(15),
	Right:  unit.Dp(35),
	Left:   unit.Dp(35),
}

// config for rows, used in article button display
var rows = layout.Flex{Axis: layout.Vertical, Spacing: layout.SpaceEvenly}

var intro = "You are one of the many tens of thousands of lone Asteroid miners out in a distant section\n" +
	"of the belt. You grew up on the Martian colonies, but you didn't like the couped up life of the domed habitats.\n" +
	"You wanted to be an explorer like the people who founded your home colony, so you decided to use your savings\n" +
	"and move out into the belt where you can explore and profit heavily from doing what you love...\n" +
	"At least that is what you thought, before you were again couped up in a little metal can with a bunch of\n" +
	"'Space Bumpkins' on the coms and local news channel. Your ship is an old throwaway Hanwa with the internet ingress\n" +
	"point ripped out by its previous owners to make an extra YenDollar. You had to switch to the emergency minimum connection ingress.\n" +
	"You are now stuck bored between farming missions.\n" +
	"You recently bought a replacement of the emergency ingress system, so now you have more than a trickle of bandwidth!\n" +
	"In between your farming missions you decide to pull together a program for AP-089, your local system, to view everyone's\n" +
	"news more easily. So you boot up your new program for the first time and see if it works..."

// return margin config
func getButtonMargins() layout.Inset {
	return buttonMargins
}
func dispalyStartingInfoWindow() {
	createInfoWindow(intro, "background")
}

// return Rigid that displays a button and manages its event
func getButtonClickable(btn *widget.Clickable, article string, name string) layout.FlexChild {
	th := material.NewTheme()
	articleButton := layout.Rigid(
		func(gtx layout.Context) layout.Dimensions {
			// sets margin config
			margins := getButtonMargins()
			return margins.Layout(gtx, func(gtx layout.Context) layout.Dimensions {
				// if button is clicked then display new window with article text
				for btn.Clicked(gtx) {
					createInfoWindow(article, name)
				}
				// set button then return dims
				btn := material.Button(th, btn, name)
				return btn.Layout(gtx)
			})
		})
	// return flex child for visualization
	return articleButton
}

/*
	creates a new window of specific size with the given params
	article: is the entire string related to the article
	name: is the articles name
*/

func createInfoWindow(article string, name string) {
	go func() {
		// set window
		w := app.NewWindow(
			// set title to name
			app.Title(name),
			// set size
			app.Size(unit.Dp(1350), unit.Dp(400)),
		)
		// if error stop window, close it
		if err := infoPanelLoop(w, article); err != nil {
			panic(err)
		}
	}()

}

/*
'draw' but for the info panel
This function provides the structure for elements and displays the article and a button to exit
w: *app.Window: used to tell what window is being added to
article: full string of the article
*/
func infoPanelLoop(w *app.Window, article string) error {
	// setting dependents
	th := material.NewTheme()
	var ops op.Ops
	// creating ok button
	okButton := new(widget.Clickable)
	for {
		// setting up switch so app can be drawn when events occur
		switch env := w.NextEvent().(type) {
		case app.FrameEvent:
			// set up context
			gtx := app.NewContext(&ops, env)
			// set layout variables of the layout
			layout.Flex{
				Axis:    layout.Vertical, // elements will be displayed vertically
				Spacing: layout.SpaceStart,
			}.Layout(gtx,
				// Rigid to display the article
				layout.Rigid(func(gtx layout.Context) layout.Dimensions {
					// set, align then return article's dims
					article := material.H6(th, article)
					article.Alignment = text.Middle
					return article.Layout(gtx)
				},
				),
				// Rigid to display button
				layout.Rigid(func(gtx layout.Context) layout.Dimensions {
					// Layout the "OK" button at the bottom
					for okButton.Clicked(gtx) {
						w.Perform(system2.ActionClose) // This will close the window
					}
					// create physical button and return dims
					btn := material.Button(th, okButton, "Close")
					return btn.Layout(gtx)
				}),
			)
			// draw frame
			env.Frame(gtx.Ops)
		case app.DestroyEvent:
			// on destroy shut dow process
			return env.Err

		}
	}
}
