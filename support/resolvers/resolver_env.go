package envarResolvers

import (
	"github.com/tedslittlerobot/go-envar/support/data"
	"os"
)

// EnvironmentVariableResolver resolves values from the environment (ie. using os.GetEnv())
type EnvironmentVariableResolver struct{}

func (resolver EnvironmentVariableResolver) Resolve(token *envarData.SourceToken) {
	token.Resolve(os.Getenv(token.Key))
}

func (resolver EnvironmentVariableResolver) PreLoad(tokens []*envarData.SourceToken) {}
