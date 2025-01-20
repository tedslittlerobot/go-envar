package envarResolvers

import (
	. "gopkg.in/check.v1"
	"testing"
)

func TestRawValueResolver(t *testing.T) { TestingT(t) }

type RawValueResolverTestSuite struct {
	Resolver RawValueResolver
}

var _ = Suite(&RawValueResolverTestSuite{})

func (s *RawValueResolverTestSuite) SetUpTest(c *C) {
	s.Resolver = RawValueResolver{}
}

func (s *RawValueResolverTestSuite) TestInterface(c *C) {}

func (s *RawValueResolverTestSuite) TestResolution(c *C) {
	token := SourceToken{
		Driver:     "Foo",
		Key:        "Bar",
		IsResolved: false,
		Value:      "",
	}

	s.Resolver.Resolve(&token)

	c.Assert(token.Value, Equals, "Bar")
}
