# Go Hello World 

Create a full working web server with the code below!

```go
package main

import (
    "fmt"
    "log"
    "net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Hello, World %s!", r.URL.Path[1:])
}

func main() {
    http.HandleFunc("/", handler)
    log.Fatal(http.ListenAndServe(":8080", nil))
}
```

## Steps

- copy the code above into a file `hello.go`
- run the following command and leave command window open
```bash
go run hello.go
```
- in browser go to (replace YOUR_NAME with your name) `http://localhost:8080/YOUR_NAME`
```bash
open http://localhost:8080/$USER
```

![hello](screens/hello-screen.png)

## Awesomeness!

So what did you do? When you ran `go run hello.go`, go compiled hello.go, created a binary in a temporary folder, and ran it.

### Some Basics

**package main**
Packages in go are similar to Java and other languages for structuring code. Executable commands must be in the `main` package.