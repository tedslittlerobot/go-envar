package envarResolvers_test

import (
	"github.com/tedslittlerobot/go-envar/support/data"
	envarResolvers "github.com/tedslittlerobot/go-envar/support/resolvers"
	. "gopkg.in/check.v1"
	"testing"
)

func TestRawValueResolver(t *testing.T) { TestingT(t) }

type RawValueResolverTestSuite struct {
	Resolver envarResolvers.RawValueResolver
}

var _ = Suite(&RawValueResolverTestSuite{})

func (s *RawValueResolverTestSuite) SetUpTest(c *C) {
	s.Resolver = envarResolvers.RawValueResolver{}
}

func (s *RawValueResolverTestSuite) TestInterface(c *C) {}

func (s *RawValueResolverTestSuite) TestResolution(c *C) {
	token := envarData.SourceToken{
		Driver:     "Foo",
		Key:        "Bar",
		IsResolved: false,
		Value:      "",
	}

	s.Resolver.Resolve(&token)

	c.Assert(token.Value, Equals, "Bar")
}
