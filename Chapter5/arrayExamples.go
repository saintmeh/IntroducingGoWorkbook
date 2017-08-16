package main

import "fmt"

func main() {
    var x [6]float64
    x[0] = 30

    fmt.Println("===x===")
    for i:=0; i<len(x); i++ {
        fmt.Println(x[i])
    }




    y  := [6]float64{30,
        3,
        5,
        7}

    fmt.Println("===y===")
    for _, value := range y {
        fmt.Println(value)
    }




    var z [4]float64
    z = [4]float64{30,
        3,
        5,
        7}

    fmt.Println("===z===")
    for _, value := range z {
        fmt.Println(value)
    }





}
