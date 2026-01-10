package cmd

import (
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{ //create struct pointer
	Use:   "task",
	Short: "Add task to a manager",
	Long:  "This add the tasks in task manager",
}

func Execute() {
	cobra.CheckErr(rootCmd.Execute()) //check for error if while createing the struct
}

func init() {
	rootCmd.AddCommand(list)
	rootCmd.AddCommand(add) // adding subcommand to it
	rootCmd.AddCommand(delete)
	rootCmd.AddCommand(update)
	rootCmd.AddCommand(check)
	rootCmd.AddCommand(uncheck)

}
