package class

import (
	"github.com/zytekaron/class/v1/class"
	"github.com/zytekaron/class/v1/types"
	"testing"
)

func TestMain(m *testing.M) {
	if err := class.Open(); err != nil {
		panic(err)
	}
	defer func() {
		if err := class.Close(); err != nil {
			panic(err)
		}
	}()

	m.Run()
}

func TestCreate(t *testing.T) {
	defer Cleanup()

	c, err := class.Create(testID)
	if err != nil {
		t.Error(err)
	}

	if c.ID != testID {
		t.Error("testID does not match:", testID+",", c.ID)
	}
}

func TestSave(t *testing.T) {
	defer Cleanup()

	err := class.Save(types.NewClass(testID))
	if err != nil {
		t.Error(err)
	}
}

// Depends on TestCreate
func TestGet(t *testing.T) {
	WithClass()
	defer Cleanup()

	c, err := class.Get(testID)
	if err != nil {
		t.Error(err)
	}

	if c.ID != testID {
		t.Error("class testID does not match:", testID+",", c.ID)
	}
}

// Depends on TestCreate
func TestBatch(t *testing.T) {
	WithClass()
	defer Cleanup()

	res, err := class.Batch([]string{testID})
	if err != nil {
		t.Error(err)
	}

	if len(res) == 0 {
		t.Error("class not found")
		return
	}

	if len(res) > 1 {
		t.Error("too many results (this shouldn't be possible):", len(res))
	}

	if res[0].ID != testID {
		t.Error("class testID does not match:", testID+",", res[0].ID)
	}
}

// Depends on TestCreate
func TestAll(t *testing.T) {
	WithClass()
	defer Cleanup()

	res, err := class.All()
	if err != nil {
		t.Error(err)
	}

	if len(res) < 1 {
		t.Error("class not found")
		return
	}

	ok := false
	for _, c := range res {
		if c.ID == testID {
			ok = true
		}
	}
	if !ok {
		t.Error("class not present in result")
	}
}

// Depends on TestCreate, TestGet
func TestDelete(t *testing.T) {
	WithClass()
	defer Cleanup()

	err := class.Delete(testID)
	if err != nil {
		t.Error(err)
	}

	c, err := class.Get(testID)
	if c != nil {
		t.Error("class was not deleted")
	}
}
