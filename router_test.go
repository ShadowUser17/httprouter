package httprouter

import (
	"fmt"
	"log"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func handler(rw http.ResponseWriter, req *http.Request, out *log.Logger) {
	rw.WriteHeader(http.StatusOK)
	fmt.Fprintf(rw, "%s %s", req.Method, req.URL)
}

func TestRouter(t *testing.T) {
	var mux = NewRouter(nil)
	mux.HandleFunc("/", http.MethodHead, handler)
	mux.HandleFunc("/", http.MethodGet, handler)
	//mux.HandleFunc("/", http.MethodPost, handler)

	var server = httptest.NewUnstartedServer(mux)
	server.Start()

	var client = server.Client()

	if resp, err := client.Head(server.URL); err != nil {
		t.Errorf("TestWebServer: %v\n", err)

	} else {
		t.Logf("TestWebServer: %s %s %s\n", resp.Request.Method, resp.Request.URL, resp.Status)
		resp.Body.Close()
	}

	if resp, err := client.Get(server.URL); err != nil {
		t.Errorf("TestWebServer: %v\n", err)

	} else {
		t.Logf("TestWebServer: %s %s %s\n", resp.Request.Method, resp.Request.URL, resp.Status)
		resp.Body.Close()
	}

	if resp, err := client.Post(server.URL, "", strings.NewReader("POST...")); err != nil {
		t.Errorf("TestWebServer: %v\n", err)

	} else {
		t.Logf("TestWebServer: %s %s %s\n", resp.Request.Method, resp.Request.URL, resp.Status)
		resp.Body.Close()
	}
}
