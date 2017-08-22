package main

import (
    "fmt"
    "math"
)

func main() {

    var c1 Circle//this is unusual.  Normally we use pointers for new instances of structs so methods can be easily attached.  And we use the "new" function to initialize everything to zero as well as alloc memory for more complicated properties.
    c1.r=234
    fmt.Println(c1.r)

    var c2 = new(Circle)//It's better to initialize everything like this
    fmt.Println(*c2)

    c3 := &Circle{x:3, y:62, r:42}  //Best practice is to give each field an explicit initial value(if possible)
    fmt.Println(*c3)

    c4 := &Circle{4,62,42}//it's possible to leave off the identifiers... but why would we?
    fmt.Println(*c4)

    fmt.Println(c4.area())

    andy := new(Android)
    andy.name="Androwd"
    andy.uptime=65446655
    andy.Talk()
    fmt.Println(andy.uptime)

    r1 := &Rectangle{0,0,50,50}
    fmt.Println(r1.area())

    fmt.Println(totalAreas(r1,c2,c3,c4))

    m := new(MultiShape)
    m.shapes = []Shape{&Rectangle{0,0,1,1}, &Circle{50,50,5} }


    fmt.Println(m.area())

    fmt.Println(m.shapes[0].perimeter())
    fmt.Println(m.shapes[1].perimeter())
}



type Circle struct {
//    x float64
//    y float64
//    r float64
//    A Better way to declare multiple variables of same type(similar to java 5.0+):
    x, y, r float64
}

func (c *Circle) area() float64 {
    return math.Pi * c.r * c.r
}

func (c *Circle) perimeter() float64 {
    return c.r * math.Pi *2
}


type Rectangle struct {
    x1, y1, x2, y2 float64
}

func (r *Rectangle) area() float64 {
    return math.Abs( (r.x1 - r.x2) * (r.y1 - r.y2) )
}


func (r *Rectangle) perimeter() float64 {
    return math.Abs( (r.x1 - r.x2)*2)+ math.Abs( (r.y1 - r.y2)*2 )
}

// Interfaces

type Shape interface {
    area() float64
    perimeter() float64
}

func totalAreas(shapes... Shape) (total float64) {
    for _, shape := range shapes {
        total += shape.area()
    }
    return
}


type MultiShape struct {
    shapes []Shape
}

func (ms *MultiShape) area() (total float64) {
    for _, shape := range ms.shapes {
        total += shape.area()
    }
    return
}




//I used Person and Android to understand is-a relationships with embedded types

type Person struct {
    name string 
}

type Machine struct {
    uptime float64
}

type Android struct {
    Person
    Machine
    model string
}

func (p *Person) Talk() {
    fmt.Println("Hello, my name is", p.name)
}



