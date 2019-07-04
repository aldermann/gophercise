package manager

func AddTask (taskName []string) error {
    s := make([]interface{}, len(taskName))
    for i, v := range taskName {
        s[i] = v
    }
    _, err := client.SAdd("taskList", s...).Result()
    return err
}

func DeleteTask (taskId []int) error {
    taskName, err := getTaskName(taskId)
    if err != nil {
        return err
    }
    s := make([]interface{}, len(taskName))
    for i, v := range taskName {
        s[i] = v
    }
    _, err = client.SRem("taskList", s...).Result()
    return err
}

func ListTask () ([]string, error) {
    s, err := client.SMembers("taskList").Result()
    if err != nil {
        return []string{}, err
    }
    res := make([]string, len(s))
    for i, v := range s {
        res[i] = v
    }
    return res, nil
}