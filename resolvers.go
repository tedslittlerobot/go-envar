package envar

import (
	"context"
	envarAws "github.com/tedslittlerobot/go-envar/support/aws"
	"os"
)

func (e *Envar) AddSsmResolver(ctx context.Context, name string, prefix string, recursive bool) {
	e.Resolvers.AddResolver(
		name,
		envarAws.MakeDefaultParameterStoreMapResolver(ctx, prefix, recursive),
	)
}

func (e *Envar) AddSsmResolverFromEnv(ctx context.Context, name string, envVar string, recursive bool) {
	prefix := os.Getenv(envVar)

	if prefix == "" {
		panic("environment variable '" + name + "' is not set")
	}

	e.AddSsmResolver(ctx, name, prefix, recursive)
}
