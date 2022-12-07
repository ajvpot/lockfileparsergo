package test

import (
	_ "embed"
)

//go:embed package.json
var Package string

//go:embed package-lock.json
var PackageLock string

//go:embed expected.json
var Expected string
