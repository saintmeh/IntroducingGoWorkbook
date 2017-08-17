#Chapter 6 Exercises

1) sum is a function that takes a slice of number and adds them together.  What would its function signature look like in Go?
```go
func sum(numbers... float64) (total float64)
```
2) Write a function that takes an integer and halves it and returns true if it was even or false if it was odd.  For example, half(1) should return (0, false) and half(2) should return (1, true).
```go
//Exercise 2: make a half function with two relevant ret vals(the halved int as int and whether the input was even).
func half(inputInt int) (outputInt int, even bool) {
    even = inputInt%2 ==0
    outputInt = int(inputInt/2)
    return
}
```
3) Write a function with one variadic parameter that finds the greatest number in a list of numbers.
```go
//EXAMPLE CALL:
//    arr2 := []float64{2,5,1,-6.6}
//    slice2 := arr2[:]
//    fmt.Println(max(slice2...) )

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
```
4) Using "makeEvenGenerator" as an example, write a makeOddGenerator function which generates odd numbers.
```go
func makeOddGenerator() func() uint {
    i := uint(1)
    return func() (ret uint) {
        ret = i 
        i += 2
        return
    }
}
```
5) The Fibonacci sequence is defined as: fib(0) = 0, fib(1) = 1, fib(n) = fib(n-1) + fib(n-2).  Write a recursive function that can find fib(n).
```go
func fib(n int)(retVal int) {
    if(n<=1) {
        return n
    } 
    retVal = fib(n-1) + fib(n-2)
    return
}
```
6) What are defer, panic, and recover?  How do you recover from a runtime panic?

**defer:  runs the given function BEFORE the enclosed method returns/panics**

**panic:  Throw an error with a specific message for troubleshooting**

**recover:  Do not throw an ABEND error, but instead make note that your code recovered from what would have been a fatal error**


7) How do you get the memory address of a variable?
```go
&myPtr
```
8) How do you assign a value to a pointer?
```go
*myPtr = 345
```
9) How do you create a new pointer?
```go
var myPtr *int

//or

myPtr := new(int)

```
10) What is the value of x after running this program:
```go
func square(x *float64) {
    *x = *x * *x
}
func main() {
    x := 1.5
    square(&x)
}
```
##Answer
**x = 2.25**

11) Write a program using pointers that can swap two integers (x := 1; y := 2; swap(&x, &y) should give you x=2 and y=1)
```go
//Exercise 11:  swap with ptrs
func swap(x *int, y *int) {
    temp := *x
    *x = *y
    *y = temp
}
```
