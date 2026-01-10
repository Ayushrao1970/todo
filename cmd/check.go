package cmd

import (
	"log"
	"strconv"

	"github.com/Ayushrao1970/todo/api"
	"github.com/spf13/cobra"
)

var check = &cobra.Command{
	Use:   "done",
	Short: "check it to mark it done",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		task := &api.TodoList{}
		task.File = "db.json"
		task.Load()
		i, err := strconv.Atoi(args[0])
		if err != nil {
			log.Fatalf("failed in conversion %v", err)
		}
		task.Check(i)
		task.Get()

	},
}
var uncheck = &cobra.Command{
	Use:   "pending",
	Short: "check it to mark it done",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		task := &api.TodoList{}
		task.Load()
		task.File = "db.json"
		i, err := strconv.Atoi(args[0])
		if err != nil {
			log.Fatalf("failed in conversion %v", err)
		}
		task.UnCheck(i)
		task.Get()

	},
}
