package main

import "fmt"

func main() {
    fmt.Println(newInt() == newIntDummy())        // false
    fmt.Println(newInt() == newInt())             // false
    fmt.Println(struct{}{} == struct{}{})         // true
    fmt.Println(&struct{}{} == &struct{}{})  // false
    fmt.Println([0]int{} == [0]int{})            // true
    fmt.Println(&[0]int{} == &[0]int{})     // false
    fmt.Printf("%T, %T\n", struct{}{}, &struct{}{}) // struct {}, *struct {}
    fmt.Printf("%T, %T\n", [0]int{}, &[0]int{})     // [0]int, *[0]int
    fmt.Printf("%T, %T\n", []int{}, &[]int{})     // []int, *[]int

    fmt.Printf("%T\n", newInt())
}

func newInt() *int {
    return new(int)
}

func newIntDummy() *int {
    var dummy int
    return &dummy
}
