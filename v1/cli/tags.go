package cli

import (
	"errors"
	"fmt"
	"strings"
)

func classAddTags(id string, args []string) error {
	if len(args) == 0 {
		return errors.New("at least one tag needs to be defined")
	}

	err := database.AddTags(id, args)
	if err != nil {
		return err
	}

	fmt.Println("Added tags:", strings.Join(args, ", "))
	return nil
}

func classRemoveTags(id string, args []string) error {
	if len(args) == 0 {
		return errors.New("at least one tag needs to be defined")
	}

	err := database.RemoveTags(id, args)
	if err != nil {
		return err
	}

	fmt.Println("Removed tags:", strings.Join(args, ", "))
	return nil
}
