##Chapter 9 Exercises

1)  Writing a good suite of tests is not always easy, but the process of writing tests often reveals more about a problem than you may first realize.  For example, with our **Average** function, what happens if you pass in an empty list([]float64{})?  How could the function be modified to return 0 in this case?

**I did something like this [four chapters ago](../Chapter5/sliceExamples.go) with the "min" function.  This precondition calls a panic instead of returning a valid result for invalid input.  It's better to fail as soon as an invalid value is discovered.  Otherwise we end up treating symptoms of symptoms of symptoms... ad naseum.  Having said that, I could have chose to "defer recover()" from the error and return a 0 float64 value.  This is acceptable, while still logging that there was cause for a potential future error.  I would normally suggest this as a compromise solution, but I wouldn't push.  Here is the code the book asked for:**
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
```go
import (
    "testing"
    "math"
)
.
.
.
//create the test pair struct
type testPair struct {
    values []float64
    result float64
}

minTests := []testPair{
    {[]float64{2,4,6,8,10}, 2},
    {[]float64{2,4,-6,8,10}, -6},
    {[]float64{1,1,1,1,1}, 1},
    {[]float64{0}, 0},
    {[]float64{0, (-1) * math.MaxFloat64}, (-1) * math.MaxFloat64},
    {[]float64{math.MaxFloat64}, math.MaxFloat64},
    {[]float64{2,math.MaxFloat64,6,(-1) * math.MaxFloat64,10}, (-1) * math.MaxFloat64},
}

maxTests := []testPair{
    {[]float64{2,4,6,8,10}, 10},
    {[]float64{-2,-4,-6,-8,-10}, -2},
    {[]float64{1,1,1,1,1}, 1},
    {[]float64{0}, 0},
    {[]float64{0, (-1) * math.MaxFloat64}, 0},
    {[]float64{math.MaxFloat64}, math.MaxFloat64},
    {[]float64{2,math.MaxFloat64,6,(-1) * math.MaxFloat64,10},  math.MaxFloat64},
}


TestMin(t *testing.T) {
    for _, testPair := range tests {
        v := Min(testPair.values)
        if (v != testPair.result){
            t.Error(
                "For", testPair.values,
                "Got", v,
                "Expected", testPair.result,
            )
        }
    }
}

TestMax(t *testing.T) {
    for _, testPair := range tests {
        v := Min(testPair.values)
        if (v != testPair.result){
            t.Error(
                "For", testPair.values,
                "Got", v,
                "Expected", testPair.result,
            )
        }
    }
}
    
```

**meh2-pc@meh2-pc ~/testbed/go/IntroducingGoWorkbook/Chapter9/math $ go test**

**PASS**

**ok  /home/meh2-pc/testbed/go/IntroducingGoWorkbook/Chapter9/math	0.002s**
