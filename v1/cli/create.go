package cli

import (
	"fmt"
	"github.com/zytekaron/class/v1/class"
)

func classCreate(id string) error {
	_, err := class.Create(id)
	if err != nil {
		return err
	}

	fmt.Println("Created class:", id)
	return nil
}
