package main

import "fmt"

func main() {
    x := "hello"
    for _, x := range x {
        fmt.Printf("origin x: %c\n", x)
        x := x + 'A' - 'a'
        fmt.Printf("x: %c\n", x)
    }
}
