##Chapter 3 Exercises

1) what are two ways to create a new variable?
```go
var x string = "blarg"
y := "blargle"
//I'm told to use the shorter version whenever possible.  (last sentence of page 40 of PDF page 19 in the corner)
```
2) What is the value of x after running
```go
x := 5
x += 1
```
**the value is six**

3)what is scope?  How do you determine the scope of a variable in Go?
```go
Scope is the area of the source code where a programmer may write code to access or modify a variable.
```

4)What is the difference between var and const?
```
const indicates the value can not be changed later
var indicates that the coder may continue to change the variable's value as long as he/she is in scope
```

5)Write a program that converts from Fahrenheit into Celsius (C=(F-32) * 5/9).

```go
package main

import (
	"fmt"
)


func main() {
  fmt.Println("===Conversion utility for Fahrenheit to Celsius===")
  fmt.Print("Enter a temperature in Fahrenheit:  ")
  var tempF float64
  fmt.Scanf("%f", &tempF)
  
  //Convert to Celsius 
  tempC := (tempF - 32) * 5/9
  
  fmt.Println(tempC)
}
```

6)Write another program which converts from feet into meters (1 ft = 0.3048 m)

```go
package main

import (
	"fmt"
)


func main() {
  fmt.Println("===Conversion utility for Feet to Meters===")
  fmt.Print("Enter a distance in Feet:  ")
  var distF float64
  fmt.Scanf("%f", &distF)
  
  //Convert to Meters
  const feetInAMeter float64 = 3.28084
  distM := distF / feetInAMeter
  
  //Handle the plurality of foot
  pluralSensitiveFoot := "feet"
  if distF == 1 {
    pluralSensitiveFoot = "foot"
  }

  fmt.Println(distF,pluralSensitiveFoot," = ",distM, "meter(s)")
}
