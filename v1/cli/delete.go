package cli

import (
	"fmt"
)

func classDelete(id string) error {
	err := database.Delete(id)
	if err != nil {
		return err
	}

	fmt.Println("Deleted class:", id)
	return nil
}
