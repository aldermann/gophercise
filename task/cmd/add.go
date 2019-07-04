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
    "strings"

    "github.com/spf13/cobra"
    "task/color"
    "task/manager"
)

// addCmd represents the add command
var addCmd = &cobra.Command{
    Use:   "add",
    Short: "Add new tasks to your list",
    Long: `Add new tasks to your list. Syntax:
task add task-name (no quote needed)
task add "[task-name-1]" "[task-name-2]" ...`,
    Args: cobra.MinimumNArgs(1),
    Run: func(cmd *cobra.Command, args []string) {
        taskList := make([]string, 0, len(args))
        currentTaskName := ""
        for id, v := range args {
            if strings.Contains(v, " ") {
                if id != 0 {
                    taskList = append(taskList, currentTaskName)
                }
                taskList = append(taskList, v)
                currentTaskName = ""
            } else {

                if currentTaskName != "" {
                    currentTaskName = currentTaskName + " " + v
                } else {
                    currentTaskName = v
                }
            }
        }
        if currentTaskName != "" {
            taskList = append(taskList, currentTaskName)
        }
        err := manager.AddTask(taskList)
        if err != nil {
            color.Println(color.Red, err)
            return
        }
        color.Println(color.Green, "Added task")
        fmt.Println("Here are your tasks")
        printTask()
    },
}

func init() {
    rootCmd.AddCommand(addCmd)

    // Here you will define your flags and configuration settings.

    // Cobra supports Persistent Flags which will work for this command
    // and all subcommands, e.g.:
    // addCmd.PersistentFlags().String("foo", "", "A help for foo")

    // Cobra supports local flags which will only run when this command
    // is called directly, e.g.:
    // addCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
