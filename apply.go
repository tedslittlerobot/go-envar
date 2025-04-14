package envar

import (
	"github.com/tedslittlerobot/go-envar/support/resolvers"
	"log"
)

func Apply(v interface{}, config Config) {
	config.ApplyDefaults()

	reflection := CreateReflection(v)
	registry := envarResolvers.SourceTokenRegistry{}

	// get list of fields
	fields := reflection.fields(&registry)

	for _, field := range fields {
		err := envarResolvers.ResolveFieldValue(field, config.Resolvers)

		if err != nil {
			log.Fatalf("field %s could not resolve a value", field.Name)
		}
	}

	reflection.SetFieldValues(fields)
}
