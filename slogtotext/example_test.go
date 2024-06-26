package slogtotext_test

import (
	"os"
	"strings"

	"github.com/michurin/human-readable-json-logging/slogtotext"
)

func Example() {
	f := slogtotext.MustFormatter(os.Stdout, `x={{.x}}{{if .ALL | rm "x"}} UNKNOWN:{{range .ALL | rm "x" "p" "q"}} {{.K}}={{.V}}{{end}}{{end}}`+"\n")
	g := slogtotext.MustFormatter(os.Stdout, `INVALID LINE: {{ .TEXT | printf "%q" }}`+"\n")
	buf := strings.NewReader(`{"x": 100}
{"x": 1, "y": { "a": 2, "b": 3 }, "p": 9, "q": 9}
here`)
	err := slogtotext.Read(buf, f, g, 1024)
	if err != nil {
		panic(err)
	}
	// Output:
	// x=100
	// x=1 UNKNOWN: y.a=2 y.b=3
	// INVALID LINE: "here"
}
