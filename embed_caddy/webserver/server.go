package webserver

import (
	"fmt"
	"io/ioutil"
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
	fmt.Printf("Loading Caddyfile: %s\n", string(contents))
	fmt.Println(serverType)
	return caddy.CaddyfileInput{
		Contents:       contents,
		Filepath:       caddy.DefaultConfigFile,
		ServerTypeName: serverType,
	}, nil
}

func StartServer() *caddy.Instance {
	caddy.AppName = "test"
	caddy.AppVersion = "0.1"

	caddyfile, err := caddy.LoadCaddyfile("http")
	if err != nil {
		panic(err)
	}

	inst, err := caddy.Start(caddyfile)
	if err != nil {
		panic(err)
	}
	return inst
}
