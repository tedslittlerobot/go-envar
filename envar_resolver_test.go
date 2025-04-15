package envar_test

import (
	"github.com/tedslittlerobot/go-envar"
	"github.com/tedslittlerobot/go-envar/support/data"
	envarResolvers "github.com/tedslittlerobot/go-envar/support/resolvers"
	. "gopkg.in/check.v1"
	"testing"
)

func TestEnvarResolvers(t *testing.T) { TestingT(t) }

type EnvarResolversTestSuite struct {
	Source    TestEnvarBasicStruct
	Resolvers envarData.ResolverRegistry
}

var _ = Suite(&EnvarResolversTestSuite{})

func (s *EnvarResolversTestSuite) SetUpTest(c *C) {
	s.Source = TestEnvarBasicStruct{}
	s.Resolvers = envarData.ResolverRegistry{
		Resolvers: map[string]envarData.ResolverInterface{
			"always": envarResolvers.RawValueResolver{},
			"never":  envarResolvers.NeverResolver{},
		},
	}
}

func (s *EnvarResolversTestSuite) TestRegister(c *C) {
	e := envar.Make()

	e.Resolvers.AddResolver("foo", envarResolvers.EnvironmentVariableResolver{})
	e.Resolvers.AddResolver("bar", envarResolvers.EnvironmentVariableResolver{})
	e.Resolvers.AddResolvers(map[string]envarData.ResolverInterface{
		"foo":    envarResolvers.EnvironmentVariableResolver{},
		"foobar": envarResolvers.EnvironmentVariableResolver{},
	})

	c.Assert(len(e.Resolvers.Resolvers), Equals, 5)
}
