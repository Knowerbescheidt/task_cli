/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"clitool/db"
	"fmt"

	"github.com/spf13/cobra"
)

// rmCmd represents the rm command
var rmCmd = &cobra.Command{
	Use:   "rm",
	Short: "Remove task",
	Long:  `Removes a task from the task list`,
	Run: func(cmd *cobra.Command, args []string) {
		ids := argsToIds(args)
		fmt.Printf("rm called, ids to rm %d", ids)
		tasks, err := db.AllTasks()
		if err != nil {
			fmt.Println(err)
			return
		}
		for _, id := range ids {
			if id <= 0 || id > len(tasks) {
				fmt.Println("Invalid task number:", id)
				continue
			}
			task := tasks[id-1]
			fmt.Printf("PReparing to delete %s", task.Value)
			err := db.DeleteTask(task.Key)
			if err != nil {
				fmt.Printf("Error occured %s", err)
			}
		}
	},
}

func init() {
	rootCmd.AddCommand(rmCmd)
}
