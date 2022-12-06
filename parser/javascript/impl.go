package javascript

import (
	"encoding/json"
	"fmt"

	"rogchap.com/v8go"

	"github.com/ajvpot/lockfileparsergo/parser"
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
