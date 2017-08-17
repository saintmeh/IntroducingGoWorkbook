package main

import "fmt"

func main() {

    fmt.Println(exampleA(4))

    fmt.Println(exampleB(42))
    context, value := exampleB(42)
    fmt.Println(context)
    fmt.Println(value)

    fmt.Println(exampleC(1,2,3,4,5))//should be 15

    //Closure... it's a thing in Go:
    exampleD := func() string{
        return "Hello waldo!"
    }
    fmt.Println(exampleD())

    fmt.Println(exampleE(3))

    fmt.Println(exampleF(2))    

    fmt.Println(exampleG(33))

    var x int = 7
    fmt.Println("x before exampleH(&x):",x)
    exampleH(&x)
    fmt.Println("x after exampleH(&x):",x)


    //another more familiar way to dereference pointers:
    y := new(int)
    *y = 6
    fmt.Println("y before exampleH(y):",*y)
    exampleH(y)
    fmt.Println("y after exampleH(y):",*y)
}

//Woah!  return values can have default names!  Cool!
func exampleA(a int) (retVal int) {
    retVal = a+1
    return
}

//Also cool, everything can return multiple values:
func exampleB(a int) (string, int) {
    return "You ran example b! You passed in parameter", a
}

//Variadic function...  these are familiar...
func exampleC(a ...int) (sum int) {
    for _, value := range a {
        sum += value
    }
    return
}

//Good old recursion...  (factorial)
func exampleE(x uint) uint {
    if(x == 0) {
      return 1 
    }
    return x * exampleE(x-1)
}

//Should inlined deferred functions be used as post conditions?
//Pros:  Post conditions are *ALWAYS* ran even when someone in the future splits the return with an if statement
//       Post conditions are easy to find.  
//Cons:  Post conditions must be written with panics in mind.
func exampleF(x int) (retVal int) {
    defer func () {
        retVal=x+3
    }()
    retVal=999//This will be overwritten by the deferred function
    return
}

//Panic and recover using a deferred function.
func exampleG(x int) int {
    defer func() {
        str := recover()
        fmt.Println(str, x)
    }()
    panic("PANIC TEST")
} 

//use of byRef or pointers to modify passed-in parameters.
func exampleH(xPtr *int) {
    *xPtr = 0
}

