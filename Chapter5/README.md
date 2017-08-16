#Chapter 5 Exercises

1) How do you access the forth element of an array or slice?>
```go
arr[3]
```



2) What is the length of a slice created using make([]int, 3, 9)
```go
    //Length is 3, capacity is 9.
```

3) Given the following array, what would x[2:5] give you?
```go
    x:= [6]string{"a","b","c","d","e","f"}
```
**c**

**d**

**and e**

4) Write a program that finds the smallest number in this list:
```go
    x:=[]int{
        48,96,86,68,
        57,82,63,70,
        37,34,83,27,
        19,97,9,17,
    }
```
##Answer
```go
    //Check each element in the list to see if it is smaller than what we already have. 
    //If it's smaller, set the smaller number to the new smallestNumberInList.
    //Start with the biggest number a signed Int could have and work downward. 
    //Error if the list is empty .
    func min(listOfInts []int) int{ 
        //Preconditions
        if(len(listOfInts)<1) {
            panic(fmt.Sprintf("Parameter must be a list with at least one element not %v", listOfInts) )
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
```
