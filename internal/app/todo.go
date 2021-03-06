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
	if err := utils.SaveDataToFile(str); err != nil {
		return err
	}

	return nil
}
