package main

import (
	"fmt"
	"log"
	"net/http"

	"time"

	"github.com/elazarl/goproxy"
)

var hosts = [...]string{"google.com", "bbc.co.uk", "basketapi.local"}

const delay = 5

func main() {
	fmt.Println("Starting Proxy")
	proxy := goproxy.NewProxyHttpServer()

	proxy.Verbose = true

	for _, host := range hosts {
		proxy.OnRequest(goproxy.DstHostIs(host)).DoFunc(addLatency)
	}

	//proxy.OnRequest(goproxy.DstHostIs("")).DoFunc(addLatency)
	proxy.OnRequest(goproxy.DstHostIs("healthcheck.local")).DoFunc(healthCheck)
	log.Fatal(http.ListenAndServe(":8080", proxy))
}

func addLatency(req *http.Request, ctx *goproxy.ProxyCtx) (*http.Request, *http.Response) {
	log.Printf("Injecting %d seconds latency for %s", delay, req.Host)
	time.Sleep(time.Second * delay)
	return req, nil
}

func healthCheck(req *http.Request, ctx *goproxy.ProxyCtx) (*http.Request, *http.Response) {
	resp := goproxy.TextResponse(req, "Running")
	resp.StatusCode = 200
	return req, resp
}
