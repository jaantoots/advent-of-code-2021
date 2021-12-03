package main

import (
    "fmt"
    "io"
)

func main() {
    var line string
    for {
        _, err := fmt.Scanln(&line)
        if err == io.EOF {
            break
        }
        if err != nil {
            panic(err)
        }
    }
}
