package class

import (
	"github.com/zytekaron/class/v1/class"
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

	// Add keys manually

	err := class.AddMeta(testID, k1, v1)
	if err != nil {
		t.Error(err)
	}

	err = class.AddMeta(testID, k2, v2)
	if err != nil {
		t.Error(err)
	}

	// Bulk add keys

	err = class.BulkAddMeta(testID, map[string]string{k3: v3, k4: v4})
	if err != nil {
		t.Error(err)
	}

	// Ensure all keys exist

	meta, err := class.GetMeta(testID)
	if err != nil {
		t.Error(err)
	}
	for key, value := range map[string]string{k1: v1, k2: v2, k3: v3, k4: v4} {
		valueOut, ok := meta[key]
		if !ok {
			t.Error("missing key:", key)
		}
		if valueOut != value {
			t.Error("mismatched value for key:", key, "|", valueOut, "|", value)
		}
	}

	// Remove keys manually

	err = class.RemoveMeta(testID, k1)
	if err != nil {
		t.Error(err)
	}

	err = class.RemoveMeta(testID, k2)
	if err != nil {
		t.Error(err)
	}

	// Bulk remove keys

	err = class.BulkRemoveMeta(testID, []string{k3, k4})
	if err != nil {
		t.Error(err)
	}

	// Ensure no keys exist

	meta, err = class.GetMeta(testID)
	if err != nil {
		t.Error(err)
	}
	for range meta {
		t.Error("meta should be empty, instead found:", meta)
		break
	}
}
