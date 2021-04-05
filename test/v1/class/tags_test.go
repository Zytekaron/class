package class

import (
	"testing"
)

// Depends on TestCreate, types.TestClassTags
func TestTags(t *testing.T) {
	WithClass()
	defer Cleanup()

	t1, t2, t3, t4 := "tag1", "tag2", "tag 3", "tag 4"
	all := []string{t1, t2, t3, t4}

	// Add tags

	err := database.AddTag(testID, t1)
	if err != nil {
		t.Error(err)
	}

	err = database.AddTag(testID, t2)
	if err != nil {
		t.Error(err)
	}

	err = database.AddTags(testID, []string{t3, t4})
	if err != nil {
		t.Error(err)
	}

	// Ensure all tags exist

	c, err := database.Get(testID)
	if err != nil {
		t.Error(err)
	}

	if !c.Tags.HasAll(all) {
		t.Error("does not contain all tags:", c.Tags, all)
	}

	// Remove tags

	err = database.RemoveTag(testID, t1)
	if err != nil {
		t.Error(err)
	}

	err = database.RemoveTag(testID, t2)
	if err != nil {
		t.Error(err)
	}

	err = database.RemoveTags(testID, []string{t3, t4})
	if err != nil {
		t.Error(err)
	}

	// Ensure no tags exist

	c, err = database.Get(testID)
	if err != nil {
		t.Error(err)
	}
	for range c.Tags {
		t.Error("tags should be empty, instead found:", c.Tags)
		break
	}
}
