package main

import (
	"bufio"
	"class/src/cli"
	"class/src/storage"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	cli.Init("0.1.0")

	err := storage.Init()
	if err != nil {
		log.Fatal(err)
	}
	defer storage.Close()

	if len(os.Args) == 1 {
		repl()
		return
	}

	err = run(os.Args[1:])
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(2)
	}
}

func repl() {
	fmt.Print("> ")
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		input := scanner.Text()

		if strings.ToLower(input) == "exit" {
			os.Exit(0)
		}

		err := run(strings.Split(input, " "))
		if err != nil {
			fmt.Println(err.Error())
		}

		fmt.Print("\n> ")
	}
}

func run(input []string) error {
	if len(input) == 1 {
		return cli.Run(input[0], []string{})
	} else {
		return cli.Run(input[0], input[1:])
	}
}
