package envar

import (
	"testing"
)

func TestNeverResolverInterfaceMatch(t *testing.T) {
	var _ ResolverInterface = NeverResolver{}
}

func TestNeverResolver(t *testing.T) {
	d := NeverResolver{}
	token := SourceToken{
		Driver:     "Foo",
		Key:        "Bar",
		IsResolved: false,
		Value:      "",
	}

	d.Resolve(&token)

	if !token.IsResolved || token.Value != "" {
		t.Error("Token somehow got resolved")
	}
}
