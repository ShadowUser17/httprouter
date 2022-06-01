package httprouter

import (
	"log"
	"net/http"
	"os"
	"sync"
)

type HandlerFunc func(http.ResponseWriter, *http.Request, *log.Logger)

type Endpoint map[string]HandlerFunc

type Router struct {
	mutex     *sync.Mutex
	logger    *log.Logger
	endpoints map[string]Endpoint
}

func New(logger *log.Logger) *Router {
	if logger == nil {
		logger = log.New(os.Stdout, "", log.Ldate|log.Ltime)
	}

	return &Router{
		mutex:     new(sync.Mutex),
		logger:    logger,
		endpoints: make(map[string]Endpoint),
	}
}

func DefaultHandler(rw http.ResponseWriter, req *http.Request, out *log.Logger) {
	rw.WriteHeader(http.StatusOK)
	rw.Header().Set("Content-Type", "text/html")
	out.Printf("%s %s %s", req.Method, req.URL, req.Proto)
}

func (router *Router) ServeHTTP(rw http.ResponseWriter, req *http.Request) {
	router.mutex.Lock()
	defer router.mutex.Unlock()

	if endpoint, ok := router.endpoints[req.URL.EscapedPath()]; !ok {
		rw.WriteHeader(http.StatusNotFound)

	} else if handler, ok := endpoint[req.Method]; !ok {
		rw.WriteHeader(http.StatusMethodNotAllowed)

	} else {
		handler(rw, req, router.logger)
	}
}

func (router *Router) HandleFunc(location, method string, handler HandlerFunc) {
	if _, ok := router.endpoints[location]; !ok {
		router.endpoints[location] = make(Endpoint)
	}

	router.endpoints[location][method] = handler
}
