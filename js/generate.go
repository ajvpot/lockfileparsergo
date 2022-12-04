package js

import (
	_ "embed"
)

//go:generate yarn install
//go:generate node build.mjs

//go:embed dist/built.js
var BuiltJS string

//go:embed shims.js
var ShimsJS string
