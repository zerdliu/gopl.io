package main

import (
    "fmt"
    "time"
)

func main() {

    // 常量是基本类型：bool string int
    // 或者基本类型的别名，比如: time.Minute
    const (
        b = true
        i = 1
        s = "liuzhuo"
        t = time.Minute
        c = 'c'
    )
    const IPv4Len = 4

    // 可用于数组长度
    var p [IPv4Len]byte
    fmt.Println(p)

    // 可以指定类型
    const noDelay time.Duration = 0
    const timeout = 5 * time.Minute

    fmt.Printf("%T %[1]v\n", noDelay)
    fmt.Printf("%T %[1]v\n", timeout)
    fmt.Printf("%T %[1]v\n", time.Minute)

    const (
        d = 1
        e
        f
        g = 10
        h
    )

    // 1 1 1 10 10
    fmt.Println(d, e, f, g, h)
}
