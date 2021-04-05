# class

<b>Version:</b> v0.2.0

A utility program to associate extra information with files, including a name, tags, and key/value data.

## Notice

This program has been tested, but complete stability is not guaranteed.
I'm not to be held responsible for any lost data.

## Build Instructions

- Install Go from https://golang.org/dl (tested: v1.15.8)
- `git clone https://github.com/Zytekaron/class`
- `cd class`
- `go build -o class main.go`
- that's it, folks

## Library Basics

This is currently untested, and may not work.

`go get github.com/zytekaron/class`

```go
package main

import (
	"github.com/zytekaron/class/v1/db"
	"log"
)

func main() {
	// Opens the database connection internally, with
	// the storage directory being a local .class folder
	database := db.NewBadger(".class")
	err := database.Open()
	if err != nil {
		log.Fatal(err)
	}
	// Closes the database connection internally
	defer database.Close()

	class, err := database.Get("123")
	// ...
}
```

## Future Plans

- This entire project will probably undergo a name change in the near future. If you have any ideas, let me know.

## License

<b>class</b> is licensed under the [MIT License](https://github.com/Zytekaron/class/blob/master/LICENSE)
