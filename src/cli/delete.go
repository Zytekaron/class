package cli

import (
	"class/src/class"
	"fmt"
)

func classDelete(id string) error {
	err := class.Delete(id)
	if err != nil {
		return err
	}

	fmt.Println("Deleted class:", id)
	return nil
}
