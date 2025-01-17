package envar

import "os"

type EnvironmentVariableResolver struct{}

func (resolver EnvironmentVariableResolver) Resolve(token *SourceToken) {
	token.Resolve(os.Getenv(token.Key))
}

func (resolver EnvironmentVariableResolver) PreLoad(tokens []*SourceToken) {}
