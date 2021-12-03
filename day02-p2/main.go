package main

import (
    "fmt"
    "io"
)

func main() {
    var cmd string
    var pos, depth, aim, x int
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
            depth += aim * x
        case "down":
            aim += x
        case "up":
            aim -= x
        default:
            panic(fmt.Sprintf("unknown command: %s", cmd))
        }
        // fmt.Printf("%d %d %d %s %d\n", pos, depth, aim, cmd, x)
    }
    fmt.Println(pos*depth)
}
