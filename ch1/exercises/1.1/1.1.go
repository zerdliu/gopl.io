package main

import (
    "fmt"
    "os"
)

func main() {
    sep, s := " ", ""
    for _, arg := range os.Args[:] {
        s += sep + arg

    }
    fmt.Println(s)
}
