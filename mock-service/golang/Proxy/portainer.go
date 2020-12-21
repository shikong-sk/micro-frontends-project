package Proxy

import (
	"io/ioutil"
	"net/http"
	"net/http/httputil"
	"net/url"
)

func portainerProxy()  {
	http.HandleFunc("/",ServeHttp)
	_ = http.ListenAndServe(":9000", nil)
}

func ServeHttp(w http.ResponseWriter, r *http.Request)  {
	remote, err := url.Parse("http://192.168.3.2:9000")
	if err != nil {
		panic(err)
	}
	proxy := httputil.NewSingleHostReverseProxy(remote)

	// 设置新的header
	for k, vs := range r.Header {
		for _, v := range vs {
			w.Header().Add(k,v)
		}
	}
	w.Header().Add("Access-Control-Allow-Origin", "*")
	w.WriteHeader(200)

	// 获取新的body
	bodyRes,err:=ioutil.ReadAll(r.Body)
	_, _ = w.Write(bodyRes)

	proxy.ServeHTTP(w, r)
}