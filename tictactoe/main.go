package main

import (
    "fmt"
    "math"
    "math/rand"
    "time"

    "tictactoe/algorithm"
    "tictactoe/game"
)

func clearScreen() {
    fmt.Print("\033[2J")
}

func estimateDepth(moveLeft int) int {
    x := 1.0
    i := 0
    threshold := math.Pow(1e8, 1)
    for i = moveLeft; i > 0; i-- {
        x = x * float64(i)
        if threshold <= x {
            break
        }
    }
    return moveLeft - i
}

func play(r, c int) {
    currentBoard := game.MakeBoard(r, c)
    moveLeft := r * c
    clearScreen()
    currentBoard.Print()
    for {
        for {
            x, y := ReadCoordinate()
            err := currentBoard.MarkMove(x, y, game.X)
            if err == nil {
                break
            } else {
                fmt.Println(err)
            }
        }
        moveLeft--
        clearScreen()
        currentBoard.Print()
        if currentBoard.GetValue() == game.Lose {
            fmt.Println("You won!!!!")
            break
        }
        if currentBoard.IsFull() {
            fmt.Println("Draw")
            break
        }
        start := time.Now()
        _, nextMove := algorithm.CalculateBestMove(currentBoard, int(-1e9), int(1e9), estimateDepth(moveLeft), true, true)
        currentBoard = nextMove
        moveLeft--
        t := time.Now()
        elapsed := t.Sub(start)
        clearScreen()
        fmt.Printf("The operation took %f s\n", float64(elapsed)/float64(time.Second))
        currentBoard.Print()
        if currentBoard.GetValue() == game.Win {
            fmt.Println("You lost")
            break
        }
        if currentBoard.IsFull() {
            fmt.Println("Draw")
            break
        }
    }
}

func main() {
    rand.Seed(time.Now().UnixNano())
    row, col := 0, 0
    var err error
    for {
        row, err = ReadInt("how many row?:")
        if err != nil {
            fmt.Println(err)
        } else {
            break
        }
    }
    for {
        col, err = ReadInt("how many column?:")
        if err != nil {
            fmt.Println(err)
        } else {
            break
        }
    }
    play(row, col)
    //b := game.MakeBoard(3, 3)
    //b.MarkMove(1, 0, game.X)
    //b.MarkMove(0, 0, game.X)
    ////b.MarkMove(1, 2, game.X)
    //b.MarkMove(1, 2, game.O)
    //b.MarkMove(2, 2, game.O)
    //v, m := algorithm.CalculateBestMove(b, false, true)
    //fmt.Println(v)
    //m.Print()
    ////fmt.Println(b.GetValue())
}
