package main

import (
    "os"
    "fmt"
    "strconv"
)

func printUsage() {
    fmt.Println("Usage:\n./test num1 num2")
}

func main() {
    if (len(os.Args) == 3) {
        num1, err := strconv.Atoi(os.Args[1])
        if (err != nil) {
            printUsage();
            os.Exit(0);
        }
        num2, err := strconv.Atoi(os.Args[2])
        if (err != nil) {
            printUsage();
            os.Exit(0);
        }
        fmt.Println(num1, " + ", num2, " = ", (num1 + num2))
    } else {
        printUsage();
        os.Exit(0);
    }
}
