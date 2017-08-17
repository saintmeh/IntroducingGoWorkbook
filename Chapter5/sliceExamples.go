package main

import "fmt"
import "math"

func main() {
    //Make func example
    fmt.Println("==========Slice using make func=========")
    x := make([]float64, 5)
    x[0]=2345
    for _, value := range x {
        fmt.Println(value)
    }


    //High:low example
    low := 2
    high := 4
    fmt.Println("=========Slice using ",low,":",high,"==========")
    arr := [5]float64{0,1,2,3,4}
    y := arr[low:high]
    for _, value := range y {
        fmt.Println(value)
    }
    low = 3
    fmt.Println("=========Slice using ",low,": ==========")
    y = arr[low:]
    for _, value := range y {
        fmt.Println(value)
    }
    high = 3
    fmt.Println("=========Slice using :",high," ==========")
    y = arr[:high]
    for _, value := range y {
        fmt.Println(value)
    }
    fmt.Println("=========Slice using : ==========")
    y = arr[:]
    for _, value := range y {
        fmt.Println(value)
    }

    //Append example
    fmt.Println("=========appending 3,3,3 ==========")
    z := arr[:]
    z = append(z,3,3,3)
    for _, value := range z {
        fmt.Println(value)
    }

    //copy example
    fmt.Println("=========copy example ==========")
    fmt.Println("before copy:\ny\tz")
    y = append(y,9)
    for  i :=  0 ; i < int(math.Max( float64(len(y)), float64(len(z)) )); i++ {
        if(i < len(y)){
            fmt.Print(y[i],"\t")
        }
        if(i < len(z)){
            fmt.Print(z[i])
        }
        fmt.Println()
   
    }
    
    copy(z,y)
    
    fmt.Println("after copy(z,y):\ny\tz")
    for  i :=  0 ; i < int(math.Max( float64(len(y)), float64(len(z)) )); i++ {
        if(i < len(y)){
            fmt.Print(y[i],"\t")
        }
        if(i < len(z)){
            fmt.Print(z[i])
        }
        fmt.Println()
   
    }

 ax:= [6]string{"a","b","c","d","e","f"}
fmt.Println(ax[2:5])


    x1:=[]int{
        48,96,86,68,
        57,82,63,70,
        37,34,83,27,
        19,97,9,17,
    }

    min(x1);
}

    
    //Check each element in the list to see if it is smaller than what we already have. 
    //If it's smaller, set the smaller number to the new smallestNumberInList.
    //Start with the biggest number a signed Int could have and work downward. 
    //Error if the list is empty .
    func min(listOfInts []int) int{ 
        //Preconditions
        if(len(listOfInts)<1) {
            panic(fmt.Sprintf("Parameter must be a list of signed integers with at least one element not %v", listOfInts) )
        }
        //initialization
        const MaxInt = int(^uint(0) >> 1)
        var smallestNumberInList int = MaxInt
        //body
        for _, value := range listOfInts{
            if value<smallestNumberInList {
                smallestNumberInList=value;
            }
        }
        fmt.Println(smallestNumberInList)
        return smallestNumberInList
    }
