package envar

import (
	envarResolvers "github.com/tedslittlerobot/go-envar/support/resolvers"
	. "gopkg.in/check.v1"
	"testing"
)

func TestConfig(t *testing.T) { TestingT(t) }

type ConfigTestSuite struct {
	Config Config
}

var _ = Suite(&ConfigTestSuite{})

func (s *ConfigTestSuite) SetUpTest(c *C) {
	s.Config = Config{}
}

func (s *ConfigTestSuite) TestGetDefaultResolvers(c *C) {
	result := s.Config.GetDefaultResolvers()

	c.Assert(result, HasLen, 2)
	c.Assert(result["env"], NotNil)
	c.Assert(result["default"], NotNil)
}

func (s *ConfigTestSuite) TestApplyDefaultsAppliesWhenResolversListIsEmpty(c *C) {
	c.Assert(s.Config.Resolvers, HasLen, 0)
	s.Config.ApplyDefaults()
	c.Assert(s.Config.Resolvers, HasLen, 2)
}

func (s *ConfigTestSuite) TestApplyDefaultsDoesNotApplyWhenResolversListIsNotEmpty(c *C) {
	s.Config.Resolvers = map[string]envarResolvers.ResolverInterface{"foo": envarResolvers.NeverResolver{}}
	c.Assert(s.Config.Resolvers, HasLen, 1)

	s.Config.ApplyDefaults()

	c.Assert(s.Config.Resolvers, HasLen, 1)
}

func (s *ConfigTestSuite) TestApplyDefaultsAppliesWhenMergeFlagIsTrue(c *C) {
	s.Config.Resolvers = map[string]envarResolvers.ResolverInterface{"foo": envarResolvers.NeverResolver{}}
	s.Config.WithDefaultResolvers = true
	c.Assert(s.Config.Resolvers, HasLen, 1)

	s.Config.ApplyDefaults()

	c.Assert(s.Config.Resolvers, HasLen, 3)
}
