package main

import (
    "fmt"
    "math/cmplx"
)

func main() {

    var x complex128 = complex(1,2)
    var y complex128 = complex(3,4)

    fmt.Println(x*y)
    fmt.Println(real(x*y), imag(x*y))

    fmt.Println(cmplx.Sqrt(-1))
}
