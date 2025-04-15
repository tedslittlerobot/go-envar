package envarResolvers

import "github.com/tedslittlerobot/go-envar/support/data"

// NeverResolver always fails resolution. Practically speaking, used for testing purposes only.
type NeverResolver struct{}

func (resolver NeverResolver) Resolve(token *envarData.SourceToken) {
	token.ResolveBlank()
}

func (resolver NeverResolver) PreLoad(tokens []*envarData.SourceToken) {}
