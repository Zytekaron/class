package cli

import (
	"errors"
	"github.com/zytekaron/class/v1/db"
)

var database db.Database
var version string

func Init(db db.Database, ver string) {
	database = db
	version = ver
}

func Run(command string, input []string) error {
	switch command {
	case "?", "h", "-h", "help":
		classHelp()
	case "-v", "version", "--version":
		classVersion()
	case "v", "view", "show":
		id, _, err := construct(input)
		if err != nil {
			return err
		}
		return classView(id)
	case "c", "new", "create":
		id, _, err := construct(input)
		if err != nil {
			return err
		}
		return classCreate(id)
	case "n", "sn", "name", "setname", "rename":
		id, args, err := construct(input)
		if err != nil {
			return err
		}
		return classSetName(id, args)
	case "t", "at", "tag", "tags", "addtag", "addtags":
		id, args, err := construct(input)
		if err != nil {
			return err
		}
		return classAddTags(id, args)
	case "ut", "rt", "untag", "rmtag", "rmtags", "removetag", "removetags":
		id, args, err := construct(input)
		if err != nil {
			return err
		}
		return classRemoveTags(id, args)
	case "m", "am", "meta", "addmeta":
		id, args, err := construct(input)
		if err != nil {
			return err
		}
		return classAddMeta(id, args)
	case "um", "rm", "dm", "unmeta", "delmeta", "rmeta", "rmmeta", "removemeta", "deletemeta":
		id, args, err := construct(input)
		if err != nil {
			return err
		}
		return classRemoveMeta(id, args)
	case "d", "r", "del", "delete", "remove":
		id, _, err := construct(input)
		if err != nil {
			return err
		}
		return classDelete(id)
	default:
		return errors.New("invalid command: " + command)
	}
	return nil
}

func construct(input []string) (string, []string, error) {
	if len(input) == 0 {
		return "", nil, errors.New("a class name must be specified")
	}
	if len(input) == 1 {
		return input[0], []string{}, nil
	}
	return input[0], input[1:], nil
}
