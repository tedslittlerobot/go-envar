package envar_test

import (
	"github.com/tedslittlerobot/go-envar"
	"github.com/tedslittlerobot/go-envar/support/data"
	"github.com/tedslittlerobot/go-envar/support/resolvers"
	. "gopkg.in/check.v1"
	"testing"
)

type TestEnvarBasicStruct struct {
	Foo string `envar:"never:baz,always:foo,foobar:foobar"`
	Bar string `envar:"always:bar"`
}

func TestEnvar(t *testing.T) { TestingT(t) }

type EnvarTestSuite struct {
	Source    TestEnvarBasicStruct
	Resolvers envarData.ResolverRegistry
}

var _ = Suite(&EnvarTestSuite{})

func (s *EnvarTestSuite) SetUpTest(c *C) {
	s.Source = TestEnvarBasicStruct{}
	s.Resolvers = envarData.ResolverRegistry{
		Resolvers: map[string]envarData.ResolverInterface{
			"always": envarResolvers.RawValueResolver{},
			"never":  envarResolvers.NeverResolver{},
		},
		DisableDefaultResolvers: true,
	}
}

func (s *EnvarTestSuite) TestApply(c *C) {
	e := envar.Envar{
		Resolvers: s.Resolvers,
	}

	e.Apply(&s.Source)

	c.Assert(s.Source.Foo, Equals, "foo")
	c.Assert(s.Source.Bar, Equals, "bar")
}

func (s *EnvarTestSuite) TestSetupDefaultResolvers(c *C) {
	e := envar.Envar{
		Resolvers: envarData.ResolverRegistry{
			Resolvers:               map[string]envarData.ResolverInterface{},
			DisableDefaultResolvers: true,
		},
	}

	e.SetupDefaultResolvers()

	c.Assert(len(e.Resolvers.Resolvers), Equals, 0)

	e.Resolvers.DisableDefaultResolvers = false

	e.SetupDefaultResolvers()

	c.Assert(len(e.Resolvers.Resolvers), Equals, 2)
}
