package main

import (
	"flag"
	httpapp "go-junior-test/stock/internal/app/http"
	"log"
	"os"
)

var opts = struct {
	httpAddr    string
	databaseURL string
}{}

func init() {
	const (
		defaultHttpAddr = ":8083"
	)

	flag.StringVar(&opts.httpAddr, "http-addr", defaultHttpAddr, "http server address, default: "+defaultHttpAddr)
	flag.Parse()

	if opts.databaseURL == "" {
		opts.databaseURL = os.Getenv("DATABASE_URL")
	}
}

func main() {
	errCh := make(chan error)
	{
		app := httpapp.NewApp(httpapp.Config{
			Address:     opts.httpAddr,
			DatabaseURL: opts.databaseURL,
		})
		go func() {
			errCh <- app.Run()
		}()
	}
	if err := <-errCh; err != nil {
		log.Fatal(err)
	}
}
