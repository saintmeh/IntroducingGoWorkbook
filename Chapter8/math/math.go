package math

//averages a slice of floats
func Average(xs []float64) (r float64) {
    total := float64(0)
    for _, x := range xs {
        total +=x
    }
    r = total / len(xs)
    return 
}

//Exercise 4:   Create Min and Max functions that find the minimum and maximum values in a slice of float64s.
func Max(numbers... float64) (maxValue float64) {
    //initialize maxValue as the lowest possible float64 value(two's compliment has been considered)
    maxValue = (-1) * math.MaxFloat64
    fmt.Println(maxValue)
    //TODO: OPTIMIZATION Opportunity:  for longer arrays, implement a list and sorting. 
    for _, value := range numbers {
        if(value>maxValue) {
            maxValue = value
        }
    }
    return
}

//Exercise 4:   Create Min and Max functions that find the minimum and maximum values in a slice of float64s.
func Min(numbers... float64) (minValue float64) {
    //initialize minValue as the highest possible float64 value(two's compliment has been considered)
    minValue = math.MaxFloat64
    fmt.Println(minValue)
    //TODO: OPTIMIZATION Opportunity:  for longer arrays, implement a list and sorting
    for _, value := range numbers {
        if(value<minValue) {
            minValue = value
        }
    }
    return
}
