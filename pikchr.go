// Package pikchr wraps a WASM build of pikchr as go package.
package pikchr

import (
	"bytes"
	"context"
	_ "embed"
	"io"

	"github.com/tetratelabs/wazero"
	"github.com/tetratelabs/wazero/imports/wasi_snapshot_preview1"
)

//go:embed pikchr.wasm
var pikchrWasm []byte

// Render converts input Pikchr figure into output SVG figure.
func Render(in io.Reader, out io.Writer) error {
	ctx := context.TODO()

	r := wazero.NewRuntimeWithConfig(ctx, wazero.NewRuntimeConfigInterpreter())
	defer r.Close(ctx)

	wasi_snapshot_preview1.MustInstantiate(ctx, r)

	code, err := r.CompileModule(ctx, pikchrWasm)
	if err != nil {
		return err
	}

	// Only stdin and stdout is exposed to the WASM module
	config := wazero.NewModuleConfig().WithStdout(out).WithStdin(in).WithArgs("pikchr", "--svg-only", "-")

	_, err = r.InstantiateModule(ctx, code, config)
	if err != nil {
		return err
	}

	return nil
}

// RenderString converts input Pikchr figure into output SVG figure.
func RenderString(text string) (string, error) {
	out := &bytes.Buffer{}
	in := bytes.NewBufferString(text)

	if err := Render(in, out); err != nil {
		return "", err
	}

	return out.String(), nil
}
