package envarResolvers

import "github.com/tedslittlerobot/go-envar/support/data"

// RawValueResolver resolves a value to its key directly. This is used for allowing a known string to be used,
// typically as a fallback in case any other resolvers cannot resolve.
type RawValueResolver struct{}

func (resolver RawValueResolver) Resolve(token *envarData.SourceToken) {
	token.Resolve(token.Key)
}

func (resolver RawValueResolver) PreLoad(tokens []*envarData.SourceToken) {}
