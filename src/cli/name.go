package cli

import (
	"class/src/class"
	"fmt"
	"strings"
)

func classSetName(id string, args []string) error {
	name := strings.Join(args, " ")

	err := class.SetName(id, name)
	if err != nil {
		return err
	}

	fmt.Println("Updated class name to:", name)
	return nil
}
