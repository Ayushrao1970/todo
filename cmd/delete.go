package cmd

import (
	"log"
	"strconv"

	"github.com/Ayushrao1970/todo/api"
	"github.com/spf13/cobra"
)

var delete = &cobra.Command{
	Use:   "delete",
	Short: "delete the task",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		task := &api.TodoList{}
		task.File = "db.json"
		task.Load()
		i, err := strconv.Atoi(args[0])
		if err != nil {
			log.Fatalf("unable convert string to int %v", err)
		}
		task.Delete(i)
		task.Get()

	},
}
