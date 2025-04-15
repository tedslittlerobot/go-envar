package envarAws_test

import (
	. "gopkg.in/check.v1"
	"testing"
)

func TestEnvironmentVariableResolver(t *testing.T) { TestingT(t) }

type EnvironmentVariableResolverTestSuite struct{}

var _ = Suite(&EnvironmentVariableResolverTestSuite{})

func (s *EnvironmentVariableResolverTestSuite) SetUpTest(c *C) {}

func (s *EnvironmentVariableResolverTestSuite) TestInterface(c *C) {}

func (s *EnvironmentVariableResolverTestSuite) TestSomething(c *C) {
	//envarAws.ParameterStoreMapResolver()
	//
}
