package envar

import (
	"testing"
)

func TestSourceTokenCollectionRegister(t *testing.T) {
	registry := SourceTokenRegistry{}

	response := registry.Register("Foo:Bar")

	if response.Driver != "Foo" {
		t.Errorf("response.ResolverInterface should be Foo, but is %s", response.Driver)
	}

	if response.Key != "Bar" {
		t.Errorf("response.Key should be Bar, but is %s", response.Key)
	}

	if registry.Tokens["Foo"]["Bar"] != response {
		t.Error("returned token pointer should match the registered one")
	}

	response.Value = "Monkeys"

	if registry.Tokens["Foo"]["Bar"].Value != "Monkeys" {
		t.Error("value should be updated")
	}
}

func TestSourceTokenCollectionRegisterChain(t *testing.T) {
	registry := SourceTokenRegistry{}

	response := registry.RegisterChain("Foo:foo,Bar:bar,Baz:baz")

	if len(response) != 3 {
		t.Errorf("Three tokens should be registered, encountered %d", len(response))
	}

	if response[0] != registry.Tokens["Foo"]["foo"] || response[1] != registry.Tokens["Bar"]["bar"] || response[2] != registry.Tokens["Baz"]["baz"] {
		t.Errorf("Items should be registered")
	}
}
