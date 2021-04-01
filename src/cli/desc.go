package cli

import (
	"class/src/storage"
	"class/src/types"
	"fmt"
	"strings"
)

func classSetDesc(id string, args []string) {
	desc := strings.Join(args, " ")
	storage.EnsureHandle(id, func(class *types.Class) bool {
		class.Desc = desc
		return true
	})
	fmt.Println("Updated class description to:", desc)
}
