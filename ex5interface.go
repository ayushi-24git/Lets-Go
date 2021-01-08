package main

import (
    "fmt"
    "os"
    "strconv"
)

type rect struct {
    l float64
    w float64
}

type circle struct {
    radius float64
}

func (objr rect) area() float64 {
    return objr.l * objr.w
}

func (objr rect) per() float64 {
    return 2*(objr.l + objr.w)
}

func (objc circle) area() float64 {
    return 3.14*objc.radius*objc.radius
}

func (objc circle) per() float64 {
    return 2*3.14*objc.radius
}


type geometry interface {
    area() float64
    per() float64
}

func ans(g geometry) {
    area := g.area()
    perimeter := g.per()
    fmt.Println("Area:", area)
    fmt.Println("Perimeter:", perimeter)
}

func main() {
    shape := os.Args[1]
    if shape == "rectangle" {
        length := os.Args[2]
        width := os.Args[3]
       
        len,_ := strconv.ParseFloat(length,64)
        wid,_ := strconv.ParseFloat(width,64)
 
        rectangle := rect{len,wid}

        ans(rectangle)
    } else {
          r := os.Args[2]
          radius,_ := strconv.ParseFloat(r,64)

          Circle := circle{radius}

          ans(Circle)
    }
}
