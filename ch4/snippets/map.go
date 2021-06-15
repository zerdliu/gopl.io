package main

import (
    "fmt"
    "sort"
)

func main() {
    var ages map[string]int
    fmt.Println(ages == nil)  // true
    fmt.Println(len(ages)==0) // true

    // 不能不初始化就使用
    //  assignment to entry in nil map
    // ages["carol"] = 21

    // 初始化后，就可以设置元素了
    // 初始化方法1
    ages = make(map[string]int)
    ages["carol"] = 21
    // map[carol:21]
    fmt.Println(ages)

    // 初始化方法2
    // map[]
    // map[carol:21]
    ages = map[string]int{}
    fmt.Println(ages)
    ages["carol"] = 21
    fmt.Println(ages)

    // 初始化方法3
    // map[alice:31 charlie:34]
    // map[alice:31 carol:21 charlie:34]
    ages = map[string]int{
        "alice": 31,
        "charlie": 34,
    }
    fmt.Println(ages)
    ages["carol"] = 21
    fmt.Println(ages)

    // map[carol:21 charlie:34]
    delete(ages, "alic")
    fmt.Println(ages)

    // map[alice:31 bob:1 carol:21 charlie:34]
    // map[alice:31 bob:2 carol:21 charlie:34]
    ages["bob"] += 1
    fmt.Println(ages)
    ages["bob"] ++
    fmt.Println(ages)

    // 不能破坏map的封装：
    //   1. hash算法的可变
    //   2. 排序的可变
    //   3. 可重新分配地址
    // cannot take the address of ages["bob"]
    //_ = &ages["bob"]

    // alice   31
    // charlie 34
    // carol   21
    // bob     2
    for name, age := range ages {
        fmt.Printf("%s\t%d\n", name, age)
    }

    var names []string
    // 0
    fmt.Println(cap(names))
    names = make([]string, 0, len(ages))
    // 4
    fmt.Println(cap(names))
    // 等价于 for name, _ := range ages {
    for name := range ages {
        names = append(names, name)
    }
    // alice   31
    // bob     2
    // carol   21
    // charlie 34
    sort.Strings(names)
    for _, name := range names {
        fmt.Printf("%s\t%d\n", name, ages[name])
    }

    // 0
    if age, ok := ages["ruby"]; !ok {
        fmt.Println(age)
    }

    // 没有输出
    ages["ruby"] = 0
    if age, ok := ages["ruby"]; !ok {
        fmt.Println(age)
    }

}
