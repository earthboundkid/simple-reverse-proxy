package main

import (
	"flag"
	"fmt"
	"os"

	"net/http"
	"net/http/httputil"
	"net/url"
)

var (
	scheme string
	host   string
	port   int
)

func init() {
	flag.IntVar(&port, "port", 8000, "Port to reverse proxy")
	flag.StringVar(&host, "host", "localhost", "Host to reverse proxy")
	flag.StringVar(&scheme, "scheme", "http", "Scheme to reverse proxy")
}

func main() {
	flag.Parse()

	url := &url.URL{
		Scheme: scheme,
		Host:   fmt.Sprintf("%s:%d", host, port),
	}

	LoggingMiddleware := func(s http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			fmt.Printf("Proxying %s to %s\n", r.URL.Path, url)
			s.ServeHTTP(w, r)
		})
	}

	proxy := LoggingMiddleware(httputil.NewSingleHostReverseProxy(url))

	fmt.Printf("Starting reverse proxy to %s...\n\n", url)

	if err := http.ListenAndServe(":80", proxy); err != nil {
		fmt.Fprintf(os.Stderr, "Error: %s\n\n", err)
		os.Exit(1)
	}

}
