```go
package main

import (
    "fmt"
    "log"
    "net/http"

    "github.com/ShadowUser17/httprouter"
)

func handler(rw http.ResponseWriter, req *http.Request, out *log.Logger) {
    rw.WriteHeader(http.StatusOK)
    fmt.Fprintf(rw, "%s %s", req.Method, req.URL)
    out.Printf("%s %s", req.Method, req.URL)
}

func main() {
    var mux = httprouter.New(nil)
    mux.HandleFunc("/", http.MethodHead, handler)
    mux.HandleFunc("/", http.MethodGet, handler)

    log.Fatal(http.ListenAndServe(":8080", mux))
}
```
