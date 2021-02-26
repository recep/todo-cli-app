package main

import (
	"fmt"
	"github.com/jroimartin/gocui"
	"log"
)

func main() {
	g, err := gocui.NewGui(gocui.OutputNormal)
	if err != nil {
		log.Fatalln(err)
	}
	defer g.Close()

	g.Cursor = true

	g.SetManagerFunc(layout)

	// quit
	if err := g.SetKeybinding("", gocui.KeyCtrlC, gocui.ModNone, quit); err != nil {
		log.Fatalln(err)
	}

	// MainLoop
	if err := g.MainLoop(); err != nil && err != gocui.ErrQuit {
		log.Panicln(err)
	}
}

func layout(g *gocui.Gui) error {
	maxX, maxY := g.Size()

	if v, err := g.SetView("appInfo", 0, 0, maxX-1, maxY/9); err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}

		v.Title = "App Info"
		v.FgColor = gocui.ColorCyan
		v.BgColor = gocui.ColorBlue
		v.Wrap = true
		v.Autoscroll = true
		fmt.Fprint(v, "GO TODO CLI APP\nv.0.0.1")
	}

	if v, err := g.SetView("Menu", 0, maxY/9+1, maxX/5, maxY/3); err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}

		v.Title = "Menu"
		v.Wrap = true
		fmt.Fprintln(v, "Add Todo")
		fmt.Fprintln(v, "Delete Todo")
	}

	if v, err := g.SetView("TODOs", maxX/5+1, maxY/9+1, maxX-1, maxY-1); err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}

		v.Title = "TODOs"
		v.Wrap = true
	}

	return nil
}

func quit(gui *gocui.Gui, view *gocui.View) error {
	return gocui.ErrQuit
}
