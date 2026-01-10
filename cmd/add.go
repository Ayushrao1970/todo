package cmd

import (
	"github.com/Ayushrao1970/todo/api"
	"github.com/spf13/cobra"
)

var add = &cobra.Command{
	Use:   "add",
	Short: "add the todo to the list",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		task := &api.TodoList{}
		task.File = "db.json"
		task.Load()
		task.Add(args[0])
		task.Get()

	},
}
