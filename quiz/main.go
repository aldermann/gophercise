package main

import (
    "bufio"
    "fmt"
    "os"
    "strings"
    "time"
)

type problem struct {
    question, answer string
}

var score = 0
var r Inputter
var inputter = bufio.NewReader(os.Stdin)

func CheckError(err error) {
    if err != nil {
        panic(err)
    }
}

func main() {
    flagParse()
    if filePath == "mock" {
        r = MockInputter{}
    } else {
        r = CSVInputter{filePath: filePath, randomized: !norand}
    }
    problemList := r.Input()
    if questionCount > len(problemList) {
        questionCount = len(problemList)
    }
    answeredChan := make(chan bool)
    defer close(answeredChan)
    timer := time.NewTimer(time.Second * time.Duration(timeOut))
    loop:
    for i, problem := range problemList {
        if questionCount != 0 && i == questionCount {
            break loop
        }
        timer.Reset(time.Second * time.Duration(timeOut))

        go func () {
            fmt.Printf("%s = ?\n", problem.question)
            input, _ := inputter.ReadString('\n')
            input = strings.TrimSpace(input)
            if input == problem.answer {
                score++
            }
            fmt.Println()
            answeredChan <- true
        }()

        select {
        case <-timer.C:
            fmt.Println("Time's up")
            break loop
        case <-answeredChan:
            continue loop
        }
    }
    fmt.Printf("You have answered correctly %d/%d problems\n", score, questionCount)
}
