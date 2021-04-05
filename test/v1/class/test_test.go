package class

import "github.com/zytekaron/class/v1/types"

// Class testID used for testing
var testID = "__test__"

func WithClass() {
	if err := database.Insert(types.NewClass(testID)); err != nil {
		panic(err)
	}
}

func Cleanup() {
	if err := database.Delete(testID); err != nil {
		panic(err)
	}
}
