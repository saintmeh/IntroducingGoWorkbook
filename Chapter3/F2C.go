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
