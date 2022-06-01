package httprouter

import (
	"net/http/httptest"
	"strings"
	"testing"
)

func TestRouter(t *testing.T) {
	var mux = New(nil)
	mux.Head("/", DefaultHandler)
	mux.Get("/", DefaultHandler)

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
