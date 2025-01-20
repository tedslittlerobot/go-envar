package envar

import (
	"github.com/tedslittlerobot/go-envar/support/resolvers"
	"log"
	"reflect"
	"strconv"
)

type Reflection struct {
	Source interface{}
	Type   reflect.Type
	Value  reflect.Value
}

func CreateReflection(v interface{}) Reflection {
	return Reflection{&v, reflect.TypeOf(v).Elem(), reflect.ValueOf(v).Elem()}

	//to set
	//reflect.ValueOf(i).Elem().FieldByName(propName).Set(reflect.ValueOf(propValue))
}

func (reflection Reflection) fields(registry *envarResolvers.SourceTokenRegistry) []*envarResolvers.Field {
	var o []*envarResolvers.Field

	for i := 0; i < reflection.Type.NumField(); i++ {
		field := reflection.Type.Field(i)
		tag := field.Tag.Get("envar")
		if tag == "" {
			//log.Printf("Found field with no envar tag [%s]", field.Name)
			continue
		}

		o = append(o, &envarResolvers.Field{
			Name:    field.Name,
			Type:    field.Type,
			Sources: registry.RegisterChain(tag),
			Value:   "",
		})
	}

	return o
}

func (reflection *Reflection) SetFieldValues(fields []*envarResolvers.Field) {
	for _, field := range fields {
		//log.Printf("updating field %v", field)
		reflection.Value.FieldByName(field.Name).Set(MakeValueForField(field))
	}
}

func MakeValueForField(field *envarResolvers.Field) reflect.Value {
	//log.Printf("making value for field %v", field)
	//log.Printf("Kind %v | %s", field.Type.Kind(), field.Type.String())

	original := field.Value

	switch field.Type.Kind() {
	case reflect.String:
		return reflect.ValueOf(original)
	case reflect.Int:
		value, err := strconv.Atoi(original)

		if err != nil {
			log.Fatalf("value [%s] could not be converted to int", original)
		}

		return reflect.ValueOf(value)
	case reflect.Bool:
		if original == "true" || original == "1" || original == "yes" {
			return reflect.ValueOf(true)
		}

		if original == "false" || original == "0" || original == "no" {
			return reflect.ValueOf(true)
		}

		log.Fatalf("value [%s] could not be converted to bool", original)
	}

	log.Fatalf("value [%s] of type [%s] is not compatible with envar", original, field.Type.String())

	return reflect.ValueOf("impossible")
}

//type ReflectionField struct {
//	Name string
//	Type string
//	Sources
//}
