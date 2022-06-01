```go
package main

import (
	"log"
	"net/http"

	"github.com/ShadowUser17/httprouter"
)

func Handler(rw http.ResponseWriter, req *http.Request, out *log.Logger) {
	rw.WriteHeader(http.StatusOK)
	out.Printf("%s %s", req.Method, req.URL)
}

func main() {
	var mux = httprouter.New(nil)
	mux.Head("/", Handler)
	mux.Get("/", Handler)

	var out = mux.GetLogger()
	var srv = http.Server{
		Addr:     ":8080",
		Handler:  mux,
		ErrorLog: out,
	}

	out.Fatal(srv.ListenAndServe())
}
```
