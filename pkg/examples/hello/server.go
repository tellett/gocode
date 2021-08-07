package hello

import (
	"github.com/tellett/gocode/pkg/examples/hello/handlers/greeter"
	"github.com/tellett/gocode/pkg/net/httpserver"

	"go.uber.org/fx"
)

var Module = fx.Options(
	httpserver.Module,
	greeter.Module,
)
