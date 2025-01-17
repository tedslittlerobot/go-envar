package envar

import (
	"os"
	"testing"
)

func TestEnvResolverInterfaceMatch(t *testing.T) {
	var _ ResolverInterface = EnvResolver{}
}

func TestEnvResolverWithNoEnvValue(t *testing.T) {
	d := EnvResolver{}

	os.Unsetenv("TEST_ENV_RESOLVER_VALUE")

	token := SourceToken{
		Driver:     "Foo",
		Key:        "Bar",
		IsResolved: false,
		Value:      "",
	}

	d.Resolve(&token)

	if !token.IsResolved || token.Value != "" {
		t.Error("Token resolution didn't go to plan")
	}
}

func TestEnvResolverWithEnvValue(t *testing.T) {
	d := EnvResolver{}

	os.Setenv("TEST_ENV_RESOLVER_VALUE", "Monkeys")

	token := SourceToken{
		Driver:     "Foo",
		Key:        "TEST_ENV_RESOLVER_VALUE",
		IsResolved: false,
		Value:      "",
	}

	d.Resolve(&token)

	os.Unsetenv("TEST_ENV_RESOLVER_VALUE")

	if !token.IsResolved || token.Value != "Monkeys" {
		t.Errorf("Token somehow got resolved to value [%s] instead of Monkeys", token.Value)
	}
}
