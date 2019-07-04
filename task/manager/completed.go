package manager
func ListFinishedTask () ([]string, error) {
    s, err := client.SMembers("completedTask").Result()
    if err != nil {
        return []string{}, err
    }
    res := make([]string, len(s))
    for i, v := range s {
        res[i] = v
    }
    return res, nil
}

func markTaskAsDone (taskName []string) error {
    s := make([]interface{}, len(taskName))
    for i, v := range taskName {
        s[i] = v
    }
    _, err := client.SRem("taskList", s...).Result()
    if err != nil {
        return err
    }
    _, err = client.SAdd("completedTask", s...).Result()
    return err
}