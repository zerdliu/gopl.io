package main

import (
    "bufio"
    "fmt"
    "os"
)

type WhichFile map[string]bool
type Counter struct {
    WhichFile
    count int
}

func main() {
    // 构建一个map
    counts := make(map[string]Counter)
    // 没提供参数（文件名）
    files := os.Args[1:]
    if len(files) == 0 {
        countLines(os.Stdin, counts)
    // 提供了参数（文件名）
    } else {
        for _, arg := range files {
            f, err := os.Open(arg)
            if err != nil {
                fmt.Fprintf(os.Stderr, "dup2: %v\n", err)
                continue
            }
            countLines(f, counts)
            f.Close()
        }
    }

    // 打印结果
    //fmt.Println(counts)
    for line, counter := range counts {
        if counter.count > 1 {
            filenameString := ""
            sep := ""
            for filename, _ := range counter.WhichFile {
                filenameString += sep + filename
                sep = ","
            }
            fmt.Printf("%d\t%s\t%s\n", counter.count, line, filenameString)
        }
    }
}

// 公共函数
func countLines(f *os.File, counts map[string]Counter) {
    input := bufio.NewScanner(f)
    for input.Scan() {
        counter, ok := counts[input.Text()]
        if ok {
            counter.WhichFile[f.Name()] = true
        } else {
            file := WhichFile{
                f.Name() : true,
            }
            counter.WhichFile = file
        }
        counter.count += 1
        counts[input.Text()] = counter
    }
}