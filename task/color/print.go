package color

import (
    "fmt"
)

func Println(color string, args ...interface{}) {
    fmt.Print(color)
    fmt.Print(args...)
    fmt.Println(Reset)
}

func Printf(color string, format string, args ...interface{}) {
    fmt.Print(color)
    fmt.Printf(format, args...)
    fmt.Print(Reset)
}

func Print(color string, args ...interface{}) {
    fmt.Print(color)
    fmt.Print(args...)
    fmt.Print(Reset)
}
