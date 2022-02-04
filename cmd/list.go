/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"clitool/db"
	"fmt"

	"github.com/spf13/cobra"
)

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List all active tasks in you tasklist",
	Long:  `A longer more detaile description`,
	Run: func(cmd *cobra.Command, args []string) {
		tasks, err := db.AllTasks()
		if err != nil {
			fmt.Println("Error occured during listing all the tasks")
		}
		for i, task := range tasks {
			fmt.Printf("%d. %s %s\n", i+1, task.Value, task.Status)
		}
	},
}

func init() {
	rootCmd.AddCommand(listCmd)
}
