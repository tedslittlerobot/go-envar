package envar_test

import (
	"github.com/tedslittlerobot/go-envar"
	"github.com/tedslittlerobot/go-envar/support/resolvers"
	. "gopkg.in/check.v1"
	"testing"
)

type TestBasicStruct struct {
	Foo string `envar:"never:baz,raw:foo,default:baz"`
	Bar string `envar:"raw:bar"`
}

func TestEnvar(t *testing.T) { TestingT(t) }

type EnvarTestSuite struct {
	Source      TestBasicStruct
	BasicConfig envar.Config
}

var _ = Suite(&EnvarTestSuite{})

func (s *EnvarTestSuite) SetUpTest(c *C) {
	s.Source = TestBasicStruct{}
	s.BasicConfig = envar.Config{
		map[string]envarResolvers.ResolverInterface{
			"raw":   envarResolvers.RawValueResolver{},
			"never": envarResolvers.NeverResolver{},
		},
		false,
	}
}

func (s *EnvarTestSuite) TestBasicImport(c *C) {
	envar.Apply(&s.Source, s.BasicConfig)

	c.Assert(s.Source.Foo, Equals, "foo")
	c.Assert(s.Source.Bar, Equals, "bar")
}
