package cli

import (
	"fmt"
	"github.com/zytekaron/class/v1/types"
)

func classCreate(id string) error {
	err := database.Insert(types.NewClass(id))
	if err != nil {
		return err
	}

	fmt.Println("Created class:", id)
	return nil
}
