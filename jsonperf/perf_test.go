package jsonperf

import (
	"encoding/json"
	"github.com/bsm/openrtb"
	"sync"
	"testing"

	"github.com/json-iterator/go"
)

func BenchmarkEncodingJSONMarshal(b *testing.B) {
	wg := &sync.WaitGroup{}
	wg.Add(b.N)
	for i := 0; i < b.N; i++ {
		go func() {
			bid := bidresponse()
			for i := 0; i < 100; i++ {
				_, err := json.Marshal(bid)
				if err != nil {
					b.Error(err.Error())
				}
			}
			wg.Done()
		}()
	}
	wg.Wait()
}

func BenchmarkJsoniterMarshal(b *testing.B) {
	wg := &sync.WaitGroup{}
	wg.Add(b.N)
	for i := 0; i < b.N; i++ {
		go func() {
			bid := bidresponse()
			for i := 0; i < 100; i++ {
				_, err := jsoniter.Marshal(bid)
				if err != nil {
					b.Error(err.Error())
				}
			}
			wg.Done()
		}()
	}
	wg.Wait()
}

func BenchmarkEncodingJSONUnmarshal(b *testing.B) {
	wg := &sync.WaitGroup{}
	wg.Add(b.N)
	for i := 0; i < b.N; i++ {
		go func() {
			bid := &openrtb.BidResponse{}
			for i := 0; i < 100; i++ {
				err := json.Unmarshal(data, bid)
				if err != nil {
					b.Error(err.Error())
				}
			}
			wg.Done()
		}()
	}
	wg.Wait()
}

func BenchmarkJsoniterUnmarshal(b *testing.B) {
	wg := &sync.WaitGroup{}
	wg.Add(b.N)
	for i := 0; i < b.N; i++ {
		go func() {
			bid := &openrtb.BidResponse{}
			for i := 0; i < 100; i++ {
				err := jsoniter.Unmarshal(data, bid)
				if err != nil {
					b.Error(err.Error())
				}
			}
			wg.Done()
		}()
	}
	wg.Wait()
}
