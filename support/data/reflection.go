package envarData

import (
	"fmt"
	"reflect"
	"strconv"
)

// Reflection is used to make a Field
type Reflection struct {
	Source interface{}
	Type   reflect.Type
	Value  reflect.Value
}

func CreateReflection(v interface{}) Reflection {
	return Reflection{&v, reflect.TypeOf(v).Elem(), reflect.ValueOf(v).Elem()}
}

// MakeFields constructs an array of Field structs for each property of the reflected struct
func (reflection *Reflection) MakeFields(registry *SourceTokenRegistry) []*Field {
	var o []*Field

	for i := 0; i < reflection.Type.NumField(); i++ {
		field := reflection.Type.Field(i)
		tag := field.Tag.Get("envar")
		if tag == "" {
			//log.Printf("Found field with no envar tag [%s]", field.Name)
			continue
		}

		o = append(o, &Field{
			Name:    field.Name,
			Type:    field.Type,
			Sources: registry.RegisterChain(tag),
			Value:   "",
		})
	}

	return o
}

func (reflection *Reflection) SetFieldValues(fields []*Field) {
	for _, field := range fields {
		//log.Printf("updating field %v", field)
		reflection.Value.FieldByName(field.Name).Set(MakeValueForField(field))
	}
}

func MakeValueForField(field *Field) reflect.Value {
	//log.Printf("making value for field %v", field)
	//log.Printf("Kind %v | %s", field.Type.Kind(), field.Type.String())

	original := field.Value

	switch field.Type.Kind() {
	case reflect.String:
		return reflect.ValueOf(original)
	case reflect.Int:
		value, err := strconv.Atoi(original)

		if err != nil {
			panic(fmt.Sprintf("value [%s] could not be converted to int", original))
		}

		return reflect.ValueOf(value)
	case reflect.Bool:
		if original == "true" || original == "1" || original == "yes" {
			return reflect.ValueOf(true)
		}

		if original == "false" || original == "0" || original == "no" {
			return reflect.ValueOf(true)
		}

		panic(fmt.Sprintf("value [%s] could not be converted to bool", original))
	default:
		panic(fmt.Sprintf("Envar does not support fields of type %s", field.Type.Kind()))
	}
}
