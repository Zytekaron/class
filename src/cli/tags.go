package cli

import (
	"class/src/storage"
	"class/src/types"
	"fmt"
	"os"
	"strings"
)

func classAddTags(id string, args []string) {
	ensureTags(args)

	var existingTags []string
	storage.EnsureHandle(id, func(class *types.Class) bool {
		for _, tag := range args {
			class.Tags.Add(tag)
		}
		existingTags = class.Tags
		return true
	})
	fmt.Println("Added tags:", strings.Join(args, ", "))
	fmt.Println("Current tags:", existingTags)
}

func classRemoveTags(id string, args []string) {
	ensureTags(args)

	var existingTags []string
	storage.Handle(id, func(class *types.Class) bool {
		for _, tag := range args {
			class.Tags.Remove(tag)
		}
		existingTags = class.Tags
		return true
	})
	fmt.Println("Removed tags:", strings.Join(args, ", "))
	fmt.Println("Current tags:", existingTags)
}

func hasAnyTag(id string, args []string) {
	storage.Handle(id, func(class *types.Class) bool {
		if class.Tags.HasAny(args) {
			fmt.Println("true")
		} else {
			fmt.Println("false")
		}
		return true
	})
}

func hasAllTags(id string, args []string) {
	storage.Handle(id, func(class *types.Class) bool {
		if class.Tags.HasAll(args) {
			fmt.Println("true")
		} else {
			fmt.Println("false")
		}
		return true
	})
}

// todo sort by most intersecting tags!

func ensureTags(args []string) {
	if len(args) == 0 {
		fmt.Println("At least one tag must be specified.")
		os.Exit(2)
	}
}
