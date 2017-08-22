package math

import "testing"


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
    average float64
}


var tests = []testpair{
    { []float64{1,2}, 1.5},
    { []float64{1,1,1,1,1}, 1},
    { []float64{-1,1}, 0},
    { []float64{0,0,0}, 0},
}

func TestAverage(t *testing.T) {
    for _, testPair := range tests {
        v := Average(testPair.values)
        if (v != testPair.average){
            t.Error(
                "For", testPair.values,
                "Got", v,
                "Expected", testPair.average,
            )
        }
    }
}
