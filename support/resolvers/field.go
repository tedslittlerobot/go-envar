package envarResolvers

import "reflect"

type Field struct {
	Name    string
	Type    reflect.Type
	Sources []*SourceToken
	Value   string
}
