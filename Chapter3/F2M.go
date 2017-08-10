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
