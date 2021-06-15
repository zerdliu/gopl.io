package main

import "fmt"

func main() {
    var a [3]int
    fmt.Println(a[0])
    fmt.Println(a[len(a)-1])

    for i, v := range a {
        fmt.Printf("%d %d\n", i, v)
    }

    var q [3]int = [3]int{1, 2, 3}
    var r [3]int = [3]int{1, 2}
    // 3 0
    fmt.Println(q[2], r[2])

    q = [...]int{1,2,3}
    // [3]int
    fmt.Printf("%T\n", q)
    // cannot use [4]int literal (type [4]int) as type [3]int in assignment
    // q = [4]int{1, 2, 3, 4}

    z := [...]int{29: -1}
    // [0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 -1]
    fmt.Println(z)

    i := [2]int{1, 2}
    j := [...]int{1, 2}
    k := [2]int{1, 3}
    // true false false
    fmt.Println( i == j, i == k , j == k )
    // [2]int [2]int [2]int
    fmt.Printf("%T %T %T\n", i, j, k)
    // l := [3]int{1, 2}
    // invalid operation: i == l (mismatched types [2]int and [3]int)
    // fmt.Println(i == l)

    var x = [32]byte{10: 1}
    // [0 0 0 0 0 0 0 0 0 0 1 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0]
    // [0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0]
    fmt.Println(x)
    zero(&x)
    fmt.Println(x)

    // [0 1 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0]
    // [0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0]
    var y = [32]byte{1: 1}
    fmt.Println(y)
    zero2(&y)
    fmt.Println(y)
}

// go用值拷贝，所以如果不用array 指针的话，数组会被copy并传入函数中
func zero(ptr *[32]byte) {
    for i := range ptr {
        ptr[i] = 0
    }
}

func zero2(ptr *[32]byte) {
    // 利用字面量来清零
    *ptr = [32]byte{}
}
