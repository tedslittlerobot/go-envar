package envarResolvers

import (
	. "gopkg.in/check.v1"
	"os"
	"testing"
)

func TestEnvironmentVariableResolver(t *testing.T) { TestingT(t) }

type EnvironmentVariableResolverTestSuite struct {
	Resolver EnvironmentVariableResolver
}

var _ = Suite(&EnvironmentVariableResolverTestSuite{})

func (s *EnvironmentVariableResolverTestSuite) SetUpTest(c *C) {
	s.Resolver = EnvironmentVariableResolver{}
	os.Unsetenv("TEST_ENV_RESOLVER_VALUE")
}

func (s *EnvironmentVariableResolverTestSuite) TearDownTest(c *C) {
	os.Unsetenv("TEST_ENV_RESOLVER_VALUE")
}

func (s *EnvironmentVariableResolverTestSuite) TestInterface(c *C) {}

func (s *EnvironmentVariableResolverTestSuite) TestResolutionWithNoEnvValueDoesNotResolve(c *C) {
	token := SourceToken{
		Driver:     "Foo",
		Key:        "ENV_VAR_WHICH_DOES_NOT_EXIST",
		IsResolved: false,
		Value:      "",
	}

	s.Resolver.Resolve(&token)

	c.Assert(token.IsResolved, Equals, true)
	c.Assert(token.Value, Equals, "")
}

func (s *EnvironmentVariableResolverTestSuite) TestResolutionWithEnvValueResolves(c *C) {
	token := SourceToken{
		Driver:     "Foo",
		Key:        "TEST_ENV_RESOLVER_VALUE",
		IsResolved: false,
		Value:      "",
	}

	os.Setenv("TEST_ENV_RESOLVER_VALUE", "Monkeys")

	s.Resolver.Resolve(&token)

	c.Assert(token.IsResolved, Equals, true)
	c.Assert(token.Value, Equals, "Monkeys")
}
