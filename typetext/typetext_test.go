package typetext

import "fmt"
import "time"
import "testing"

func TestTypeText(t *testing.T) {
	retVal := TypeText(time.Duration(1000000), "This is a test")

	if !retVal {
		t.Error("Some error occured, did not get a true back!", retVal)
	}
	fmt.Println()
}
