package main

import "fmt"

func main() {
    text := getData()
    typeText(text)
}

func getData() string {
    str := "This is a test, this is only a test!"
    return str
}

func typeText(str string) {
    fmt.Println(str)
}
