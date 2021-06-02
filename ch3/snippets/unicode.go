package main

import (
    "fmt"
    "unicode/utf8"
)

func main() {

    // rune == int32

    c1 := '世'
    c2 := '\u4e16'
    c3 := '\U00004e16'
    //invalid character literal (more than one character)
    //c4 := '世界'

    // 世, 世, 世
    fmt.Printf("%c, %[1]c, %[1]c\n", c1, c2, c3)
    // 19990 19990 19990
    fmt.Println(c1, c2, c3)

    s1 := "世界"
    //世界
    fmt.Println(s1)

    // 世
    // 通常不要按照下面这样写，采用\u或者\U
    s2 := "\xe4\xb8\x96"
    fmt.Println(s2)
    // invalid character literal (more than one character)
    // c5 := '\xe4\xb8\x96'

    s := "Hello, 世界"
    fmt.Println(len(s)) // 13 = 7 + 2 * 3
    fmt.Println(utf8.RuneCountInString(s)) // 8 = 7 + 2

    /*
    0       H       1  Hello, 世界
    1       e       1  ello, 世界
    2       l       1  llo, 世界
    3       l       1  lo, 世界
    4       o       1  o, 世界
    5       ,       1  , 世界
    6               1   世界
    7       世      3  世界
    10      界      3  界
    */
    for i := 0; i < len(s); {
        // 解析一个字符串的第一个rune
        r, size := utf8.DecodeRuneInString(s[i:])
        fmt.Printf("%d\t%c\t%d\t%s\n", i, r, size, s[i:])
        i += size
    }

    /*
    0       H       72
    1       e       101
    2       l       108
    3       l       108
    4       o       111
    5       ,       44
    6               32
    7       世      19990
    10      界      30028
    */
    for i, r := range s {
        fmt.Printf("%d\t%c\t%d\n", i, r, r)
    }

    n := 0
    for range s {
        n++
    }
    // 9 13 9
    fmt.Println(n, len(s), utf8.RuneCountInString(s))

    s3 := "世界"
    // e4 b8 96 e7 95 8c
    // [4e16 754c]
    // 注意格式字符串
    fmt.Printf("% x\n", s3)
    r3 := []rune(s3)
    fmt.Printf("%x\n", r3)
    // 世界
    fmt.Println(string(r3))

    // A 京 �
    fmt.Println(string(65), string(0x4eac), string(1234567))
}

