package main

import (
    "fmt"
)

func main() {
    type Weekday int

    const (
        Sunday Weekday = iota
        Monday
        Tuesday
        Wednesday
        Thursday
        Friday
        Saturday
    )

    // 0 1 2 3 4 5 6
    fmt.Println(Sunday, Monday, Tuesday, Wednesday, Thursday, Friday, Saturday)

    type Flags uint

    const (
        FlagUp Flags = 1 << iota
        FlagBroadcast
        FlagLoopback
        FlagPointToPoint
        FlagMulticast
    )

    // 1 10 100 1000 10000
    // 1 2 4 8 16
    fmt.Printf("%b %b %b %b %b \n", FlagUp, FlagBroadcast, FlagLoopback, FlagPointToPoint, FlagMulticast)
    fmt.Println(FlagUp, FlagBroadcast, FlagLoopback, FlagPointToPoint, FlagMulticast)

    const (
        _ = 1 << (10 * iota)
        KiB
        MiB
        GiB
        TiB
        PiB
        ZiB
        YiB
    )

    // 1024 1048576 1073741824 1099511627776 1125899906842624
    fmt.Println(KiB, MiB, GiB, TiB, PiB)

    // constant 1180591620717411303424 overflows int
    // fmt.Println(ZiB, YiB)

    // 虽然溢出，但无类型常量因为推迟了计算，所以下面的表达式仍然能计算出值
    // 1024
    fmt.Println(YiB/ZiB)

}

