package main

import (
    "fmt"
    "os"
)

func main() {

    for idx, arg := range os.Args[:] {
        fmt.Printf("idx=%d arg=%s \n", idx, arg)
    }

}
