package class

import (
	"github.com/zytekaron/class/v1/class"
	"github.com/zytekaron/class/v1/types"
	"testing"
)

// Depends on TestCreate, types.TestClassTags
func TestTags(t *testing.T) {
	WithClass()
	defer Cleanup()

	t1, t2, t3, t4 := "tag1", "tag2", "tag 3", "tag 4"
	all := []string{t1, t2, t3, t4}

	// Add tags manually

	err := class.AddTag(testID, t1)
	if err != nil {
		t.Error(err)
	}

	err = class.AddTag(testID, t2)
	if err != nil {
		t.Error(err)
	}

	// Bulk add tags

	err = class.AddTags(testID, []string{t3, t4})
	if err != nil {
		t.Error(err)
	}

	// Ensure all tags exist

	tags, err := class.GetTags(testID)
	if err != nil {
		t.Error(err)
	}
	tl := types.TagList(tags)
	if !tl.HasAll(all) {
		t.Error("does not contain all tags:", tags, all)
	}

	// Remove tags manually

	err = class.RemoveTag(testID, t1)
	if err != nil {
		t.Error(err)
	}

	err = class.RemoveTag(testID, t2)
	if err != nil {
		t.Error(err)
	}

	// Bulk remove tags

	err = class.RemoveTags(testID, []string{t3, t4})
	if err != nil {
		t.Error(err)
	}

	// Ensure no tags exist

	tags, err = class.GetTags(testID)
	if err != nil {
		t.Error(err)
	}
	for range tags {
		t.Error("tags should be empty, instead found:", tags)
		break
	}
}
