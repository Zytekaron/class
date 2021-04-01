package cli

import (
	"class/src/storage"
	"class/src/types"
	"fmt"
	"os"
	"strings"
)

func classAddMeta(id string, args []string) {
	if len(args) < 2 {
		fmt.Println("A key and value must be provided.")
		os.Exit(2)
	}
	key := args[0]
	value := strings.Join(args[1:], " ")

	storage.EnsureHandle(id, func(class *types.Class) bool {
		class.Meta[key] = value
		return true
	})
	fmt.Println("Updated class meta.")
	fmt.Println("Key:", key)
	fmt.Println("Value:", value)
}

func classRemoveMeta(id string, args []string) {
	if len(args) == 0 {
		fmt.Println("A key must be provided.")
		os.Exit(2)
	}
	key := args[0]

	storage.EnsureHandle(id, func(class *types.Class) bool {
		delete(class.Meta, key)
		return true
	})
	fmt.Println("Updated class meta; deleted key:", key)
}
