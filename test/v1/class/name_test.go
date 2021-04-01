package class

import (
	"github.com/zytekaron/class/v1/class"
	"testing"
)

// Depends on TestCreate
func TestName(t *testing.T) {
	WithClass()
	defer Cleanup()

	name := "Class Name Example"
	err := class.SetName(testID, name)
	if err != nil {
		t.Error(err)
	}

	n, err := class.GetName(testID)
	if err != nil {
		t.Error(err)
	}

	if n != name {
		t.Error("class names do not match:", n+",", name)
	}
}
