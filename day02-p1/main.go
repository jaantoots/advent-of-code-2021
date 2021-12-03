package main

import (
    "fmt"
    "io"
)

func main() {
    var cmd string
    var pos, depth, x int
    for {
        _, err := fmt.Scanln(&cmd, &x)
        if err == io.EOF {
            break
        }
        if err != nil {
            panic(err)
        }

        switch cmd {
        case "forward":
            pos += x
        case "down":
            depth += x
        case "up":
            depth -= x
        default:
            panic(fmt.Sprintf("unknown command: %s", cmd))
        }
        // fmt.Printf("%d %d %s %d\n", pos, depth, cmd, x)
    }
    fmt.Println(pos*depth)
}
