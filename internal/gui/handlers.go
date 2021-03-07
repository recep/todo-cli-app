package gui

import (
	"github.com/jroimartin/gocui"
	"github.com/recep/todo-cli-app/internal/app"
	"log"
)

func nextView(g *gocui.Gui, v *gocui.View) error {
	if v == nil || v.Name() != "menu" {
		_, err := g.SetCurrentView("menu")
		g.Cursor = false
		return err
	}
	_, err := g.SetCurrentView("todos")

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
	if msgV, err := g.SetView("msg", maxX/2-30, maxY/2, maxX/2+30, maxY/2+2); err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}

		msgV.Editable = true

		if _, err := g.SetCurrentView("msg"); err != nil {
			return err
		}
	}

	return nil
}

// Task saving handler
func save(g *gocui.Gui, v *gocui.View) error {
	// Get string of the line
	_, cy := v.Cursor()
	str, err := v.Line(cy)
	if err != nil {
		return err
	}

	// Delete view
	if err := g.DeleteView("msg"); err != nil {
		return err
	}

	// Change current view to menu
	if _, err := g.SetCurrentView("menu"); err != nil {
		return err
	}

	// Get todos view
	todos, err := g.View("todos")
	if err != nil {
		return err
	}

	// Append new task to file
	if err := app.AddTodo(str, todos); err != nil {
		return err
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
