package envar

import (
	"testing"
)

func TestRawValueResolverInterfaceMatch(t *testing.T) {
	var _ ResolverInterface = RawValueResolver{}
}

func TestRawValueResolver(t *testing.T) {
	d := RawValueResolver{}
	token := SourceToken{
		Driver:     "Foo",
		Key:        "Bar",
		IsResolved: false,
		Value:      "",
	}

	d.Resolve(&token)

	if token.Value != "Bar" {
		t.Error("Token did not get resolved")
	}
}
