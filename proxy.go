package main

import (
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
)

type Transport struct{}

// function to modifying Transport
func (t *Transport) RoundTrip(req *http.Request) (*http.Response, error) {
	req.Host = req.URL.Host
	return http.DefaultTransport.RoundTrip(req)
}

// handler for proxy
func proxyHandler(w http.ResponseWriter, r *http.Request) {

	// url example httpbin.org
	url, err := url.Parse("http://httpbin.org/headers")

	if err != nil {
		panic(err)
	}

	proxy := httputil.ReverseProxy{
		// control request to endpoint targer.
		Director: func(r *http.Request) {
			r.URL.Host = url.Host
			r.URL.Scheme = url.Scheme
			r.URL.Path = url.Path // golang pass pattern to the proxy as path so we modifyied it befor send to the endpoint
		},
	}
	proxy.Transport = &Transport{}
	proxy.ServeHTTP(w, r)

}

func main() {
	http.HandleFunc("/proxy", proxyHandler)
	// start server
	log.Fatal(http.ListenAndServe(":3000", nil))
}
