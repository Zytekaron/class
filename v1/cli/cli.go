package cli

import (
	"errors"
	"fmt"
	"os"
)

var version string

func Init(ver string) {
	version = ver
}

func Run(command string, input []string) error {
	switch command {
	case "?", "h", "-h", "help":
		classHelp()
	case "-v", "version", "--version":
		classVersion()
	case "v", "view", "show":
		id, _ := construct(input)
		return classView(id)
	case "c", "new", "create":
		id, _ := construct(input)
		return classCreate(id)
	case "n", "sn", "name", "setname", "rename":
		id, args := construct(input)
		return classSetName(id, args)
	case "d", "sd", "desc", "setdesc", "description":
		id, args := construct(input)
		return classSetDesc(id, args)
	case "t", "at", "tag", "tags", "addtag", "addtags":
		id, args := construct(input)
		return classAddTags(id, args)
	case "ut", "rt", "untag", "rmtag", "rmtags", "removetag", "removetags":
		id, args := construct(input)
		return classRemoveTags(id, args)
	case "m", "am", "meta", "addmeta":
		id, args := construct(input)
		return classAddMeta(id, args)
	case "um", "unmeta", "delmeta", "rmeta", "rmmeta", "removemeta":
		id, args := construct(input)
		return classRemoveMeta(id, args)
	case "r", "rm", "del", "delete", "remove":
		id, _ := construct(input)
		return classDelete(id)
	default:
		return errors.New("invalid command: " + command)
	}
	return nil
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
