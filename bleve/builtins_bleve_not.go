//go:build !b_bleve
// +build !b_bleve

package bleve

import (
	"rye/env"
)

var Builtins_bleve = map[string]*env.Builtin{}
