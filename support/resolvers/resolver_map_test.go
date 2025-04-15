package envarResolvers_test

import (
	"github.com/tedslittlerobot/go-envar/support/data"
	envarResolvers "github.com/tedslittlerobot/go-envar/support/resolvers"
	. "gopkg.in/check.v1"
	"testing"
)

func TestMapResolver(t *testing.T) { TestingT(t) }

type MapResolverTestSuite struct {
	Resolver envarResolvers.MapResolver
}

var _ = Suite(&MapResolverTestSuite{})

func (s *MapResolverTestSuite) SetUpTest(c *C) {
	s.Resolver = envarResolvers.MapResolver{
		Contents: map[string]string{
			"foo": "foovalue",
			"bar": "barvalue",
			"baz": "bazvalue",
		},
	}
}

func (s *MapResolverTestSuite) TestInterface(c *C) {}

func (s *MapResolverTestSuite) TestMatchingResolution(c *C) {
	token := envarData.SourceToken{
		Driver:     "Foo",
		Key:        "foo",
		IsResolved: false,
		Value:      "",
	}

	s.Resolver.Resolve(&token)

	c.Assert(token.IsResolved, Equals, true)
	c.Assert(token.Value, Equals, "foovalue")
}

func (s *MapResolverTestSuite) TestNonMatchingResolution(c *C) {
	token := envarData.SourceToken{
		Driver:     "Foo",
		Key:        "monkeys",
		IsResolved: false,
		Value:      "",
	}

	s.Resolver.Resolve(&token)

	c.Assert(token.IsResolved, Equals, true)
	c.Assert(token.Value, Equals, "")
}
