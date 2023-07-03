package main

import (
	"github.com/caddyserver/caddy/v2"
	"github.com/caddyserver/caddy/v2/caddyconfig/caddyfile"
)

func init() {
	caddy.RegisterModule(ExamplePlugin{})
	caddy.RegisterParser("example", caddyfile.UnmarshalCaddyfile)
}

type ExamplePlugin struct {
	Message string `json:"message,omitempty"`
}

func (p ExamplePlugin) ServeHTTP(w http.ResponseWriter, r *http.Request) error {
	// Your plugin logic goes here
	return nil
}
