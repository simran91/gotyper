package typetext

import "fmt"
import "time"

// TypeText writes out the string iterating over each byte
// (ranging from the lowest byte (based on input) to the current byte)...
func TypeText(sleepTime time.Duration, str string) bool {

	// a string is just a slice of bytes... but make it explicit...
	bites := []byte(str)

	// output is what we will be writing to the screen (building it up as wel go...)
	output := []byte{}

	// the smallest byte in the input...
	var smallestByte byte

	// get the smallest and biggest bytes so that we know what to "range" over later...
	for _, b := range bites {
		if smallestByte == 0 {
			smallestByte = b
		} else if b < smallestByte {
			smallestByte = b
		}
	}

	// range over all the bytes and print them out...
	for _, b := range bites {
		output = append(output, b)
		position := len(output)

		for i := smallestByte; i <= b; i++ {
			output[position-1] = i
			fmt.Printf("%v\r", string(output))
			time.Sleep(sleepTime)
		}
	}

	// return true that all succeeded... not much use at this point, but just planning for the future when we can test
	// if we were writing to a valid terminal, etc...
	return true
}
