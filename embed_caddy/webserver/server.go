package webserver

import (
	"io/ioutil"
	"log"
	"os"

	"github.com/mholt/caddy"
	_ "github.com/mholt/caddy/caddyhttp"
)

func init() {
	caddy.SetDefaultCaddyfileLoader("default", caddy.LoaderFunc(defaultLoader))
}

func defaultLoader(serverType string) (caddy.Input, error) {
	contents, err := ioutil.ReadFile(caddy.DefaultConfigFile)
	if err != nil {
		if os.IsNotExist(err) {
			return nil, nil
		}
		return nil, err
	}
	log.Printf("Loading Caddyfile: %s\n", string(contents))
	log.Println(serverType)
	return caddy.CaddyfileInput{
		Contents:       contents,
		Filepath:       caddy.DefaultConfigFile,
		ServerTypeName: serverType,
	}, nil
}

func GetServer() *caddy.Instance {
	caddy.AppName = "test"
	caddy.AppVersion = "0.1"

	caddyfile, err := caddy.LoadCaddyfile("http")
	if err != nil {
		log.Fatalf("Unable to load caddyfile: %s", err)
	}

	inst, err := caddy.Start(caddyfile)
	if err != nil {
		log.Fatalf("Unable to start caddy server: %s", err)
	}
	return inst
}
