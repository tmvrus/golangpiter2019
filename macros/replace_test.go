package macros

import (
	"bytes"
	"fmt"
	"github.com/valyala/fasttemplate"
	"io"
	"strings"
	"testing"
	"text/template"
)

func BenchmarkTextTemplate(b *testing.B) {
	tpl := strings.ReplaceAll(tplLayout, "{", "{{.")
	tpl = strings.ReplaceAll(tpl, "}", "}}")

	b.ResetTimer()

	b.RunParallel(func(pb *testing.PB) {
		var buf bytes.Buffer
		for pb.Next() {
			t, err := template.New("").Parse(tpl)
			if err != nil {
				b.Fatalf("got parse error: %q", err.Error())
			}

			if err := t.Execute(&buf, macrosData); err != nil {
				b.Fatalf("got execute error: %q", err.Error())
			}

			if buf.String() != result {
				b.Fatal("unexpected result")
			}

			buf.Reset()
		}
	})
}

func BenchmarkStringsReplacer(b *testing.B) {
	data := make([]string, 0, len(macrosData)*2)
	for k, v := range macrosData {
		k := fmt.Sprintf("{%s}", k)
		data = append(data, k, v)
	}

	b.ResetTimer()

	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			replacer := strings.NewReplacer(data...)

			if s := replacer.Replace(tplLayout); s != result {
				b.Errorf("unexpected result: %q", s)
			}
		}
	})

}

func BenchmarkStringsReplaceAll(b *testing.B) {
	data := make(map[string]string, len(macrosData))

	for k, v := range macrosData {
		data[fmt.Sprintf("{%s}", k)] = v
	}

	b.ResetTimer()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			tpl := tplLayout
			for k, v := range data {
				tpl = strings.ReplaceAll(tpl, k, v)
			}

			if result != tpl {
				b.Errorf("unexpected result: %q", tpl)
			}
		}
	})
}

func BenchmarkFasttemplate(b *testing.B) {
	data := make(map[string][]byte, len(macrosData))
	for k, v := range macrosData {
		data[k] = []byte(v)
	}
	b.ResetTimer()

	b.RunParallel(func(pb *testing.PB) {
		var buf bytes.Buffer

		for pb.Next() {
			tpl, err := fasttemplate.NewTemplate(tplLayout, "{", "}")
			if err != nil {
				b.Fatalf("got tplLayout error: %s", err.Error())
			}

			_, err = tpl.ExecuteFunc(&buf, func(w io.Writer, tag string) (int, error) {
				return w.Write(data[tag])
			})

			if err != nil {
				b.Fatalf("got executing error: %s", err.Error())
			}

			if s := buf.String(); s != result {
				b.Fatalf("unexpected result: %s", s)
			}

			buf.Reset()
		}
	})
}

var macrosData = map[string]string{
	"CLICK_URL":     "https://click.url.com/",
	"PIXEL_URL":     "https://pixel.url.com/",
	"TEST_ID":       "some test id",
	"CONVERSION_ID": "random id",
	"SUB_1":         "sub 1",
	"SUB_2":         "sub 2",
	"SUB_3":         "sub 3",
	"SUB_4":         "sub 4",
	"SUB_5":         "sub 5",
	"SUB_6":         "sub 6",
	"SUB_7":         "sub 7",
	"SUB_8":         "sub 8",
	"SUB_9":         "sub 9",
	"SUB_10":        "sub 10",
}

var tplLayout = `<html><a href="{CLICK_URL}""><button>Click!</button></a><img src="{PIXEL_URL}"><h1>{TEST_ID}</h1><h2>{CONVERSION_ID}</h2><ul><li>{SUB_1}</li><li>{SUB_2}</li><li>{SUB_3}</li><li>{SUB_4}</li><li>{SUB_5}</li><li>{SUB_6}</li><li>{SUB_7}</li><li>{SUB_8}</li><li>{SUB_9}</li><li>{SUB_10}</li><ul></html>`

var result = `<html><a href="https://click.url.com/""><button>Click!</button></a><img src="https://pixel.url.com/"><h1>some test id</h1><h2>random id</h2><ul><li>sub 1</li><li>sub 2</li><li>sub 3</li><li>sub 4</li><li>sub 5</li><li>sub 6</li><li>sub 7</li><li>sub 8</li><li>sub 9</li><li>sub 10</li><ul></html>`
