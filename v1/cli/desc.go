package cli

import (
	"fmt"
	"github.com/zytekaron/class/v1/class"
	"strings"
)

func classSetDesc(id string, args []string) error {
	desc := strings.Join(args, " ")
	err := class.SetDesc(id, desc)
	if err != nil {
		return err
	}

	fmt.Println("Updated desc to:", desc)
	return nil
}
