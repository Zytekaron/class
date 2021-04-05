package class

import (
	"testing"
)

// Depends on TestCreate
func TestName(t *testing.T) {
	WithClass()
	defer Cleanup()

	name := "Class Name Example"
	err := database.SetName(testID, name)
	if err != nil {
		t.Error(err)
	}

	c, err := database.Get(testID)
	if err != nil {
		t.Error(err)
	}

	if c.Name != name {
		t.Error("class names do not match:", c.Name+",", name)
	}
}
