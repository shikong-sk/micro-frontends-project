package Proxy

import (
	"fmt"
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
)

func portainerProxy()  {
	http.HandleFunc("/",ServeHttp)
	_ = http.ListenAndServe(":9000", nil)
}

func ServeHttp(w http.ResponseWriter, r *http.Request)  {
	remote, err := url.Parse("http://192.168.3.2:9000/#!/")
	if err != nil {
		panic(err)
	}
	proxy := httputil.NewSingleHostReverseProxy(remote)
	r.URL.Scheme = remote.Scheme

	req, _ := httputil.DumpRequest(r, true)
	fmt.Println(string(req))

	proxy.Transport = &TransportInterceptor{}
	proxy.ServeHTTP(w, r)
}

type TransportInterceptor struct {
	// 捕获传输
	// CapturedTransport http.RoundTripper
}

func (t *TransportInterceptor) RoundTrip(request *http.Request) (*http.Response, error) {
	response, err := http.DefaultTransport.RoundTrip(request)
	// or, if you captured the transport
	// response, err := t.CapturedTransport.RoundTrip(request)

	// The httputil package provides a DumpResponse() func that will copy the
	// contents of the body into a []byte and return it. It also wraps it in an
	// ioutil.NopCloser and sets up the response to be passed on to the client.
	body, err := httputil.DumpResponse(response, true)
	if err != nil {
		// copying the response body did not work
		return nil, err
	}
	response.Header.Add("Access-Control-Allow-Origin", "*")
	// You may want to check the Content-Type header to decide how to deal with
	// the body. In this case, we're assuming it's text.
	log.Print(string(body))

	return response, err
}