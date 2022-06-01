```go
package main

import (
    "fmt"
    "log"
    "net/http"

    "github.com/ShadowUser17/httprouter"
)

func Handler(rw http.ResponseWriter, req *http.Request, out *log.Logger) {
    rw.WriteHeader(http.StatusOK)
    fmt.Fprintf(rw, "%s %s", req.Method, req.URL)
    out.Printf("%s %s", req.Method, req.URL)
}

func main() {
    var mux = httprouter.New(nil)
    mux.HEAD("/", Handler)
    mux.GET("/", Handler)

    log.Fatal(http.ListenAndServe(":8080", mux))
}
```
