package envarResolvers

import (
	. "gopkg.in/check.v1"
	"testing"
)

func TestNeverResolver(t *testing.T) { TestingT(t) }

type NeverResolverTestSuite struct {
	Resolver NeverResolver
}

var _ = Suite(&NeverResolverTestSuite{})

func (s *NeverResolverTestSuite) SetUpTest(c *C) {
	s.Resolver = NeverResolver{}
}

func (s *NeverResolverTestSuite) TestInterface(c *C) {}

func (s *NeverResolverTestSuite) TestResolution(c *C) {
	token := SourceToken{
		Driver:     "Foo",
		Key:        "Bar",
		IsResolved: false,
		Value:      "",
	}

	s.Resolver.Resolve(&token)

	c.Assert(token.IsResolved, Equals, true)
	c.Assert(token.Value, Equals, "")
}
