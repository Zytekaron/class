package cli

import "fmt"

func classVersion() {
	fmt.Println("Class version: v" + version)
}

func classHelp() {
	fmt.Println("Class v" + version)
	fmt.Println("`class view <id>`")
	fmt.Println("`class create <id>`")
	fmt.Println("`class delete <id>`")
	fmt.Println("`class name <id> <name>`")
	fmt.Println("`class desc <id> <desc>`")
	fmt.Println("`class tag <id> <tags...>`")
	fmt.Println("`class untag <id> <tags...>`")
	fmt.Println("`class addmeta <id> <desc>`")
	fmt.Println("`class delmeta <id> <desc>`")
}
