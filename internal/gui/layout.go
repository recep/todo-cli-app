package gui

import (
	"fmt"
	"github.com/jroimartin/gocui"
)

// layout function
func Layout(g *gocui.Gui) error {
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
		fmt.Fprint(v, "GO TODO CLI APP\nv.0.5.0")
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

	if v, err := g.SetView("keyshortcuts",0,maxY/3+1,maxX/5,maxY/2) ; err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}

			v.Title = "Key Shortcuts"

			fmt.Fprintln(v,"TAB - Switch tabs")
			fmt.Fprintln(v,"CTRL+S - Complete Current Todo")
			fmt.Fprintln(v,"CTRL+C - Exit")
	}

	return nil
}
