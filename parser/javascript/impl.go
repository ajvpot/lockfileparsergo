package javascript

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"

	"go.kuoruan.net/v8go-polyfills/console"
	"go.kuoruan.net/v8go-polyfills/timers"
	"rogchap.com/v8go"

	"github.com/ajvpot/lockfileparsergo/js"
	"github.com/ajvpot/lockfileparsergo/parser"
	"github.com/ajvpot/lockfileparsergo/pkg/reader"
)

func mustMarshal(v any) []byte {
	r, err := json.Marshal(v)
	if err != nil {
		panic(err)
	}
	return r
}

func buildDepTree(ctx *v8go.Context, manifestPath, lockfilePath string) (*parser.PkgTree, error) {
	val, err := ctx.RunScript(fmt.Sprintf(`module.exports.buildDepTreeFromFiles("./", %s, %s);`, mustMarshal(manifestPath), mustMarshal(lockfilePath)), `invoke.js`)
	if err != nil {
		return nil, err
	}

	valPromise, err := val.AsPromise()
	if err != nil {
		return nil, err
	}

	jb, err := valPromise.Result().MarshalJSON()
	if err != nil {
		return nil, err
	}

	var out parser.PkgTree
	err = json.Unmarshal(jb, &out)
	if err != nil {
		panic(err)
	}
	return &out, nil
}

func loadFile(ctx *v8go.Context, name, data string) error {
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

// loadJS loads the built js from esbuild into the v8 context.
func loadJS(ctx *v8go.Context) error {
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

// newV8Context initializes a new v8 isolate and injects some polyfills.
func newV8Context() (*v8go.Context, error) {
	iso := v8go.NewIsolate()
	global := v8go.NewObjectTemplate(iso)
	if err := timers.InjectTo(iso, global); err != nil {
		panic(err)
	}
	ctx := v8go.NewContext(iso, global)
	if err := console.InjectTo(ctx); err != nil {
		panic(err)
	}
	err := loadJS(ctx)
	if err != nil {
		return nil, err
	}
	return ctx, nil
}

// loadFileFromNamedReader loads a file with a name into the memory filesystem of V8.
func loadFileFromNamedReader(ctx *v8go.Context, rdr reader.NamedReadCloser) error {
	manifestData, err := io.ReadAll(rdr)
	if err != nil {
		return err
	}
	err = loadFile(ctx, rdr.Name(), string(manifestData))
	if err != nil {
		return err
	}
	return nil
}
