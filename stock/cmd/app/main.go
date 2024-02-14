package main

import (
	"flag"
	httpapp "go-junior-test/stock/internal/app/http"
	"log"
)

var opts = struct {
	httpAddr string
}{}

func init() {
	const (
		defaultHttpAddr = ":8083"
	)

	flag.StringVar(&opts.httpAddr, "http-addr", defaultHttpAddr, "http server address, default: "+defaultHttpAddr)
	flag.Parse()
}

func main() {
	errCh := make(chan error)
	{
		app := httpapp.NewApp(httpapp.Config{
			Address: opts.httpAddr,
		})
		go func() {
			errCh <- app.Run()
		}()
	}
	if err := <-errCh; err != nil {
		log.Fatal(err)
	}
}
