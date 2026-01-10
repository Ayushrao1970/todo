package cmd

import (
	"github.com/Ayushrao1970/todo/api"
	"github.com/spf13/cobra"
)

var list = &cobra.Command{
	Use:   "list",
	Short: "lists the todos",
	Run: func(cmd *cobra.Command, args []string) {
		task := &api.TodoList{}
		task.File = "db.json"
		task.Load()
		task.Get()

	},
}
