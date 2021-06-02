package main

import "fmt"

func main() {

    var s string
    // 短路，所以s[0]写法是安全的
    fmt.Println(s != "" && s[0] == 'x')

    fmt.Println(btoi(true), itob(12))
}

func btoi(b bool) int {
    if b {
        return 1
    }
    return 0
}

func itob(i int) bool {
    return i != 0
}
