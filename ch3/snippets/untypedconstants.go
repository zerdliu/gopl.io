package main

import (
    "fmt"
    "math"
)

func main() {
    var x float32 = math.Pi
    var y float64 = math.Pi
    var z complex128 = math.Pi
    // 3.1415927 3.141592653589793 (3.141592653589793+0i)
    fmt.Println(x, y, z)

    var f float64 = 212
    /*
        100
        0
        100
     */
    // 注意计算次序导致的不同结果
    fmt.Println((f - 32) * 5 / 9)
    fmt.Println( 5 / 9 * (f - 32))
    fmt.Println(5.0 / 9.0 * (f - 32))

    /*
    number:
    untyped integer -> int
    untyped floating-point -> float64
    untyped complex -> complex128
    string:
    untyped rune -> int32 (rune)
    untyped string
    boolean:
    untyped boolean
    */
    fmt.Printf("%T\n", 0)
    fmt.Printf("%T\n", 0.0)
    fmt.Printf("%T\n", 0i)
    fmt.Printf("%T\n", '\000')

    // 类型转换的方式
    var i = int8(0)
    var j int8 = 0
    fmt.Println(i, j)

    const (
        deadbeef = 0xdeadbeef
        a = uint32(deadbeef)
        b = float32(deadbeef)
        c = float64(deadbeef)
        // ./untypedconstants.go:52:18: constant 3735928559 overflows int32
        // ./untypedconstants.go:53:20: constant 1e+309 overflows float64
        // ./untypedconstants.go:54:17: constant -1 overflows uint
        // d = int32(deadbeef)
        // e = float64(1e309)
        // g = uint(-1)
    )
    // int
    // 3735928559 3.7359286e+09 3.735928559e+09
    fmt.Printf("%T\n", deadbeef)
    fmt.Println(a, b, c)
}

