package main

import (
	"gioui.org/app"
	"gioui.org/unit"
	"log"
	"os"
)

func main() {
	go func() {
		// create new window
		w := app.NewWindow(
			app.Title("Passing Comet"),
			app.Size(unit.Dp(600), unit.Dp(490)),
		)
		if err := draw(w); err != nil {
			log.Fatal(err)
		}
		os.Exit(0)
	}()
	app.Main()

}
