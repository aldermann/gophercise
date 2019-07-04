package game

import (
    "fmt"
)

func (b Board) Print() {
    fmt.Print(" x")
    for i := 0; i < b.Col; i++ {
        fmt.Print(" ", i)
    }
    fmt.Println()
    fmt.Println("y")
    for i, row := range b.container {
        fmt.Print(i, " ")
        for _, cell := range row {
            if cell == X {
                fmt.Print(" X")
            } else if cell == O {
                fmt.Print(" O")
            } else {
                fmt.Print(" _")
            }
        }
        fmt.Println()
    }
    fmt.Println()
}
