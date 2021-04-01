package main

import (
	"class/src/cli"
	"class/src/storage"
	"fmt"
	"log"
	"os"
)

func main() {
	cli.Init("1.0")

	db, err := storage.Init()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	if len(os.Args) == 1 {
		fmt.Println("usage: class <action> <id> [data...]")
		fmt.Println("try `class help`")
		os.Exit(2)
	}

	if len(os.Args) == 2 {
		cli.Run(os.Args[1], []string{})
	}

	cli.Run(os.Args[1], os.Args[2:])
}
