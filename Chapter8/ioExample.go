package main

import (
    "fmt"
    "os"
    "io/ioutil"
)

func main() {
    attempt1()
    attempt2()
    create("test2.txt")
    readFileNamesFromDir(".")
}

//My super lame attempt at guessing how to read from a file
func attempt1() {
    fmt.Println("Opening file...")
    file, err := os.Open("./test1.txt")
    defer func () {
        fmt.Println("Closing file...")
        file.Close()
    }()
    //Risk of buffer overflow -_-
    //REALLY bad code... but I'm also guessing at a new lang without Google...
    //Runs... but with errors... (11 <nil> <nil>)
    var bs []byte = []byte("                                                                                     ")
    intofsomething, readErr := file.Read(bs)
    fmt.Println(string(bs),intofsomething,readErr, err)
}

//Actually reading how to properly read from a file(wow that was simple!
func attempt2() {
    bs, err := ioutil.ReadFile("test1.txt")
    if err != nil {
        //handle the error here!  I'll learn more about this next chapter.
        return
    }
    str := string(bs)
    fmt.Println(str)
}

func create(filename string) {
    file, err := os.Create(filename)
    if err != nil {
        fmt.Println("there was an error")
        panic(err)
    }
    defer file.Close()
    file.WriteString("Hellowwwwwww")

    
}

func readFileNamesFromDir(dirName string) {
    dir, err := os.Open(dirName)
    if err != nil {
        fmt.Println("there was an error")
        panic(err)
    }
    fileInfos, err := dir.Readdir(-1) 
    if err != nil {
        fmt.Println("there was an error")
        panic(err)
    }

    for _, fileInfo := range fileInfos {
        fmt.Println( fileInfo.Name())
    }

    fmt.Printf("%v",fileInfos)
    defer dir.Close()
}


