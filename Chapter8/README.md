##Chapter 8 Exercises 

1)  Why do we use packages?

**Three reasons:  To maintain scope between identical variable/method names.  And to organize code to make it easier to reuse.  It also speeds up compile time by breaking code up into relevant chunks."**

2)  What is the diference between an identifier that starts with a capital letter and one that doesn't (e.g. Average vs average)?
"Average" can be seen by other packages.  "average" can only be seen by the package it was declared in.


3)  What is a package alias?  How do you make one?

```go
import m "myawesome/directory/structure/math"
.
.
.
x := m.SquareRoot(9)
```

**An alias is just a way to specifically reference a package.  Let's say that you had two packages by the same name; you could use an alias to seperate them**

4)  Create Min and Max functions that find the minimum and maximum values in a slice of float64s.
```go
//ALSO FOUND IN MATH.GO:
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

```


5)  How would you document the functions you created in #4?

**run godoc command to spin up web server pointed to /my/directory/package**

**Another way(for small scripts) is to just cat to std(or a webfile) out using -html flag and then pointing to the package**
```
godoc /home/meh2-pc/testbed/go/IntroducingGoWorkbook/Chapter8/math/ > /var/www/developmentServerReverseProxyA/html/docs/scriptA/math.htm
```

