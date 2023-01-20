package arch

import "net/http"

// Way is an object that describes the **way** to handle a request
type Way interface {
	Analyze(req *http.Request, ctx *Context) (*Analysis, error)
	Route(req *http.Request, ctx *Context) (*Route, error)
	Cache(req *http.Request, ctx *Context) (*Cache, error)
	Handle(req *http.Request, ctx *Context) (*Response, error)
}
