package envarData

import (
	"reflect"
)

// Field is a struct representing a field/property with a tag, serialised into an array of SourceToken objects.
type Field struct {
	Name    string
	Type    reflect.Type
	Sources []*SourceToken
	Value   string
}
