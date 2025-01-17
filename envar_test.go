package envar

import "testing"

type TestBasicStruct struct {
	Foo string `envar:"never:baz,raw:foo,default:baz"`
	Bar string `envar:"raw:bar"`
}

func TestEnvar(t *testing.T) {
	s := TestBasicStruct{}

	c := Config{
		map[string]ResolverInterface{
			"raw":   RawValueResolver{},
			"never": NeverResolver{},
		},
		false,
	}

	Envar(&s, c)

	if s.Foo != "foo" {
		t.Errorf("Foo should be assigned the value 'foo', found [%s]", s.Foo)
	}

	if s.Bar != "bar" {
		t.Errorf("Bar should be assigned the value 'bar', found [%s]", s.Foo)
	}
}
