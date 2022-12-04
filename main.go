package main

import (
	"fmt"
	"os"
	"path"

	"github.com/ajvpot/lockfileparsergo/parser/javascript"
	"github.com/ajvpot/lockfileparsergo/pkg/reader"
)

func main() {
	parser := javascript.NewParser()

	manifest, err := os.Open(os.Args[1])
	if err != nil {
		panic(err)
	}
	lockfile, err := os.Open(os.Args[2])
	if err != nil {
		panic(err)
	}

	tree, err := parser.BuildDepTree(reader.NewNamedReadCloser(path.Base(os.Args[1]), manifest), reader.NewNamedReadCloser(path.Base(os.Args[2]), lockfile))
	if err != nil {
		panic(err)
	}
	fmt.Printf("result: %v", tree)
}
