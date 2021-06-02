package main

import "fmt"

func main() {


    var x uint8 = 1<<1 | 1<<5
    var y uint8 = 1<<1 | 1<<2

    fmt.Printf("%08b\n", x)
    fmt.Printf("%08b\n", y)

    fmt.Printf("%08b\n", x&y)
    fmt.Printf("%08b\n", x|y)
    fmt.Printf("%08b\n", x^y)
    fmt.Printf("%08b\n", x&^y) // 按y的位对x进行清零



    fmt.Printf("%08b\n", x<<1)
    fmt.Printf("%08b\n", x>>1)

    whichBitTrue(x<<1)
    whichBitTrue(x>>1)

    var z int8 = -10
    fmt.Printf("%08b\n", z)
    fmt.Printf("%08b\n", z>>1)
}

func whichBitTrue(n uint8) {
    for i := uint(0); i< 8; i++ {
        if n&(1<<i) != 0 {
            fmt.Println(i)
        }
    }
}