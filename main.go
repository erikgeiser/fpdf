package main

import (
  "github.com/leaanthony/mewn"
  "github.com/wailsapp/wails"
)

func basic() string {
  return "Hello World!"
}

func main() {

	js := mewn.String("./frontend/build/main.js")
	css := mewn.String("./frontend/build/main.css")

  app := wails.CreateApp(&wails.AppConfig{
    Width:  1024,
    Height: 768,
    Title:  "fpdf",
    JS:     js,
    CSS:    css,
    Colour: "#131313",
  })
  app.Bind(basic)
  app.Run()
}