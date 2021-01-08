package main

import "fmt"

func sum(Sum chan int) {
    s := 0
    for i := 1; i <50 ; i++ {
        if i %3 == 0 {
	    s = s + i
	} else {
	      if i % 5 ==0 {
	          s = s + i
              }
	  }
    }
    Sum <- s
}

func main() {
    makechan := make(chan int)
    go sum(makechan)
    fmt.Println("Required sum = ", <-makechan)
}
