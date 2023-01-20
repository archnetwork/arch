package arch

import (
	"net/http"
	"strings"

	"github.com/dimfeld/httptreemux/v5"
)

type Server struct {
	router *httptreemux.TreeMux
}

type App interface {
	Config() *Config
	Routes() map[string]Way
}

type Config struct{}

type route struct {
	method string
	path   string
}

func New(app App) *Server {
	router := httptreemux.New()

	routes := app.Routes()

	for r, way := range routes {
		route := parseRouteString(r)

		router.Handle(route.method, route.path, handleWithWay(way))
	}

	s := &Server{
		router: httptreemux.New(),
	}

	return s
}

// Start starts the Server
func (s *Server) Start() error {
	srv := http.Server{
		Addr:    ":8080",
		Handler: s.router,
	}

	return srv.ListenAndServe()
}

func parseRouteString(routeString string) route {
	method := "GET"
	path := ""

	if strings.HasPrefix(routeString, "GET") {
		path = strings.TrimPrefix(routeString, "GET ")
	}

	r := route{
		method: method,
		path:   path,
	}

	return r
}
