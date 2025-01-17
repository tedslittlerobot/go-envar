package envar

import (
	"testing"
)

func TestPlainValueResolverInterfaceMatch(t *testing.T) {
	var _ ResolverInterface = PlainValueResolver{}
}

func TestPlainValueResolver(t *testing.T) {
	d := PlainValueResolver{}
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
