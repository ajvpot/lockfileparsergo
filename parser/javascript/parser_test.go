package javascript

import (
	"encoding/json"
	"reflect"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/ajvpot/lockfileparsergo/parser"
	"github.com/ajvpot/lockfileparsergo/pkg/reader"
	"github.com/ajvpot/lockfileparsergo/test"
)

func TestParserSnapshot(t *testing.T) {
	p := NewParser()
	tree, err := p.BuildDepTree(reader.NewNamedReadCloser("package.json", strings.NewReader(test.Package)), reader.NewNamedReadCloser("package-lock.json", strings.NewReader(test.PackageLock)))
	assert.NoError(t, err)

	var expected parser.PkgTree
	err = json.Unmarshal([]byte(test.Expected), &expected)
	assert.NoError(t, err)

	assert.True(t, reflect.DeepEqual(&expected, tree), "trees are deeply equal")
}
