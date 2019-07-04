package utils

import (
    "fmt"
    "os"
)

func CheckError (err error, fatal bool) {
    if err != nil {
        fmt.Println(err)
        if fatal {
            os.Exit(1)
        }
    }
}