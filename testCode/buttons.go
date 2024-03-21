package testCode

import (
	"fmt"
	"gioui.org/app"
	"gioui.org/layout"
	"gioui.org/op"
	"gioui.org/unit"
	"gioui.org/widget"
	"gioui.org/widget/material"
	"log"
)

/*
	func main() {
		go startApp()
		app.Main()
	}
*/
func startApp() {
	th := material.NewTheme()
	w := app.NewWindow(
		app.Title("Dynamic Buttons"),
		app.Size(unit.Dp(800), unit.Dp(600)),
	)

	if err := loop(w, th); err != nil {
		log.Fatal(err)
	}
}

func loop(w *app.Window, th *material.Theme) error {
	var ops op.Ops
	button1 := new(widget.Clickable)
	button2 := new(widget.Clickable)
	buttonMap1 := map[string]*widget.Clickable{}
	buttonMap2 := map[string]*widget.Clickable{}

	// Sample data for our buttons
	data1 := map[string]string{"A": "Apple", "B": "Banana", "C": "Cherry"}
	data2 := map[string]string{"1": "One", "2": "Two", "3": "Three"}

	list := layout.List{Axis: layout.Vertical}

	showData1 := false
	showData2 := false

	for {
		switch e := w.NextEvent().(type) {
		case app.DestroyEvent:
			return e.Err
		case app.FrameEvent:
			gtx := app.NewContext(&ops, e)

			layout.Flex{Axis: layout.Horizontal}.Layout(gtx,
				layout.Flexed(0.5, func(gtx layout.Context) layout.Dimensions {
					for button1.Clicked(gtx) {
						showData1 = true
						showData2 = false
					}
					return material.Button(th, button1, "Show Data 1").Layout(gtx)
				}),
				layout.Flexed(0.5, func(gtx layout.Context) layout.Dimensions {
					for button2.Clicked(gtx) {
						showData1 = false
						showData2 = true
					}
					return material.Button(th, button2, "Show Data 2").Layout(gtx)
				}),
			)

			if showData1 {
				displayButtons(th, &gtx, &list, data1, buttonMap1)
			}
			if showData2 {
				displayButtons(th, &gtx, &list, data2, buttonMap2)
			}

			e.Frame(gtx.Ops)
		}
	}

}

func displayButtons(th *material.Theme, gtx *layout.Context, list *layout.List, data map[string]string, buttons map[string]*widget.Clickable) {
	list.Layout(*gtx, len(data), func(gtx layout.Context, i int) layout.Dimensions {
		var key string
		var value string
		var count int
		for k, v := range data {
			if count == i {
				key = k
				value = v
				break
			}
			count++
		}

		if _, exists := buttons[key]; !exists {
			buttons[key] = &widget.Clickable{}
		}

		for buttons[key].Clicked(gtx) {
			fmt.Println("Button pressed:", value)
		}

		return material.Button(th, buttons[key], value).Layout(gtx)
	})
}
