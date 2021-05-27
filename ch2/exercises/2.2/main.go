// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 43.
//!+

// Cf converts its numeric argument to Celsius and Fahrenheit.
package main

import (
    "fmt"
    "os"
    "strconv"

    "github.com/zerdliu/gopl.io/ch2/exercises/2.2/lengthconv"
    "github.com/zerdliu/gopl.io/ch2/exercises/2.2/weightconv"
)

func main() {
    for _, arg := range os.Args[1:] {
        n, err := strconv.ParseFloat(arg, 64)
        if err != nil {
            fmt.Fprintf(os.Stderr, "cf: %v\n", err)
            os.Exit(1)
        }
        f := lengthconv.Foot(n)
        m := lengthconv.Meter(n)

        p := weightconv.Pound(n)
        k := weightconv.Kilogram(n)

        fmt.Printf("%s = %s, %s = %s\n", f, lengthconv.FToM(f), m, lengthconv.MToF(m))
        fmt.Printf("%s = %s, %s = %s\n", p, p.ToKilogram(), k, k.ToPound())
    }
}

//!-
