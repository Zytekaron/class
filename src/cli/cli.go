package cli

import (
	"fmt"
	"os"
)

var version string

func Init(ver string) {
	version = ver
}

func Run(command string, input []string) {
	switch command {
	case "?", "h", "-h", "help":
		classHelp()
	case "-v", "version", "--version":
		classVersion()
	case "n", "sn", "name", "setname", "rename":
		id, args := construct(input)
		classSetName(id, args)
	case "d", "sd", "desc", "setdesc", "description":
		id, args := construct(input)
		classSetDesc(id, args)
	case "t", "at", "tag", "tags", "addtag", "addtags":
		id, args := construct(input)
		classAddTags(id, args)
	case "ut", "rt", "untag", "rmtag", "rmtags", "removetag", "removetags":
		id, args := construct(input)
		classRemoveTags(id, args)
	case "m", "am", "meta", "addmeta":
		id, args := construct(input)
		classAddMeta(id, args)
	case "um", "rm", "unmeta", "delmeta", "rmvmeta", "rmmeta", "removemeta":
		id, args := construct(input)
		classRemoveMeta(id, args)
	case "v", "view", "show":
		id, _ := construct(input)
		classView(id)
	default:
		fmt.Println("invalid command:", command)
		os.Exit(2)
	}
}

func construct(input []string) (string, []string) {
	if len(input) == 0 {
		fmt.Println("A class id must be specified.")
		os.Exit(2)
	}
	if len(input) == 1 {
		return input[0], []string{}
	}
	return input[0], input[1:]
}
