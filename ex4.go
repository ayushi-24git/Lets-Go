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

func main() {
    shape := os.Args[1]
    if shape == "rectangle" {
        length := os.Args[2]
        width := os.Args[3]
       
        len,_ := strconv.ParseFloat(length,64)
        wid,_ := strconv.ParseFloat(width,64)
 
        rectangle := rect{len,wid}

        Area := rectangle.area()
        fmt.Println("Area:", Area)

        Perimeter := rectangle.per()
        fmt.Println("Perimeter:",Perimeter)
    } else {
          r := os.Args[2]
          radius,_ := strconv.ParseFloat(r,64)

          Circle := circle{radius}

          Area := Circle.area()
          fmt.Println("Area:", Area)

          Perimeter := Circle.per()
          fmt.Println("Perimeter:",Perimeter)
    }
}
