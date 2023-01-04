# Pikchr for Go

This package wraps [Pikchr](https://pikchr.org) as a Go package. It doesn't rely
on cgo or an external binary, instead using the [wazero](https://wazero.io/)
WebAssembly runtime to embed Pikchr.

Wazero is configured to sandbox Pikchr such that Pikchr only has access to the
input and output data.

## Performance

For a small Pikchr figure, go-pikchr takes about 7.5ms:

```shell
$ go test -test.bench .
goos: linux
goarch: amd64
pkg: github.com/rigtorp/go-pikchr
cpu: Intel(R) Core(TM) i7-8550U CPU @ 1.80GHz
BenchmarkPikchr-8            142           7520289 ns/op
PASS
ok      github.com/rigtorp/go-pikchr    1.471s
```

## Building Pikchr for WASM

Download [wasi-sdk](https://github.com/WebAssembly/wasi-sdk) and extract it
somewhere.

Next build Pikchr:

```shell
export WASI_SDK_PATH=/path/to/wasi/sdk
$WASI_SDK_PATH/bin/clang pikchr.c -o pikchr.wasm -DPIKCHR_SHELL -Os -DNDEBUG -s
```

## Acknowledgements

I got the idea for this approach of embedding an external program in Go using
WebAssembly from the blog post [The carcinization of Go
programs](https://xeiaso.net/blog/carcinization-golang).
