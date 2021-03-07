package main

import (
	"fmt"
	"github.com/jroimartin/gocui"
	"github.com/recep/todo-cli-app/internal/gui"
	"github.com/recep/todo-cli-app/internal/utils"
	"log"
)

func main() {
	g, err := gocui.NewGui(gocui.OutputNormal)
	if err != nil {
		log.Fatal(err)
	}
	defer g.Close()

	// Active Layout Highlighting
	g.Highlight = true
	g.SelFgColor = gocui.ColorGreen

	// Setup layout and keybindings
	g.SetManagerFunc(gui.Layout)
	if err := gui.Keybindings(g); err != nil {
		log.Fatal(err)
	}

	// Get todos from the file
	todos, err := utils.ReadData("./storage/todos.txt")
	if err != nil {
		log.Fatalln(err)
	}

	// Write todos
	g.Update(func(g *gocui.Gui) error {
		tView, err := g.View("todos")
		if err != nil {
			return err
		}

		for _, todo := range todos {
			fmt.Fprintln(tView, todo)
		}
		return nil
	})

	// MainLoop
	if err := g.MainLoop(); err != nil && err != gocui.ErrQuit {
		log.Fatal(err)
	}
}
