package main

import (
	"gioui.org/app"
	"gioui.org/unit"
	"log"
	"os"
)

/*
Author Aidan Scott
Starts program of fictional news viewer called Passing Comet
it is named this way because a comet 'drops' off the news
*/
func main() {
	go func() {
		// create new window
		w := app.NewWindow(
			// set title
			app.Title("Passing Comet"),
			// set size
			app.Size(unit.Dp(600), unit.Dp(490)),
		)
		// if error is returned then log the error and exit
		if err := draw(w); err != nil {
			log.Fatal(err)
		}
		// exit nominally
		os.Exit(0)
	}()
	app.Main()

}
