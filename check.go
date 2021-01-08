package main

import "fmt"
import "os"
import "strconv"

func main() {
    i := os.Args[1]
    n,err := strconv.Atoi(i)
    if err != nil {
    	fmt.Println("Not a number!")
        os.Exit(1)
    }
    if n % 2 == 0 {
        fmt.Println("It's an even number!")
    } else {
        fmt.Println("It's an odd number!")
    }
}
