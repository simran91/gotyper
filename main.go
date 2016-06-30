package main

import "fmt"
import "time"
import "./typetext"

func main() {
	sleepTime := time.Duration(3500000)
	text := getData()
	typetext.TypeText(sleepTime, text)
	fmt.Println()
}

func getData() string {
	str := "This is a test, this is only a test!"
	return str
}
