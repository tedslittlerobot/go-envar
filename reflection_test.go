package envar

import (
	"fmt"
	"reflect"
	"testing"
)

type TestReflectionValueSettingStruct struct {
	Name    string
	Age     int
	IsHuman bool
}

func TestNewReflection(t *testing.T) {
	ts := TestReflectionValueSettingStruct{}

	CreateReflection(&ts)
}

func TestSetFieldValuesSetsSimpleStringValue(t *testing.T) {
	ts := TestReflectionValueSettingStruct{}

	reflection := CreateReflection(&ts)

	reflection.SetFieldValues([]*Field{
		{
			"Name",
			reflect.TypeOf(""),
			[]*SourceToken{},
			"Monkey",
		},
	})

	if ts.Name != "Monkey" {
		t.Errorf(fmt.Sprintf("name should have the value of 'Monkey' applied, found [%s]", ts.Name))
	}
}

func TestSetFieldValuesSetsSimpleIntValue(t *testing.T) {
	ts := TestReflectionValueSettingStruct{}

	reflection := CreateReflection(&ts)

	reflection.SetFieldValues([]*Field{
		{
			"Age",
			reflect.TypeOf(0),
			[]*SourceToken{},
			"42",
		},
	})

	if ts.Age != 42 {
		t.Errorf(fmt.Sprintf("Age should have the value of 42 applied, found [%d]", ts.Age))
	}
}

func TestSetFieldValuesSetsSimpleBoolValue(t *testing.T) {
	ts := TestReflectionValueSettingStruct{}

	reflection := CreateReflection(&ts)

	reflection.SetFieldValues([]*Field{
		{
			"IsHuman",
			reflect.TypeOf(false),
			[]*SourceToken{},
			"true",
		},
	})

	if !ts.IsHuman {
		t.Errorf(fmt.Sprintf("IsHuman should have the value of true applied, found [%v]", ts.IsHuman))
	}
}
