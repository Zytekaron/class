package cli

import (
	"fmt"
	"github.com/zytekaron/class/v1/class"
)

func classDelete(id string) error {
	err := class.Delete(id)
	if err != nil {
		return err
	}

	fmt.Println("Deleted class:", id)
	return nil
}
