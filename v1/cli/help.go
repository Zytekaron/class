package cli

import "fmt"

func classVersion() {
	fmt.Println("Class version: v" + version)
}

func classHelp() {
	fmt.Println("Class v" + version)
	fmt.Println("  class view <id>")
	fmt.Println("  class create <id>")
	fmt.Println("  class delete <id> <desc>")
	fmt.Println("  class name <id> <name>")
	fmt.Println("  class addtag <id> <tags...>")
	fmt.Println("  class deltag <id> <tags...>")
	fmt.Println("  class addmeta <id> <desc>")
	fmt.Println("  class delmeta <id> <desc>")
}
