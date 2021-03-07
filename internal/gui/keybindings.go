package gui

import (
	"github.com/jroimartin/gocui"
	"log"
)

// Keybindings
func Keybindings(g *gocui.Gui) error {
	// Key Arrow Down menu and todoView
	if err := g.SetKeybinding("", gocui.KeyArrowDown, gocui.ModNone, cursorDown); err != nil {
		return err
	}

	// Key Arrow Up menu and todoView
	if err := g.SetKeybinding("", gocui.KeyArrowUp, gocui.ModNone, cursorUp); err != nil {
		return err
	}

	// Edit
	if err := g.SetKeybinding("menu", gocui.KeyEnter, gocui.ModNone, getLineMenu); err != nil {
		return err
	}

	// Save
	if err := g.SetKeybinding("msg", gocui.KeyEnter, gocui.ModNone, save); err != nil {
		return err
	}

	// Get options
	if err := g.SetKeybinding("todos", gocui.KeyEnter, gocui.ModNone, getOptions); err != nil {
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
