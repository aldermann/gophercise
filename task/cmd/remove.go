/*
Copyright © 2019 NAME HERE <EMAIL ADDRESS>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package cmd

import (
	"fmt"
	"strconv"

	"github.com/spf13/cobra"
	"task/color"
	"task/manager"
)

// removeCmd represents the remove command
var removeCmd = &cobra.Command{
	Use:   "remove",
	Short: "Remove a task",
	Long: `Remove a task from your active task list, without moving it into the completed list. 
Input your task index(s), which is given when listing the tasks.
Syntax:
	task remove [task-index-1] [task-index-2] (with task indexes shown in the task list)
`,
	Run: func(cmd *cobra.Command, args []string) {
		taskId := make ([]int, len(args))
		for i, v := range args {
			var err error
			taskId[i], err = strconv.Atoi(v)
			if err != nil {
				color.Println(color.Red, "your input is not valid index (number)")
			}
		}
		err := manager.DeleteTask(taskId)
		if err != nil {
			color.Println(color.Red, err)
		}
		fmt.Println("Your current tasks")
		printTask()
	},
}

func init() {
	rootCmd.AddCommand(removeCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// removeCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// removeCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
