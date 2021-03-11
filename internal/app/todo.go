package app

import (
	"encoding/json"
	"fmt"
	"github.com/jroimartin/gocui"
	"github.com/recep/todo-cli-app/internal/utils"
)

type Todo struct {
	Task      string `json:"task"`
	Completed bool   `json:"completed"`
}

func AddTodo(str string, v *gocui.View) error {
	v.SelBgColor = gocui.ColorGreen
	v.SelFgColor = gocui.ColorBlack
	v.Wrap = true
	fmt.Fprintln(v, str)

	// Create new model
	todo := Todo{
		Task:      str,
		Completed: false,
	}

	todos, err := GetAllTodos()
	if err != nil {
		return err
	}

	todos = append(todos, &todo)

	bytes, err := json.MarshalIndent(todos, "", " ")
	if err != nil {
		return err
	}

	if err := utils.SaveDataToFile(bytes, "./storage/tasks.json"); err != nil {
		return err
	}

	return nil
}

func CompleteTask(todo string) error {
	// Get all todos from the json file
	todos, err := GetAllTodos()
	if err != nil {
		return err
	}

	for _, t := range todos {
		if t.Task == todo {
			t.Completed = true
		}
	}

	bytes, err := json.MarshalIndent(todos, "", " ")
	if err != nil {
		return err
	}

	if err := utils.SaveDataToFile(bytes, "./storage/tasks.json"); err != nil {
		return err
	}

	return nil
}

func GetAllTodos() ([]*Todo, error) {
	// Read data from the file
	data, err := utils.ReadData("./storage/tasks.json")
	if err != nil {
		return nil, err
	}

	var todos []*Todo
	if err := json.Unmarshal(data, &todos); err != nil {
		return nil, err
	}

	return todos, err
}
