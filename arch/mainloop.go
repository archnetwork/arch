package arch

import (
	"net/http"

	"github.com/dimfeld/httptreemux/v5"
)

// handleWithWay handles all requests as defined by a 'Way'
func handleWithWay(way Way) httptreemux.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request, params map[string]string) {
		ctx := &Context{}
		var err error

		ctx.Analysis, err = way.Analyze(r, ctx)
		if err != nil {
			return
		}

		ctx.Route, err = way.Route(r, ctx)
		if err != nil {
			return
		}

		ctx.Cache, err = way.Cache(r, ctx)
		if err != nil {
			return
		}

		resp, err := way.Handle(r, ctx)
		if err != nil {
			return
		}

		for k, v := range resp.headers {
			w.Header().Set(k, v)
		}

		w.Write(resp.body)
		w.WriteHeader(resp.status)
	}
}
