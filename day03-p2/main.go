package main

import (
    "fmt"
    "io"
)

func rating(values []uint, shift int, reverse bool) uint {
    if len(values) == 0 {
        return 0
    }
    if len(values) == 1 {
        return values[0]
    }

    var ones, count uint
    for _, v := range values {
        if (v>>shift)&1 == 1 {
            ones++
        }
        count++
        // fmt.Printf("%012b %d %d\n", v, ones, count)
    }

    var ref uint
    if 2*ones >= count {
        ref = 1
    }
    if reverse {
        ref ^= 1
    }

    var keep []uint
    for _, v := range values {
        if (v>>shift)&1 == ref {
            keep = append(keep, v)
        }
        // fmt.Printf("%012b %012b %d\n", v, v >> shift, ref)
    }
    return rating(keep, shift-1, reverse)
}

func main() {
    var line string
    var values []uint
    var length int
    for {
        _, err := fmt.Scanln(&line)
        if err == io.EOF {
            break
        }
        if err != nil {
            panic(err)
        }
        length = len(line)

        var val uint
        fmt.Sscanf(line, "%b", &val)
        values = append(values, val)
    }
    // fmt.Println(values)

    oxygen := rating(values, length-1, false)
    co2 := rating(values, length-1, true)
    // fmt.Println(oxygen, co2)
    fmt.Println(oxygen * co2)
}
