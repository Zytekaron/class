# class

<b>Version:</b> v0.1.2

A utility program to associate extra information with files,
including a name, description, tags, and key/value data.



## Notice

This program has undergone basic testing, but may be unstable. Use it with caution.

I also suggest you wait until the v1.0.0 release, as there may be drastic changes to how classes are stored by then.


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
    "github.com/zytekaron/class/v1/class"
    "log"
)

func main() {
    err := class.Open() // Opens the database connection internally
    if err != nil {
    	log.Fatal(err)
    }
    defer class.Close() // Closes the database connection internally
    
    c, err := class.Get("123")
    // ...
}
```


## Future Plans

- All metadata will be queryable, which will allow efficient searching based on tags and other data.
- Configuration may see some changes or new options to allow central configuration instead of per-directory.
  This will allow for web servers that can access the entirety of the data, as opposed to one directory at a time.
  This will also make querying more encompassing.
- This entire project will probably undergo a name change in the near future.
If you have any ideas, let me know.



## License

<b>class</b> is licensed under the [MIT License](https://github.com/Zytekaron/class/blob/master/LICENSE)
