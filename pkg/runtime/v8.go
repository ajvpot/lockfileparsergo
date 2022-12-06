package runtime

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"

	"go.kuoruan.net/v8go-polyfills/console"
	"go.kuoruan.net/v8go-polyfills/timers"
	"rogchap.com/v8go"

	"github.com/ajvpot/lockfileparsergo/js"
	"github.com/ajvpot/lockfileparsergo/pkg/reader"
)

// LoadFile loads a file with contents data at path name.
func LoadFile(ctx *v8go.Context, name, data string) error {
	fn, err := json.Marshal(name)
	if err != nil {
		return err
	}
	dataJ, err := json.Marshal(data)
	if err != nil {
		return err
	}

	_, err = ctx.RunScript(fmt.Sprintf("module.exports.fs.writeFileSync(%s, %s);", string(fn), string(dataJ)), `loader.js`)
	if err != nil {
		return err
	}
	return nil
}

// LoadJS loads the built js from esbuild into the v8 context.
func LoadJS(ctx *v8go.Context) error {
	_, err := ctx.RunScript(js.ShimsJS, "shims.js") // executes a script on the global context
	if err != nil {
		return err
	}

	_, err = ctx.RunScript(js.BuiltJS, "built.js") // executes a script on the global context
	if err != nil {
		if jserr, ok := err.(*v8go.JSError); ok {
			je, _ := json.Marshal(jserr)
			return errors.New(string(je))
		}
		return err
	}
	return err
}

// NewV8Context initializes a new v8 isolate and injects some polyfills.
func NewV8Context() (*v8go.Context, error) {
	iso := v8go.NewIsolate()
	global := v8go.NewObjectTemplate(iso)
	if err := timers.InjectTo(iso, global); err != nil {
		panic(err)
	}
	ctx := v8go.NewContext(iso, global)
	if err := console.InjectTo(ctx); err != nil {
		panic(err)
	}
	err := LoadJS(ctx)
	if err != nil {
		return nil, err
	}
	return ctx, nil
}

// LoadFileFromNamedReader loads a file with a name into the memory filesystem of V8.
func LoadFileFromNamedReader(ctx *v8go.Context, rdr reader.NamedReadCloser) error {
	manifestData, err := io.ReadAll(rdr)
	if err != nil {
		return err
	}
	err = LoadFile(ctx, rdr.Name(), string(manifestData))
	if err != nil {
		return err
	}
	return nil
}
