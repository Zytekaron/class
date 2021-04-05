package main

import (
	"bufio"
	"fmt"
	"github.com/spf13/pflag"
	"github.com/zytekaron/class/v1/cli"
	"github.com/zytekaron/class/v1/db"
	"log"
	"os"
	"strings"
)

var (
	dbType string
	// badger
	dir string
	// mongo
	uri     string
	dbName  string
	colName string
)

func init() {
	pflag.StringVarP(&dbType, "type", "t", "badger", "The type of database to use (badger, mongodb)")
	pflag.StringVarP(&dir, "dir", "d", ".class", "The storage directory (badger only)")
	pflag.StringVarP(&uri, "uri", "u", "", "The connection uri (mongodb only)")
	pflag.StringVarP(&dbName, "dbName", "n", "", "The database name (mongodb only)")
	pflag.StringVarP(&colName, "colName", "c", "", "The collection name (mongodb only)")
	pflag.Parse()
}

func main() {
	dbase := initDb()
	if err := dbase.Open(); err != nil {
		log.Fatal(err)
	}
	defer dbase.Close()

	cli.Init(dbase, "0.2.0")

	args := pflag.Args()
	if len(args) == 0 {
		repl()
		return
	}
	if err := run(args); err != nil {
		fmt.Println(err.Error())
		os.Exit(2)
	}
}

func initDb() db.Database {
	switch dbType {
	case "badger":
		if dir == "" {
			log.Println("dir must be provided")
		}
		return db.NewBadger(dir)
	case "mongo", "mongodb":
		if uri == "" || dbName == "" || colName == "" {
			log.Fatal("uri, dbName, and colName must be provided")
		}
		return db.NewMongo(uri, dbName, colName)
	default:
		log.Fatal("db type must be one of: badger, mongodb")
		return nil
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
