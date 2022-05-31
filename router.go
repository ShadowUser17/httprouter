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

func NewRouter(logger *log.Logger) *Router {
	if logger == nil {
		logger = log.New(os.Stdout, "", log.Ldate|log.Ltime)
	}

	return &Router{
		mutex:     new(sync.Mutex),
		logger:    logger,
		endpoints: make(map[string]Endpoint),
	}
}

func (router *Router) GetHandler(location, method string) (HandlerFunc, int) {
	router.mutex.Lock()
	defer router.mutex.Unlock()

	if endpoint, ok := router.endpoints[location]; ok {
		if handler, ok := endpoint[method]; ok {
			return handler, http.StatusOK

		} else {
			return nil, http.StatusMethodNotAllowed
		}
	}

	return nil, http.StatusNotFound
}

func (router *Router) ServeHTTP(rw http.ResponseWriter, req *http.Request) {
	var handler, status = router.GetHandler(req.URL.EscapedPath(), req.Method)

	if status == http.StatusOK {
		handler(rw, req, router.logger)

	} else {
		rw.WriteHeader(status)
	}
}

func (router *Router) HandleFunc(location, method string, handler HandlerFunc) {
	if _, ok := router.endpoints[location]; !ok {
		router.endpoints[location] = make(Endpoint)
	}

	router.endpoints[location][method] = handler
}
