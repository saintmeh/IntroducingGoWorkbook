##Chapter 9 Exercises

1)  Writing a good suite of tests is not always easy, but the process of writing tests often reveals more about a problem than you may first realize.  For example, with our **Average** function, what happens if you pass in an empty list([]float64{})?  How could the function be modified to return 0 in this case?

**I did something like this [four chapters ago](../Chapter5/sliceExamples.go) with the "min" function.  This precondition calls a panic instead of returning a valid result for invalid input.  It's better to fail as soon as an invalid value is discovered.  Otherwise we end up treating symptoms of symptoms of symptoms... ad naseum.  Having said that, I could have chose to "defer recover()" from the error and return a 0 float64 value.  This is sometimes acceptable, while still logging that there was an error.  I would do this like so:**
```go
 
//averages a slice of floats
func Average(xs []float64) (r float64) {
    //Preconditions
    if(len(listOfInts)<1) {
        //TODO verify that we don't panic on a zero-length list.  panic(fmt.Sprintf("Parameter must be a list with at least one element not %v", listOfInts) )
        return 0
    }
    total := float64(0)
    for _, x := range xs {
        total +=x
    }
    r = total / float64(len(xs))
    return 
}
    
```



2)  Write a series of tests for the Min and Max functions you wrote in the previous chapter.


