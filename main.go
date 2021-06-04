package main

import (
	"flag"
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
)

func main() {
	addr := flag.String("addr", ":8080", "listen addr")
	target := flag.String("target", "https://github.com/", "target backend")
	flag.Parse()

	parsedTarget, err := url.Parse(*target)
	if err != nil {
		panic(err)
	}

	proxy := &httputil.ReverseProxy{
		Director: func(r *http.Request) {
			r.Host = ""
			r.URL.Scheme = parsedTarget.Scheme
			r.URL.Host = parsedTarget.Host
		},
	}

	log.Printf("listening on %s; proxying to %s\n", *addr, *target)
	err = http.ListenAndServe(*addr, proxy)
	if err != nil {
		panic(err)
	}
}
