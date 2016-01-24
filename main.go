package main

import (
	"flag"
	"fmt"
	"log"
	"net"
	"os"

	"net/http"
	"net/http/httputil"
	"net/url"
)

var (
	Url      *url.URL
	ListenOn string
	FileName string
)

func init() {
	proxyUrl := flag.String("proxy", "http://localhost:8000/", "URL to proxy requests to")
	listeningPort := flag.Int("listening-port", 80, "Port to listen on")
	allowExternal := flag.Bool("allow-external-connections", false, "Allow other computers to connect to your HTTP server")
	flag.Parse()

	var err error
	Url, err = url.Parse(*proxyUrl)
	if err != nil {
		fmt.Fprint(os.Stderr, err)
		flag.Usage()
		os.Exit(2)
	}

	if Url.Scheme == "unix" {
		FileName = Url.Host
		// Handle URLs written like unix:whatever instead of unix://whatever
		if FileName == "" {
			FileName = Url.Opaque
		}
		Url = &url.URL{
			Scheme: "http",
			Host:   "127.0.0.1",
		}
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

func SocketDialer(network, addr string) (conn net.Conn, err error) {
	return net.Dial("unix", FileName)
}

func main() {
	proxied := Url.String()
	rp := httputil.NewSingleHostReverseProxy(Url)

	// Handle unix socket connections
	if FileName != "" {
		proxied = FileName
		rp.Transport = &http.Transport{
			Dial: SocketDialer,
		}
	}

	fmt.Printf("Started simple reverse proxy from %s to %s...\n\n", ListenOn, proxied)

	if err := http.ListenAndServe(ListenOn, LoggingMiddleware(rp)); err != nil {
		fmt.Fprintf(os.Stderr, "Error: %s\n\n", err)
		os.Exit(1)
	}
}
