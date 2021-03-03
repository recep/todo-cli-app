package main

import (
	"github.com/jroimartin/gocui"
	"github.com/recep/todo-cli-app/internal/gui"
	"log"
)

var index int

func main() {
	g, err := gocui.NewGui(gocui.OutputNormal)
	if err != nil {
		log.Fatal(err)
	}
	defer g.Close()

	g.SetManagerFunc(gui.Layout)

	if err := gui.Keybindings(g); err != nil {
		log.Fatal(err)
	}

	// MainLoop
	if err := g.MainLoop(); err != nil && err != gocui.ErrQuit {
		log.Fatal(err)
	}
}
