package cli

import (
	"fmt"
	"strings"
)

func classView(id string) error {
	c, err := database.Get(id)
	if err != nil {
		return err
	}

	fmt.Println("ID:  ", c.ID)
	fmt.Println("Name:", c.Name)
	fmt.Println("Tags:", strings.Join(c.Tags, ", "))
	fmt.Println("Meta:")
	for key, value := range c.Meta {
		fmt.Println(key, "=>", value)
	}

	return nil
}
