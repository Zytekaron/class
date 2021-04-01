# class

<b>Version:</b> v0.1.2

A file utility program to associate extra information with files,
including a name, description, tags, and key/value data.



## Notice

This program has not undergone strict testing and may be unstable. Use it with caution.



## Build Instructions

- Install Go from https://golang.org/dl (tested: v1.15.8)
- `go build -o class main.go`
- that's it, folks



## Library Basics

This is currently also untested, so it may not work.

`go get github.com/zytekaron/class`

```go
package main

import (
    "github.com/zytekaron/class/v1/class"
    "log"
)

func main() {
    err := class.Init() // Opens the database connection internally
    if err != nil {
    	log.Fatal(err)
    }
    defer class.Close() // Closes the database connection internally
    
    c, err := class.Get("123")
    // ...
}
```


## Future Plans

- Packages will be changed to improve quality and simplify usage.
- Formal tests will be written for all behavior.
- All metadata will be queryable, which will allow efficient searching based on tags and other data.
- This entire project will probably undergo a name change in the near future.
If you have any ideas, let me know.



## License

<b>class</b> is licensed under the [MIT License](https://github.com/Zytekaron/class/blob/master/LICENSE)
