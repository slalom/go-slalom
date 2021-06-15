# Go Hello World

Create a working web app with the code below!

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

- create a `hello.go` file using below

```bash
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

- run the following command and leave console open

```bash
go run hello.go
```

- in browser go to (replace YOUR_NAME with your name) `http://localhost:8080/YOUR_NAME` or in separate console

```bash
open http://localhost:8080/$USER
```

![hello](screens/hello-screen.png)

## Awesomeness

So what did you do? When you ran `go run hello.go`, go compiled hello.go, created a binary in a temporary folder, and ran it.

Alternatively you could build and run it with

```bash
go build hello.go
./hello
```

In above, go compiled hello.go and created a binary in current path. You then can run it.

### compiled server

In many other languages you need to deploy your app to a web app server or "bundle" a server with your app. With go the server is built and compiled into your app (binary).

## Next

[Run go-slalom](build-go-slalom.md)

![gopher ninja](images/gopher-ninja.png)

