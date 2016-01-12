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
	Url      url.URL
	ListenOn string
)

func init() {
	scheme := flag.String("scheme", "http", "Scheme to reverse proxy")
	host := flag.String("host", "localhost", "Host to reverse proxy")
	port := flag.Int("port", 8000, "Port to reverse proxy")
	listeningPort := flag.Int("listening-port", 80, "Port to listen on")
	allowExternal := flag.Bool("allow-external-connections", false, "Allow other computers to connect to your HTTP server")
	flag.Parse()

	Url = url.URL{
		Scheme: *scheme,
		Host:   fmt.Sprintf("%s:%d", *host, *port),
	}

	if *allowExternal {
		ListenOn = fmt.Sprintf(":%d", *listeningPort)
	} else {
		ListenOn = fmt.Sprintf("127.0.0.1:%d", *listeningPort)
	}

}

func LoggingMiddleware(s http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Printf("Reverse proxying %s ...\n", r.URL.Path)
		s.ServeHTTP(w, r)
	})
}

func main() {
	fmt.Printf("Started simple reverse proxy from %s to %s...\n\n", ListenOn, &Url)

	proxy := LoggingMiddleware(httputil.NewSingleHostReverseProxy(&Url))

	if err := http.ListenAndServe(ListenOn, proxy); err != nil {
		fmt.Fprintf(os.Stderr, "Error: %s\n\n", err)
		os.Exit(1)
	}
}
