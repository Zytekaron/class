package types

import (
	"github.com/zytekaron/class/v1/types"
	"testing"
)

var testID = "__test__"

func TestClass(t *testing.T) {
	c := types.NewClass(testID)
	if c.ID != testID {
		t.Error("class id does not match:", c.ID+",", testID)
	}
}

func TestClassTags(t *testing.T) {
	c := types.NewClass(testID)

	c.Tags.Add("1")
	c.Tags.Add("2")
	c.Tags.Add("3")

	if !c.Tags.HasAny([]string{"2"}) {
		t.Error("tags failed HasAll check with 2")
	}
	if !c.Tags.HasAny([]string{"5", "4", "3"}) {
		t.Error("tags failed HasAll check with 5, 4, 3")
	}
	if c.Tags.HasAny([]string{"6", "7"}) {
		t.Error("tags failed HasAll check with 6, 7")
	}

	if !c.Tags.HasAll([]string{"2"}) {
		t.Error("tags failed HasAll check with 2")
	}
	if !c.Tags.HasAll([]string{"1", "2", "3"}) {
		t.Error("tags failed HasAll check with 1, 2, 3")
	}
	if c.Tags.HasAll([]string{"3", "4"}) {
		t.Error("tags failed HasAll check with 1, 2, 3")
	}

	c.Tags.Remove("3")
	if c.Tags.Has("3") {
		t.Error("tags contains 3 after removal of 3")
	}
	if !c.Tags.HasAll([]string{"1", "2"}) {
		t.Error("tags does not contain 1, 2 after removal of 3")
	}
}
