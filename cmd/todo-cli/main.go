package main

import (
	"fmt"
	"github.com/jroimartin/gocui"
	"log"
)

func main() {
	g, err := gocui.NewGui(gocui.OutputNormal)
	if err != nil {
		log.Fatal(err)
	}
	defer g.Close()

	g.SetManagerFunc(layout)

	if err := keybindings(g); err != nil {
		log.Fatal(err)
	}

	// MainLoop
	if err := g.MainLoop(); err != nil && err != gocui.ErrQuit {
		log.Fatal(err)
	}
}

func layout(g *gocui.Gui) error {
	maxX, maxY := g.Size()

	if v, err := g.SetView("info", 0, 0, maxX-1, maxY/9); err != nil {
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
		fmt.Fprintln(v, "Add Task")
		fmt.Fprintln(v, "Exit")
	}

	if v, err := g.SetView("todos", maxX/5+1, maxY/9+1, maxX-20, maxY-1); err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}

		v.Title = "TODOs"
		v.Highlight = true
		v.SelBgColor = gocui.ColorCyan
		v.SelFgColor = gocui.ColorWhite
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
	// Key Arrow Down
	if err := g.SetKeybinding("menu", gocui.KeyArrowDown, gocui.ModNone, cursorDown); err != nil {
		return err
	}

	// Key Arrow Up
	if err := g.SetKeybinding("menu", gocui.KeyArrowUp, gocui.ModNone, cursorUp); err != nil {
		return err
	}

	// Next view
	if err := g.SetKeybinding("", gocui.KeyTab, gocui.ModNone, nextView); err != nil {
		return err
	}

	if err := g.SetKeybinding("menu", gocui.KeyTab, gocui.ModNone, nextView); err != nil {
		return err
	}

	if err := g.SetKeybinding("todos", gocui.KeyTab, gocui.ModNone, nextView); err != nil {
		return err
	}

	// Quit
	if err := g.SetKeybinding("", gocui.KeyCtrlC, gocui.ModNone, quit); err != nil {
		log.Fatalln(err)
	}

	return nil
}

func nextView(g *gocui.Gui, v *gocui.View) error {
	if v == nil || v.Name() != "menu" {
		_, err := g.SetCurrentView("menu")
		g.Cursor = false
		return err
	}
	_, err := g.SetCurrentView("todos")
	g.Cursor = true

	return err
}

func cursorUp(g *gocui.Gui, v *gocui.View) error {
	if v != nil {
		ox, oy := v.Origin()
		cx, cy := v.Cursor()
		if err := v.SetCursor(cx, cy-1); err != nil && oy > 0 {
			if err := v.SetOrigin(ox, oy-1); err != nil {
				return err
			}
		}
	}
	return nil
}

func cursorDown(g *gocui.Gui, v *gocui.View) error {
	if v != nil {
		cx, cy := v.Cursor()
		if err := v.SetCursor(cx, cy+1); err != nil {
			ox, oy := v.Origin()
			if err := v.SetOrigin(ox, oy+1); err != nil {
				return err
			}
		}
	}
	return nil
}

func quit(gui *gocui.Gui, view *gocui.View) error {
	return gocui.ErrQuit
}
