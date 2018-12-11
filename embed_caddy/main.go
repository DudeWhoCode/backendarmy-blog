package main

import (
	"github.com/dudewhocode/backendarmy-blog/embed_caddy/simpleserver"
	"github.com/dudewhocode/backendarmy-blog/embed_caddy/webserver"
)

func main() {
	caddyInst := webserver.GetServer()
	simpleserver.StartHTTPServer()
	caddyInst.Wait()
}
