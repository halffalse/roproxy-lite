package main

import (
	"log"
	"time"
	"os" // remove if not using os.Getenv("PORT") for port.
	"github.com/valyala/fasthttp"
)

// options
var timeout = 5 // timeout in seconds
var retries = 5 // number of request attempts. minimum: 1
var port = os.Getenv("PORT") // port to use. if not using heroku, replace with port in string format (e.g. "8080")

var client *fasthttp.Client

func main() {
	h := requestHandler
	
	client = &fasthttp.Client{
		ReadTimeout: time.Duration(timeout) * time.Second,
		MaxIdleConnDuration: 60 * time.Second,
	}

	if err := fasthttp.ListenAndServe(":" + port, h); err != nil {
		log.Fatalf("Error in ListenAndServe: %s", err)
	}
}

func requestHandler(ctx *fasthttp.RequestCtx) {
	response := makeRequest(ctx, 1)

	defer fasthttp.ReleaseResponse(response)

	body := response.Body()
	ctx.SetBody(body)
	ctx.SetStatusCode(response.StatusCode())
	response.Header.VisitAll(func (key, value []byte) {
		ctx.Response.Header.Set(string(key), string(value))
	})
}

func makeRequest(ctx *fasthttp.RequestCtx, attempt int) *fasthttp.Response {
	if attempt > retries {
		resp := fasthttp.AcquireResponse()
		resp.SetBody([]byte("Proxy failed to connect. Please try again."))
		resp.SetStatusCode(500)

		return resp
	}

	req := fasthttp.AcquireRequest()
	defer fasthttp.ReleaseRequest(req)
	req.Header.SetMethod(string(ctx.Method()))
	req.SetRequestURI(string(ctx.Request.Header.RequestURI())[1:])
	req.SetBody(ctx.Request.Body())
	ctx.Request.Header.VisitAll(func (key, value []byte) {
		req.Header.Set(string(key), string(value))
	})
	req.Header.Set("User-Agent", "RoProxy")
	req.Header.Del("Roblox-Id")
	resp := fasthttp.AcquireResponse()

	err := client.Do(req, resp)

    if err != nil {
		fasthttp.ReleaseResponse(resp)
        return makeRequest(ctx, attempt + 1)
    } else {
		return resp
	}
}