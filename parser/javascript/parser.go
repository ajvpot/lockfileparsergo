package javascript

import (
	"mkm.pub/syncpool"
	"rogchap.com/v8go"

	"github.com/ajvpot/lockfileparsergo/parser"
	"github.com/ajvpot/lockfileparsergo/pkg/reader"
)

// Parser parses JavaScript project files.
type Parser interface {
	BuildDepTree(manifest, lockfile reader.NamedReadCloser) (*parser.PkgTree, error)
}

type jsDepTreeParser struct {
	v8Pool syncpool.Pool[*v8go.Context]
}

// NewParser creates a new javascript lockfile parser
func NewParser() Parser {
	return &jsDepTreeParser{v8Pool: syncpool.New(func() *v8go.Context {
		ctx, err := newV8Context()
		if err != nil {
			return nil
		}
		return ctx
	})}
}

// BuildDepTree returns a dependency tree given streams for a JavaScript project manifest and lockfile.
func (j *jsDepTreeParser) BuildDepTree(manifest, lockfile reader.NamedReadCloser) (*parser.PkgTree, error) {
	ctx := j.v8Pool.Get()
	defer j.v8Pool.Put(ctx)

	err := loadFileFromNamedReader(ctx, manifest)
	if err != nil {
		return nil, err
	}
	err = loadFileFromNamedReader(ctx, lockfile)
	if err != nil {
		return nil, err
	}

	return buildDepTree(ctx, manifest.Name(), lockfile.Name())
}
