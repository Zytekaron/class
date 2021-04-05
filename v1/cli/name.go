package cli

import (
	"fmt"
	"strings"
)

func classSetName(id string, args []string) error {
	name := strings.Join(args, " ")

	err := database.SetName(id, name)
	if err != nil {
		return err
	}

	fmt.Println("Updated name:", name)
	return nil
}
