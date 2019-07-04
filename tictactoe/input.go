package main

import (
    "bufio"
    "errors"
    "fmt"
    "os"
    "strconv"
    "strings"
)

var rd *bufio.Reader

func GetReader() *bufio.Reader {
    if rd == nil {
        rd = bufio.NewReader(os.Stdin)
    }
    return rd
}

func ReadInt (prompt string) (int, error) {
    rd := GetReader()
    fmt.Print(prompt, " ")
    inp, err := rd.ReadString('\n')
    if err != nil {
        panic(err)
    }
    inp = strings.TrimSpace(inp)
    res, err := strconv.Atoi(inp)
    if err != nil {
        return 0, errors.New("invalid input. please input a valid number")
    }
    return res, nil
}

func ReadCoordinate() (int, int) {
    x, y := 0, 0
    var err error
    fmt.Println("Input the coordinate x, y: ")
    for {
        x, err = ReadInt("Input x:")
        if err != nil{
            fmt.Println(err)
        } else {
            break
        }
    }
    for {
        y, err = ReadInt("Input y:")
        if err != nil{
            fmt.Println(err)
        } else {
            break
        }
    }
    return x, y
}
