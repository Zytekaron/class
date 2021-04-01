package cli

import (
	"class/src/storage"
	"class/src/types"
	"fmt"
	"strings"
)

func classSetName(id string, args []string) {
	name := strings.Join(args, " ")
	storage.EnsureHandle(id, func(class *types.Class) bool {
		class.Name = name
		return true
	})
	fmt.Println("Updated class name to:", name)
}
