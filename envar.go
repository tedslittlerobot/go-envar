package envar

import (
	"log"
	"reflect"
)

type Field struct {
	Name    string
	Type    reflect.Type
	Sources []*SourceToken
	Value   string
}

func Envar(v interface{}, config Config) {
	config = ApplyDefaultsToConfig(config)

	reflection := CreateReflection(v)
	registry := SourceTokenRegistry{}

	// get list of fields
	fields := reflection.fields(&registry)

	for _, field := range fields {
		err := ResolveFieldValue(field, config.Resolvers)

		if err != nil {
			log.Fatalf("field %s could not resolve a value", field.Name)
		}
	}

	reflection.SetFieldValues(fields)
}
