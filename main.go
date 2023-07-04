package main

import (
	"net/http"

	"github.com/caddyserver/caddy/v2"
	"github.com/caddyserver/caddy/v2/modules/caddyhttp"
)

type CustomHeader struct {
	Name  string `json:"name,omitempty"`
	Value string `json:"value,omitempty"`
}

func (h CustomHeader) ServeHTTP(w http.ResponseWriter, r *http.Request, next caddyhttp.Handler) error {
	w.Header().Add(h.Name, h.Value)
	return next.ServeHTTP(w, r)
}

func CustomHeaderDirective() caddyhttp.MiddlewareHandler {
	return func(next caddyhttp.Handler) caddyhttp.Handler {
		return &CustomHeader{}
	}
}

func init() {
	caddy.RegisterModule(&CustomHeader{})
	caddy.RegisterModule(CustomHeaderDirective)
}