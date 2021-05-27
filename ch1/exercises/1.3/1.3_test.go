package main

import (
    "fmt"
    "os"
    "strings"
    "testing"
)



func echo() {
    sep, s := " ", ""
    for _, arg := range os.Args[:] {
        s += sep + arg

    }
    fmt.Println(s)
}

func echoQuick(){
    fmt.Println(strings.Join(os.Args[:], " "))
}

func BenchmarkEcho(b *testing.B) {
    for i := 0; i < b.N; i++ {
        echo()
    }
}
