package class

import (
	"github.com/zytekaron/class/v1/class"
)

// Class testID used for testing
var testID = "__test__"

func WithClass() {
	if _, err := class.Create(testID); err != nil {
		panic(err)
	}
}

func Cleanup() {
	if err := class.Delete(testID); err != nil {
		panic(err)
	}
}
