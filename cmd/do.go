/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"clitool/db"
	"fmt"

	"github.com/spf13/cobra"
)

// doCmd represents the do command
var doCmd = &cobra.Command{
	Use:   "do",
	Short: "Marks a task as done",
	Long:  `A longer description`,
	Run: func(cmd *cobra.Command, args []string) {
		ids := argsToIds(args)
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
			fmt.Printf("PReparing to update %s", task.Value)
			err := db.UpdateTask(task)
			if err != nil {
				fmt.Printf("Failed to mark \"%d\" as completed. Error: %s \n", id, err)
			} else {
				fmt.Printf("Marked \"%d\" as completed\n", id)
			}
		}

	},
}

func init() {
	rootCmd.AddCommand(doCmd)
}
