package main

import (
    "flag"
    "fmt"
    "os"
    "strconv"
)

func main() {
    flag.Parse()
    i := flag.Arg(0)
    j := flag.Arg(1)
    a, err := strconv.Atoi(i)
    b, err := strconv.Atoi(j)
    if err != nil {
        fmt.Println(err)
        os.Exit(2)
    }
    fmt.Println(a+b)
}
