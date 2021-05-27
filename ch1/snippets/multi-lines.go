package main

import "fmt"

func main() {
    fmt.Println(add(1,2))
}

func add(x int, y int) int {
    // + 后可以省略换行符，但是不能在下一行的开头
    return x +
        y
}