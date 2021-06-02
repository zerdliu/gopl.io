package main

import "fmt"

func main() {

    s := "hello, world"
    fmt.Println(len(s))
    fmt.Println(s[0], s[7])
    fmt.Printf("%c, %[1]x\n", s[0])

    //panic: runtime error: index out of range [12] with length 12
    //fmt.Println(s[len(s)])

    fmt.Println(s[:5])
    fmt.Println(s[7:])
    fmt.Println(s[:])

    fmt.Println("goodbye" + s[5:])

    // 字符串是不可变的
    s = "left foot"
    t := s
    s += ", right foot"

    fmt.Println(s)
    fmt.Println(t)

    //cannot assign to s[0]
    //s[0] = 'L'
}
