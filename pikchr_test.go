package pikchr

import (
	_ "embed"
	"fmt"
	"io"
	"testing"
)

//go:embed test.pikchr
var testPikchr string

//go:embed test.svg
var testSvg string

func TestPikchr(t *testing.T) {

	res, err := RenderString(testPikchr)
	if err != nil {
		t.Fatal(err)
	}

	if res != testSvg {
		t.Fatal("wrong output")
	}
}

func BenchmarkPikchr(b *testing.B) {
	b.RunParallel(benchStep)
}

func benchStep(pb *testing.PB) {
	for pb.Next() {
		result, err := RenderString(testPikchr)
		if err != nil {
			panic(err)
		}
		fmt.Fprintln(io.Discard, result)
	}
}
