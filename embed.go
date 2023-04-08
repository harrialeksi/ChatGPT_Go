/*
In Go, you can embed data and files directly into your compiled binary using the embed package introduced in Go 1.16. This makes it easy to 
distribute a single binary file without any external dependencies or configuration files.

Here's an example of embedding a file named data.txt into a Go program:

In this example, we use the embed package to embed the contents of the data.txt file into a variable named data. We then use fmt.Println to print out 
the contents of data.

Note the use of the //go:embed directive, which tells the Go compiler to include the specified file in the embedded resources of the binary.

You can also embed an entire directory of files by specifying a glob pattern:
*/

package main

import (
    "embed"
    "fmt"
)

//go:embed data.txt
var data []byte

func main() {
    fmt.Println(string(data))
}

/*
In this example, we use the embed package to embed all files in the static directory into an embed.FS type variable named staticFiles. We then use 
staticFiles.ReadFile to read the contents of a file named example.txt from the embedded directory.
*/



package main

import (
    "embed"
    "fmt"
)

//go:embed static/*
var staticFiles embed.FS

func main() {
    // Read the contents of a file from the embedded static directory
    fileContents, err := staticFiles.ReadFile("static/example.txt")
    if err != nil {
        fmt.Println(err)
        return
    }
    fmt.Println(string(fileContents))
}
