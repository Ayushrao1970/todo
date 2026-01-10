package cmd

import (
	"log"
	"strconv"

	"github.com/Ayushrao1970/todo/api"
	"github.com/spf13/cobra"
)

var update = &cobra.Command{
	Use:   "update",
	Short: "update the todo by giving index of it",
	Args:  cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		task := &api.TodoList{}
		task.File = "db.json"
		task.Load()
		i, err := strconv.Atoi(args[0])
		if err != nil {
			log.Fatalf("unable to convert the string to int %v", err)
		}
		task.Update(i, args[1])

	},
}
