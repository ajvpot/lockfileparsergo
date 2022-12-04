# lockfileparsergo

[![PkgGoDev](https://pkg.go.dev/badge/github.com/ajvpot/lockfileparsergo)](https://pkg.go.dev/github.com/ajvpot/lockfileparsergo)

This repository contains a collection of JS libraries for parsing package manager lockfiles and glue code to make them
interoperable with golang. It uses [v8go](https://github.com/rogchap/v8go) to execute the scripts and returns native types.

Parsers are available for:
* JavaScript ([snyk-nodejs-lockfile-parser](https://github.com/snyk/nodejs-lockfile-parser))