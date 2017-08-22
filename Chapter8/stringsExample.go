package main

import (
    "fmt"
    "strings"
)

func main() {
    //Contains(s, substring string ) bool
    fmt.Println(strings.Contains("test", "st")  )
    
    //Count(s string, search string) int
    fmt.Println(strings.Count("test", "t") ) 

    //HasPrefix(s string, prefix string) bool
    fmt.Println(strings.HasPrefix("test", "tes") )
    fmt.Println(strings.HasPrefix("test", "est") )

    //HasSuffix(s string, prefix string) bool
    fmt.Println(strings.HasSuffix("test", "tes") )
    fmt.Println(strings.HasSuffix("test", "est") )

    //Index(s string, search string) int
    //This returns the first index... like in Javascript/Java/C++/etc
    fmt.Println(strings.Index("test", "e") )
    fmt.Println(strings.Index("test", "t") )

    //Join(joinableStrings []string, delimeter string)
    fmt.Println(strings.Join([]string{"sine","sine","cosine","sine","3.14159!\nGo nerds! Go! Go! Go Nerds!"}, " ") )

    //Repeat(repeatingString string, numberOfTimesToRepeat int)
    fmt.Println( strings.Repeat("testing ",2), "123.  Is this thing on?")

    //Split(stringToSplitIntoArray string, delimeter string) []string
    array1 := strings.Split("lu min at ion", " ")
    fmt.Println(array1[1],array1[3])

    //ToLower(stringToLowercase string) string
    fmt.Println(strings.ToLower("Brandy, you're a fine girl.") )

}
