package class

import (
	"github.com/zytekaron/class/v1/class"
	"testing"
)

// Depends on TestCreate
func TestDesc(t *testing.T) {
	WithClass()
	defer Cleanup()

	name := "Class description example"
	err := class.SetDesc(testID, name)
	if err != nil {
		t.Error(err)
	}

	descOut, err := class.GetDesc(testID)
	if err != nil {
		t.Error(err)
	}

	if descOut != name {
		t.Error("class descriptions do not match:", descOut+",", name)
	}
}
