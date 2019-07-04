package utils

import (
    "flag"
)

func ParseFlag() string {
    filePath := flag.String("f", "null", "Specify the .yml file to parse")
    flag.Parse()
    return *filePath
}
