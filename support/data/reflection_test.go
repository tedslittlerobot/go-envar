package envarData_test

import (
	"github.com/tedslittlerobot/go-envar/support/data"
	. "gopkg.in/check.v1"
	"reflect"
	"testing"
)

type TestReflectionValueSettingStruct struct {
	Name    string
	Age     int
	IsHuman bool
}

func TestReflection(t *testing.T) { TestingT(t) }

type ReflectionTestSuite struct {
	Source     TestReflectionValueSettingStruct
	Reflection envarData.Reflection
}

var _ = Suite(&ReflectionTestSuite{})

func (s *ReflectionTestSuite) SetUpTest(c *C) {
	s.Source = TestReflectionValueSettingStruct{}
	s.Reflection = envarData.CreateReflection(&s.Source)
}

func (s *ReflectionTestSuite) TestNewReflection(c *C) {
	//
}

func (s *ReflectionTestSuite) TestSetFieldValuesSetsSimpleStringValue(c *C) {
	s.Reflection.SetFieldValues([]*envarData.Field{
		{
			"Name",
			reflect.TypeOf(""),
			[]*envarData.SourceToken{},
			"Monkey",
		},
		{
			"Age",
			reflect.TypeOf(0),
			[]*envarData.SourceToken{},
			"42",
		},
		{
			"IsHuman",
			reflect.TypeOf(false),
			[]*envarData.SourceToken{},
			"true",
		},
	})

	c.Assert(s.Source.Name, Equals, "Monkey")
	c.Assert(s.Source.Age, Equals, 42)
	c.Assert(s.Source.IsHuman, Equals, true)
}
