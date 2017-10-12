package main

import (
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
)

type Transport struct{}

func (t *Transport) RoundTrip(req *http.Request) (*http.Response, error) {
	req.Host = req.URL.Host
	return http.DefaultTransport.RoundTrip(req)
}

func proxyHandler(w http.ResponseWriter, r *http.Request) {

	url, err := url.Parse("http://httpbin.org/headers")

	if err != nil {
		panic(err)
	}

	proxy := httputil.ReverseProxy{
		Director: func(r *http.Request) {
			r.URL.Host = url.Host
			r.URL.Scheme = url.Scheme
			r.URL.Path = url.Path
		},
	}
	proxy.Transport = &Transport{}
	proxy.ServeHTTP(w, r)

}

func main() {

	http.HandleFunc("/proxy", proxyHandler)
	// note the handler passed to ListenAndServe.
	log.Fatal(http.ListenAndServe(":3000", nil))
}
