package main

import (
    "fmt"
    "math"
)

func main() {

    for x := 0; x < 8; x++ {
        fmt.Printf("x = %d ex = %8.3f\n", x, math.Exp(float64(x)))
        fmt.Printf("x = %d ex = %8.3e\n", x, math.Exp(float64(x)))
        fmt.Printf("x = %d ex = %g\n", x, math.Exp(float64(x)))
    }

    var f float32 = 1<<24
    fmt.Println(f == f+1)
    fmt.Printf("%g %g\n", f, f+1)

    var z float64
    fmt.Println(z, -z, 1/z, -1/z, z/z)
    fmt.Println(math.IsNaN(z/z))   // true
    // nan与nan比较总是不成立
    fmt.Println(z/z == math.NaN()) // false

    fmt.Println(math.IsInf(-1/z, -1)) // true
    fmt.Println(math.IsInf(1/z, 1))  // true

    fmt.Println(math.IsInf(z, 1))    // false
}
