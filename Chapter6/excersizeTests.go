package main

import "fmt"
import "math"


func main() {
    //Exercise 1: make a sum func with a slice param.
    fmt.Println("sum:")
    arr := []float64{5,8,1,4,4,6}
    slice := arr[3:5]
    fmt.Println( sum(slice) ) 

    //Exercise 2: make a half function with two relevant ret vals.
    fmt.Println("half:")
    num1, even1 := half(4)
    num2, even2 := half(5)
    num3, even3 := half(6)
    fmt.Println( num1, even1 )
    fmt.Println( num2, even2 )
    fmt.Println( num3, even3 )
    
    //Exercise 3 Use variadic param to find max number in slice/array
    fmt.Println("variadic:")
    arr2 := []float64{2,5,1,-6.6}
    slice2 := arr2[:]
    fmt.Println(max(slice2...) )

    //Exercise 4:  simple modification to existing nested closure function
    fmt.Println("makeOddGenerator seq:")
    nextOdd := makeOddGenerator()
    fmt.Println(nextOdd())
    fmt.Println(nextOdd())

    //Exercise 5
    fmt.Println("Fibonacci seq:")
    fmt.Println(fib(0))
    fmt.Println(fib(1))
    fmt.Println(fib(2))
    fmt.Println(fib(3))
    fmt.Println(fib(4))
    fmt.Println(fib(5))
    fmt.Println(fib(6))
    fmt.Println(fib(7))
    fmt.Println(fib(8))
    fmt.Println(fib(9))

    //Exercise 10:  Pointers
    fmt.Println("Pointer value of x:")
    x := 1.5
    square(&x)
    fmt.Println(x)

    //Exercise 11:  swap with ptrs
    x2 := 1
    y2 := 2
    fmt.Println("old values x:",x2,"and y:",y2)
    swap(&x2,&y2)
    fmt.Println("new values x:",x2,"and y:",y2)
}


    //Exercise 1: make a sum func with a slice param.
func sum(numbers []float64 ) (retVal float64) {
    for _, value := range numbers {
        retVal += value
    }
    return
}


//Exercise 2: make a half function with two relevant ret vals.
func half(inputInt int) (outputInt int, even bool) {
    even = inputInt%2 ==0
    outputInt = int(inputInt/2)
    return
}

//Exercise 3:  Use variadic param to find max number in slice/array
func max(numbers... float64) (maxValue float64) {
    //initialize maxValue as the lowest possible float64 value(two's compliment has been considered)
    maxValue = (-1) * math.MaxFloat64
    fmt.Println(maxValue)
    //
    for _, value := range numbers {
        if(value>maxValue) {
            maxValue = value
        }
    }
    return
}

//Exercise 4:
func makeOddGenerator() func() uint {
    i := uint(1)
    return func() (ret uint) {
        ret = i 
        i += 2
        return
    }
}

//Exercise 5:Write a recursive function that can find fib(n).
func fib(n int)(retVal int) {
    if(n<=1) {
        return n
    } 
    retVal = fib(n-1) + fib(n-2)
    return
}


//Exercise 10:  Pointers
func square(x *float64) {
    *x = *x * *x
}

//Exercise 11:  swap with ptrs
func swap(x *int, y *int) {
    temp := *x
    *x = *y
    *y = temp
}
