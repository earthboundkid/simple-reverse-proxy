package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"net/http"
	"net/http/httputil"
	"net/url"
)

var (
	Url      *url.URL
	ListenOn string
)

func init() {
	var (
		scheme        string
		host          string
		port          int
		listeningPort int
		allowExternal bool
	)
	flag.IntVar(&port, "port", 8000, "Port to reverse proxy")
	flag.StringVar(&host, "host", "localhost", "Host to reverse proxy")
	flag.StringVar(&scheme, "scheme", "http", "Scheme to reverse proxy")
	flag.IntVar(&listeningPort, "listening-port", 80, "Port to listen on")
	flag.BoolVar(&allowExternal, "allow-external-connections", false, "Allow other computers to connect to your HTTP server")
	flag.Parse()

	Url = &url.URL{
		Scheme: scheme,
		Host:   fmt.Sprintf("%s:%d", host, port),
	}

	if allowExternal {
		ListenOn = fmt.Sprintf(":%d", listeningPort)
	} else {
		ListenOn = fmt.Sprintf("127.0.0.1:%d", listeningPort)
	}

}

func main() {

	LoggingMiddleware := func(s http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			log.Printf("Proxying %s to %s\n", r.URL.Path, Url)
			s.ServeHTTP(w, r)
		})
	}

	proxy := LoggingMiddleware(httputil.NewSingleHostReverseProxy(Url))


	fmt.Printf("Starting reverse proxy from %s to %s...\n\n", ListenOn, Url)

	if err := http.ListenAndServe(ListenOn, proxy); err != nil {
		fmt.Fprintf(os.Stderr, "Error: %s\n\n", err)
		os.Exit(1)
	}

}
