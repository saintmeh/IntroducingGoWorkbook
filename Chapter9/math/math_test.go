package math

import (
    "testing"
    "math"
)


/*func TestAverage(t *testing.T) {
    v := Average([]float64{1,2})
    if v != 1.5 {
        t.Error("Expected 1.5, got: ", v)
    }
}*/
//After above function.  I was able to test with "go test"
//meh2-pc@meh2-pc ~/testbed/go/IntroducingGoWorkbook/Chapter9/math $ go test
//PASS
//ok  	_/home/meh2-pc/testbed/go/IntroducingGoWorkbook/Chapter9/math	0.002s


//Better way to set up tests:
type testpair struct {
    values[] float64
    result float64
}


var tests = []testpair{
    { []float64{1,2}, 1.5},
    { []float64{1,1,1,1,1}, 1},
    { []float64{-1,1}, 0},
    { []float64{0,0,0}, 0},
}


var minTests = []testpair{
    {[]float64{2,4,6,8,10}, 2},
    {[]float64{2,4,-6,8,10}, -6},
    {[]float64{1,1,1,1,1}, 1},
    {[]float64{0}, 0},
    {[]float64{0, (-1) * math.MaxFloat64}, (-1) * math.MaxFloat64},
    {[]float64{math.MaxFloat64}, math.MaxFloat64},
    {[]float64{2,math.MaxFloat64,6,(-1) * math.MaxFloat64,10}, (-1) * math.MaxFloat64},
}

var maxTests = []testpair{
    {[]float64{2,4,6,8,10}, 10},
    {[]float64{-2,-4,-6,-8,-10}, -2},
    {[]float64{1,1,1,1,1}, 1},
    {[]float64{0}, 0},
    {[]float64{0, (-1) * math.MaxFloat64}, 0},
    {[]float64{math.MaxFloat64}, math.MaxFloat64},
    {[]float64{2,math.MaxFloat64,6,(-1) * math.MaxFloat64,10},  math.MaxFloat64},
}


func TestMin(t *testing.T) {
    for _, test := range minTests {
        v := Min(test.values...)
        if (v != test.result){
            t.Error(
                "For", test.values,
                "Got", v,
                "Expected", test.result,
            )
        }
    }
}

func TestMax(t *testing.T) {
    for _, test := range maxTests {
        v := Max(test.values...)
        if (v != test.result){
            t.Error(
                "For", test.values,
                "Got", v,
                "Expected", test.result,
            )
        }
    }
}

func TestAverage(t *testing.T) {
    for _, test := range tests {
        v := Average(test.values)
        if (v != test.result){
            t.Error(
                "For", test.values,
                "Got", v,
                "Expected", test.result,
            )
        }
    }
}
