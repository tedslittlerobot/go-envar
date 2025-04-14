package envarResolvers_test

import (
	envarResolvers "github.com/tedslittlerobot/go-envar/support/resolvers"
	. "gopkg.in/check.v1"
	"testing"
)

func TestSourceRegistry(t *testing.T) { TestingT(t) }

type SourceRegistryTestSuite struct {
	Registry envarResolvers.SourceTokenRegistry
}

var _ = Suite(&SourceRegistryTestSuite{})

func (s *SourceRegistryTestSuite) SetUpTest(c *C) {
	s.Registry = envarResolvers.SourceTokenRegistry{}
}

func (s *SourceRegistryTestSuite) TestSourceTokenCollectionRegister(c *C) {
	response := s.Registry.Register("Foo:Bar")

	c.Assert(response.Driver, Equals, "Foo")
	c.Assert(response.Key, Equals, "Bar")
	c.Assert(s.Registry.Tokens["Foo"], NotNil)
	c.Assert(s.Registry.Tokens["Foo"]["Bar"], NotNil)

	response.Value = "Monkeys"

	c.Assert(s.Registry.Tokens["Foo"]["Bar"].Value, Equals, "Monkeys")
}

func (s *SourceRegistryTestSuite) TestSourceTokenCollectionRegisterChain(c *C) {
	response := s.Registry.RegisterChain("Foo:foo,Bar:bar,Baz:baz")

	c.Assert(response, HasLen, 3)

	c.Assert(s.Registry.Tokens["Foo"], NotNil)
	c.Assert(s.Registry.Tokens["Bar"], NotNil)
	c.Assert(s.Registry.Tokens["Baz"], NotNil)
	c.Assert(s.Registry.Tokens["Foo"]["foo"], Equals, response[0])
	c.Assert(s.Registry.Tokens["Bar"]["bar"], Equals, response[1])
	c.Assert(s.Registry.Tokens["Baz"]["baz"], Equals, response[2])
}
