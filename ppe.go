package ppe

import (
	"github.com/k0kubun/pp"
	"os"
)

func E(a ...interface{}) (n int, err error) {
	return pp.Fprintln(os.Stderr, a...)
}
