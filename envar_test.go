package envar

import "testing"

type TestBasicStruct struct {
	Foo string `envar:"never:baz,default:foo,default:baz"`
	Bar string `envar:"default:bar"`
}

func TestEnvar(t *testing.T) {
	s := TestBasicStruct{}

	c := Config{map[string]ResolverInterface{
		"default": PlainValueResolver{},
		"never":   NeverResolver{},
	}}

	Envar(&s, c)

	if s.Foo != "foo" {
		t.Errorf("Foo should be assigned the value 'foo', found [%s]", s.Foo)
	}

	if s.Bar != "bar" {
		t.Errorf("Bar should be assigned the value 'bar', found [%s]", s.Foo)
	}
}
