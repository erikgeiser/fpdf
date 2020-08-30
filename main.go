package main

import (
	"fmt"

	"github.com/leaanthony/mewn"
	"github.com/wailsapp/wails"
)

func basic() string {
	return "Hello World!"
}

func main() {
	html := mewn.String("./frontend/public/index.html")
	js := mewn.String("./frontend/public/build/bundle.js")
	css := mewn.String("./frontend/public/build/bundle.css")

	app := wails.CreateApp(&wails.AppConfig{
		Width:            1024,
		Height:           768,
		Title:            "fpdf",
		HTML:             html,
		JS:               js,
		CSS:              css,
		Resizable:        true,
		DisableInspector: true,
	})

	app.Bind(basic)

	err := app.Run()
	if err != nil {
		fmt.Printf("Error: %v", err)
	}
}
