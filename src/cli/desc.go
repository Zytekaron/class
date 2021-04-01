package cli

import (
	"class/src/class"
	"fmt"
	"strings"
)

func classSetDesc(id string, args []string) error {
	desc := strings.Join(args, " ")
	err := class.SetDesc(id, desc)
	if err != nil {
		return err
	}

	fmt.Println("Updated class description to:", desc)
	return nil
}
