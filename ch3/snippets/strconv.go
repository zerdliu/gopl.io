package main

import (
    "fmt"
    "strconv"
)

func main() {
    x := 123
    y := fmt.Sprintf("%d", x)
    fmt.Println(y, strconv.Itoa(x))

    fmt.Println(strconv.FormatInt(int64(x), 2))

    // 使用Sprintf生成字符串会更方便
    s := fmt.Sprintf("x=%b", x)
    fmt.Println(s)

    a, err := strconv.Atoi("123")
    if err != nil {
        fmt.Println(err)
    }
    b, err := strconv.ParseInt("111", 2, 64)
    if err != nil {
        fmt.Println(err)
    }

    // 123 7
    // int int64
    fmt.Println(a, b)
    fmt.Printf("%T %T\n", a, b)
}
