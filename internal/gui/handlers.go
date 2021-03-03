package gui

import (
	"github.com/jroimartin/gocui"
	"log"
)

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

func getLine(g *gocui.Gui, v *gocui.View) error {
	_, cy := v.Cursor()
	str, err := v.Line(cy)
	if err != nil {
		log.Fatalln(err)
	}

	if str == "Exit" {
		return gocui.ErrQuit
	}

	maxX, maxY := g.Size()
	if v, err := g.SetView("msg", maxX/2-30, maxY/2, maxX/2+30, maxY/2+2); err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}

		v.Editable = true

		if _, err := g.SetCurrentView("msg"); err != nil {
			return err
		}
	}

	return nil
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

