#Chapter 7 Exercises

1)  What's the difference between a method and a function?

**A method is an action you would expect to perform on or with its parent object.  Functions are not tied to an object.**

2)  Why would you use an embedded anonymous field instead of a normal named field?

**Embedding an anonymous field type instead of a normal named field creates an is-a relationship without relying on an interface.  A struct can have multiple such is-a relationships.**

3)  Add a new perimeter method to the shape interface to calculate the perimeter of a shape.  Implement the method for circle and rectangle.

**Full code is in [chapt7Structs.go](./chapt7Structs.go)**
```go

type Shape interface {
    area() float64
    perimeter() float64
}

func (c *Circle) perimeter() float64 {
    return c.r * math.Pi *2
}

func (r *Rectangle) perimeter() float64 {
    return math.Abs( (r.x1 - r.x2)*2)+ math.Abs( (r.y1 - r.y2)*2 )
}

```
