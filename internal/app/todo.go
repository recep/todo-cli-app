package app

import (
	"fmt"
	"github.com/jroimartin/gocui"
	"github.com/recep/todo-cli-app/internal/utils"
)

func AddTodo(str string, v *gocui.View) error {
	v.SelBgColor = gocui.ColorGreen
	v.SelFgColor = gocui.ColorBlack
	v.Wrap = true
	fmt.Fprintln(v, str)

	// Save task to txt file
	if err := utils.SaveDataToFile(str, "./storage/todos.txt"); err != nil {
		return err
	}

	return nil
}

func CompleteTodo(todo string) error {
	// Get Todos
	todos, err := utils.ReadData("./storage/todos.txt")
	if err != nil {
		return err
	}

	// Delete task from the slice
	for i, v := range todos {
		if v == todo {
			todos = append(todos[:i], todos[i+1:]...)
		}
	}

	if err := utils.UpdateFile(todos, "./storage/todos.txt"); err != nil {
		return err
	}

	return nil
}
