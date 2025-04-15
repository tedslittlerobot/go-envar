package envarResolvers_test

import (
	"github.com/tedslittlerobot/go-envar/support/data"
	envarResolvers "github.com/tedslittlerobot/go-envar/support/resolvers"
	. "gopkg.in/check.v1"
	"testing"
)

func TestNeverResolver(t *testing.T) { TestingT(t) }

type NeverResolverTestSuite struct {
	Resolver envarResolvers.NeverResolver
}

var _ = Suite(&NeverResolverTestSuite{})

func (s *NeverResolverTestSuite) SetUpTest(c *C) {
	s.Resolver = envarResolvers.NeverResolver{}
}

func (s *NeverResolverTestSuite) TestInterface(c *C) {}

func (s *NeverResolverTestSuite) TestResolution(c *C) {
	token := envarData.SourceToken{
		Driver:     "Foo",
		Key:        "Bar",
		IsResolved: false,
		Value:      "",
	}

	s.Resolver.Resolve(&token)

	c.Assert(token.IsResolved, Equals, true)
	c.Assert(token.Value, Equals, "")
}
