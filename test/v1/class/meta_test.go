package class

import (
	"testing"
)

// Depends on TestCreate
func TestMeta(t *testing.T) {
	WithClass()
	defer Cleanup()

	k1, v1 := "key1", "value1"
	k2, v2 := "key2", "value2"
	k3, v3 := "key3", "value3"
	k4, v4 := "key4", "value4 long"

	// Add meta

	err := database.AddMeta(testID, k1, v1)
	if err != nil {
		t.Error(err)
	}

	err = database.AddMeta(testID, k2, v2)
	if err != nil {
		t.Error(err)
	}

	err = database.AddMetaBulk(testID, map[string]interface{}{k3: v3, k4: v4})
	if err != nil {
		t.Error(err)
	}

	// Ensure all meta exists

	c, err := database.Get(testID)
	if err != nil {
		t.Error(err)
	}
	meta := c.Meta
	for key, value := range map[string]string{k1: v1, k2: v2, k3: v3, k4: v4} {
		valueOut, ok := meta[key]
		if !ok {
			t.Error("missing key:", key)
		}
		if valueOut != value {
			t.Error("mismatched value for key:", key, "|", valueOut, "|", value)
		}
	}

	// Remove meta

	err = database.RemoveMeta(testID, k1)
	if err != nil {
		t.Error(err)
	}

	err = database.RemoveMeta(testID, k2)
	if err != nil {
		t.Error(err)
	}

	err = database.RemoveMetaBulk(testID, []string{k3, k4})
	if err != nil {
		t.Error(err)
	}

	// Ensure no meta exists

	c, err = database.Get(testID)
	if err != nil {
		t.Error(err)
	}
	meta = c.Meta
	for range meta {
		t.Error("meta should be empty, instead found:", meta)
		break
	}
}
