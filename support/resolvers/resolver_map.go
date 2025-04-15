package envarResolvers

import (
	"github.com/tedslittlerobot/go-envar/support/data"
)

// MapResolver resolves values from the environment (ie. using os.GetEnv())
type MapResolver struct {
	Contents map[string]string
}

func (resolver MapResolver) Resolve(token *envarData.SourceToken) {
	value, exists := resolver.Contents[token.Key]

	if exists {
		token.Resolve(value)
	} else {
		token.ResolveBlank()
	}
}

func (resolver MapResolver) PreLoad(tokens []*envarData.SourceToken) {}
