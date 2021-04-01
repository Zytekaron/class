package cli

import (
	"class/src/class"
	"fmt"
)

func classCreate(id string) error {
	_, err := class.Create(id)
	if err != nil {
		return err
	}

	fmt.Println("Created class:", id)
	return nil
}
