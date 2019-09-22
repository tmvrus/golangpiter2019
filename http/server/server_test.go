package server

// see https://github.com/valyala/fasthttp/blob/master/server_timing_test.go

import (
	"bytes"
	"fmt"
	"github.com/valyala/fasthttp"
	"io"
	"io/ioutil"
	"net"
	"net/http"
	"testing"
	"time"
)

const (
	clientsCount         = 100
	requestPerConnection = 1000
)

func BenchmarkFastHTTPServer(b *testing.B) {
	requestPipe := make(chan struct{}, b.N)
	server := &fasthttp.Server{
		Handler: func(ctx *fasthttp.RequestCtx) {
			body := ctx.Request.Body()
			if !bytes.Equal(body, responseBody) {
				b.Fatalf("Unexpected body %q. Expected %q", body, responseBody)
			}
			ctx.Success("application-json", body)

			select {
			case requestPipe <- struct{}{}:
			default:
				b.Fatalf("request count exceed %d", b.N)
			}
		},
		Concurrency: 16 * clientsCount,
	}

	doBenchmark(b, server, clientsCount, requestPerConnection, postRequest)
	verifyRequestsServed(b, requestPipe)
}

func BenchmarkNetHTTPServer(b *testing.B) {
	requestPipe := make(chan struct{}, b.N)
	server := &http.Server{
		Handler: http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
			_, err := ioutil.ReadAll(req.Body)
			if err != nil {
				b.Fatalf("failed to read body: %s", err.Error())
			}

			if err := req.Body.Close(); err != nil {
				b.Fatalf("failed to close body: %s", err.Error())
			}

			h := w.Header()
			h.Set("Content-Type", "application-json")

			if _, err := w.Write(responseBody); err != nil {
				b.Fatalf("failed to write response body: %s", err.Error())
			}

			select {
			case requestPipe <- struct{}{}:
			default:
				b.Fatalf("request count exceed %d", b.N)
			}

		}),
	}

	doBenchmark(b, server, clientsCount, requestPerConnection, postRequest)
	verifyRequestsServed(b, requestPipe)
}

type server interface {
	Serve(ln net.Listener) error
}

func doBenchmark(b *testing.B, s server, clientsCount, requestsPerConn int, requestBody []byte) {
	ln := newFakeListener(b.N, clientsCount, requestsPerConn, requestBody)
	done := make(chan struct{})
	go func() {
		if err := s.Serve(ln); err != nil && err != io.EOF {
			b.Fatalf("got Serve error: %s", err.Error())
		}

		done <- struct{}{}
	}()

	<-ln.done

	select {
	case <-done:
	case <-time.After(10 * time.Second):
		b.Fatalf("Server.Serve() didn't stop")
	}
}

func verifyRequestsServed(b *testing.B, ch <-chan struct{}) {
	requestsServed := 0
	for len(ch) > 0 {
		<-ch
		requestsServed++
	}

	requestsSent := b.N
	for requestsServed < requestsSent {
		select {
		case <-ch:
			requestsServed++
		case <-time.After(100 * time.Millisecond):
			b.Fatalf("Unexpected number of requests served %d. Expected %d", requestsServed, requestsSent)
		}
	}
}

var postRequest = []byte(fmt.Sprintf("POST /foobar?baz HTTP/1.1\r\nHost: google.com\r\nContent-Type: application-json\r\nContent-Length: %d\r\n"+
	"User-Agent: Opera Chrome MSIE Firefox and other/1.2.34\r\nReferer: http://google.com/aaaa/bbb/ccc\r\n"+
	"Cookie: foo=bar; baz=baraz; aa=aakslsdweriwereowriewroire\r\n\r\n%s",
	len(responseBody), responseBody))
