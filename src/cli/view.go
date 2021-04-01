package cli

import (
	"class/src/storage"
	"class/src/types"
	"fmt"
	"strings"
)

func classView(id string) {
	exists := storage.Handle(id, func(class *types.Class) bool {
		fmt.Println("ID:  ", class.ID)
		fmt.Println("Name:", class.Name)
		fmt.Println("Desc:", class.Desc)
		fmt.Println("Tags:", strings.Join(class.Tags, ", "))
		fmt.Println("Meta:")
		for key, value := range class.Meta {
			fmt.Println(key, "=>", value)
		}
		return false
	})
	if !exists {
		fmt.Println("A class by that id does not exist.")
	}
}
