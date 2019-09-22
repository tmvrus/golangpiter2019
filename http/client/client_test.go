package client

// see https://github.com/valyala/fasthttp/blob/master/client_timing_test.go
import (
	"bytes"
	"fmt"
	"github.com/valyala/fasthttp"
	"io/ioutil"
	"net"
	"net/http"
	"runtime"
	"sync"
	"sync/atomic"
	"testing"
	"time"
)

var connPool sync.Pool

type Conn struct {
	net.Conn
	s  []byte
	n  int
	ch chan struct{}
}

func (c *Conn) LocalAddr() net.Addr {
	return &net.TCPAddr{
		IP:   []byte{1, 2, 3, 4},
		Port: 1234,
	}
}

func (c *Conn) RemoteAddr() net.Addr {
	return &net.TCPAddr{
		IP:   []byte{1, 2, 3, 4},
		Port: 1234,
	}
}

func (c *Conn) Close() error {
	fmt.Println("close")
	c.n = 0
	connPool.Put(c)
	return nil
}

func (c *Conn) Write(b []byte) (int, error) {
	c.ch <- struct{}{}
	return len(b), nil
}

func (c *Conn) Read(b []byte) (int, error) {
	if c.n == 0 {
		<-c.ch
	}

	n := 0
	for len(b) > 0 {
		if c.n == len(c.s) {
			c.n = 0
			return n, nil
		}

		n = copy(b, c.s[c.n:])
		c.n += n
		b = b[n:]
	}

	return n, nil

}

//acquireFakeServerConn
func acquireConn(s []byte) *Conn {
	c := connPool.Get()
	if c == nil {
		c := &Conn{
			s:  s,
			ch: make(chan struct{}, 100000),
		}
		return c
	}

	return c.(*Conn)
}

func BenchmarkFastHTTPClient(b *testing.B) {
	fullResponse := []byte(fmt.Sprintf("HTTP/1.1 200 OK\r\nContent-Type: application-json\r\nContent-Length: %d\r\n\r\n%s", len(responseBody), responseBody))
	client := &fasthttp.Client{
		Dial: func(addr string) (net.Conn, error) {
			return acquireConn(fullResponse), nil
		},
		MaxConnsPerHost: runtime.GOMAXPROCS(-1),
	}

	siteNum := uint32(0)
	b.RunParallel(func(pb *testing.PB) {
		var req fasthttp.Request
		var resp fasthttp.Response
		req.SetBody(requestBody)
		req.Header.SetMethod(fasthttp.MethodPost)
		req.SetRequestURI(fmt.Sprintf("http://site%d.com/path", atomic.AddUint32(&siteNum, 1)))
		req.Header.Add("Content-Type", "application/json")

		for pb.Next() {
			if err := client.DoTimeout(&req, &resp, time.Second); err != nil {
				b.Fatalf("got fasthttp error: %s", err.Error())
			}

			if status := resp.Header.StatusCode(); status != fasthttp.StatusOK {
				b.Fatalf("unexpected status code: %d", status)
			}

			if !bytes.Equal(resp.Body(), responseBody) {
				b.Fatalf("unepected body diff")
			}
		}
	})
}

func BenchmarkNetHTTPClient(b *testing.B) {
	fullResponse := []byte(fmt.Sprintf("HTTP/1.1 200 OK\r\nContent-Type: application-json\r\nContent-Length: %d\r\n\r\n%s", len(responseBody), responseBody))
	client := &http.Client{
		Transport: &http.Transport{
			Dial: func(_, _ string) (net.Conn, error) {
				return acquireConn(fullResponse), nil
			},
			MaxConnsPerHost: runtime.GOMAXPROCS(-1),
		},
	}

	siteNum := uint32(0)
	b.RunParallel(func(pb *testing.PB) {
		req, err := http.NewRequest(http.MethodPost, fmt.Sprintf("http://site%d.com/path", atomic.AddUint32(&siteNum, 1)), nil)
		if err != nil {
			b.Fatalf("failed to create http request: %s", err.Error())
		}

		for pb.Next() {
			resp, err := client.Do(req)

			if err != nil {
				b.Fatalf("unexpected client error: %s", err.Error())
			}

			if resp.StatusCode != fasthttp.StatusOK {
				b.Fatalf("unexpected status code: %d", resp.StatusCode)
			}

			readBody, err := ioutil.ReadAll(resp.Body)

			if err != nil {
				b.Fatalf("failed to read body: %s", err.Error())
			}

			if err := resp.Body.Close(); err != nil {
				b.Fatalf("failed to close body: %s", err.Error())
			}

			if !bytes.Equal(readBody, responseBody) {
				b.Fatalf("unepected body diff")
			}
		}
	})
}
