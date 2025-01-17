package envar

import "os"

type EnvResolver struct{}

func (resolver EnvResolver) Resolve(token *SourceToken) {
	token.Resolve(os.Getenv(token.Key))
}

func (resolver EnvResolver) PreLoad(tokens []*SourceToken) {}
