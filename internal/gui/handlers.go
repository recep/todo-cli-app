package gui

import (
	"fmt"
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
	_, err := g.SetCurrentView("tasks")
	g.Cursor = false

	return err
}

func getLineMenu(g *gocui.Gui, v *gocui.View) error {
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
		g.Cursor = true

		if _, err := g.SetCurrentView("msg"); err != nil {
			return err
		}
	}

	return nil
}

func getOptions(g *gocui.Gui, v *gocui.View) error {
	maxX, maxY := g.Size()
	if oV, err := g.SetView("options", maxX/5+15, maxY/9+5, maxX/3+10, maxY/9+10); err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}

		oV.Highlight = true
		oV.SelBgColor = gocui.ColorGreen
		oV.SelFgColor = gocui.ColorBlack
		oV.Wrap = true
		fmt.Fprintln(oV, "Complete")
		//fmt.Fprintln(oV, "Change Color")

		if _, err := g.SetCurrentView("options"); err != nil {
			return err
		}
	}

	return nil
}

func completeTask(g *gocui.Gui, v *gocui.View) error {
	// Get string of the line
	_, cy := v.Cursor()
	task, err := v.Line(cy)
	if err != nil {
		return err
	}

	// Complete task
	if err := app.CompleteTask(task); err != nil {
		return err
	}

	// Refresh view
	if err := app.RefreshTasksView(v); err != nil {
		return err
	}

	// Get completedView
	gV, err := g.View("completed")
	if err != nil {
		return err
	}

	if err := app.RefreshCompletedView(gV); err != nil {
		return err
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
	g.Cursor = false

	// Change current view to menu
	if _, err := g.SetCurrentView("menu"); err != nil {
		return err
	}

	// Get todos view
	todos, err := g.View("tasks")
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
