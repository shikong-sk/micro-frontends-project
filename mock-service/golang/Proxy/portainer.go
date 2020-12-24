package Proxy

import (
	"fmt"
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
