package manager

import (
    "errors"
)

func getTaskName(taskIndex []int) (taskName []string, err error) {
    list, err := ListTask()
    if err != nil {
        return
    }
    taskName = make([]string, len(taskIndex))
    for i, v := range taskIndex {
        if v < len(list) && v > -1 {
            taskName[i] = list[v]
        } else {
            err = errors.New("task index not exists")
            return
        }
    }
    return
}

func DoTask(taskIndex []int) (err error) {
    taskName, err := getTaskName(taskIndex)
    if err != nil {
        return
    }
    err = markTaskAsDone(taskName)
    return
}
