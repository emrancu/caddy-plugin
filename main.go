package main

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/caddyserver/caddy/v2/modules/caddyhttp"
	"github.com/caddyserver/caddy/v2/modules/caddyhttp/caddylocale"
)

func init() {
	caddyhttp.RegisterModule(ExtraString{})
	caddyhttp.RegisterModule(caddylocale.Localization{})
	caddyhttp.RegisterHandlerConstructor("extra_string", parseExtraString)
}

// ExtraString represents the custom handler that adds an extra string to the response body.
type ExtraString struct {
	Next   caddyhttp.Handler
	String string `json:"string,omitempty"`
}

// ServeHTTP implements the caddyhttp.Handler interface.
func (es ExtraString) ServeHTTP(w http.ResponseWriter, r *http.Request) error {
	rec := w.(caddyhttp.ResponseWriter)
	rec.Header().Add("X-Extra-String", es.String)
	return es.Next.ServeHTTP(rec, r)
}

// parseExtraString creates a handler instance based on the provided config.
func parseExtraString(h caddyhttp.Helper) (caddyhttp.Handler, error) {
	var es ExtraString
	err := es.UnmarshalCaddyfile(h.Dispenser)
	return es, err
}

// UnmarshalCaddyfile sets up the handler from Caddyfile tokens.
func (es *ExtraString) UnmarshalCaddyfile(d *caddyfile.Dispenser) error {
	for d.Next() {
		if d.NextArg() {
			es.String = d.Val()
		}
	}
	return nil
}
