package main

import "fmt"

func main() {
    s := "世界"
    b := []byte(s)
    r := []rune(s)
    s2 := string(b)

    // 世界 [228 184 150 231 149 140] [19990 30028] 世界
    // string []uint8 []int32 string
    // []byte []rune 的type不同
    fmt.Println(s, b, r, s2)
    fmt.Printf("%T %T %T %T\n", s, b, r, s2)
}
