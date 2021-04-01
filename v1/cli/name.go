package cli

import (
	"fmt"
	"github.com/zytekaron/class/v1/class"
	"strings"
)

func classSetName(id string, args []string) error {
	name := strings.Join(args, " ")

	err := class.SetName(id, name)
	if err != nil {
		return err
	}

	fmt.Println("Updated name:", name)
	return nil
}
