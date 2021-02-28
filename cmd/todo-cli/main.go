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

	err = keybindings(g)
	if err != nil {
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

	if v, err := g.SetView("menu", 0, maxY/9+1, maxX/5, maxY/3); err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}

		v.Title = "Menu"
		v.Highlight = true
		v.SelBgColor = gocui.ColorGreen
		v.SelFgColor = gocui.ColorBlack
		fmt.Fprintln(v, "[1] Add Todo")
		fmt.Fprintln(v, "[2] Complete Todo")
		fmt.Fprintln(v, "[3] Delete Todo")
	}

	if v, err := g.SetView("todos", maxX/5+1, maxY/9+1, maxX-20, maxY-1); err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}

		v.Title = "TODOs"
		v.Wrap = true
	}

	if v, err := g.SetView("completed", maxX/2+1, maxY/9+1, maxX-1, maxY-1); err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}

		v.Title = "Completed TODOs"
		v.Wrap = true
	}

	return nil
}

func keybindings(g *gocui.Gui) error {
	// quit
	if err := g.SetKeybinding("", gocui.KeyCtrlC, gocui.ModNone, quit); err != nil {
		log.Fatalln(err)
	}

	return nil
}

func quit(gui *gocui.Gui, view *gocui.View) error {
	return gocui.ErrQuit
}
