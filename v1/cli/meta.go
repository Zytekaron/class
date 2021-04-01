package cli

import (
	"errors"
	"fmt"
	"github.com/zytekaron/class/v1/class"
	"strings"
)

func classAddMeta(id string, args []string) error {
	if len(args) < 2 {
		return errors.New("a key and value must be provided")
	}
	key := args[0]
	value := strings.Join(args[1:], " ")

	err := class.AddMeta(id, key, value)
	if err != nil {
		return err
	}

	fmt.Println("Added meta:", key, "=>", value)
	return nil
}

func classRemoveMeta(id string, args []string) error {
	if len(args) == 0 {
		return errors.New("a key must be provided")
	}
	key := args[0]

	err := class.RemoveMeta(id, key)
	if err != nil {
		return err
	}

	fmt.Println("Deleted meta:", key)
	return nil
}
