package main

import (
	"github.com/tellett/gocode/pkg/examples/hello"
	"go.uber.org/fx"
)

var Module = fx.Options(
	hello.Module,
)

func main() {
	fx.New(Module).Run()
}
