package envarResolvers

import (
	"errors"
	"fmt"
)

type ResolverInterface interface {
	PreLoad([]*SourceToken)
	Resolve(*SourceToken)
}

func ResolveFieldValue(field *Field, resolvers map[string]ResolverInterface) error {
	for _, token := range field.Sources {
		if !token.IsResolved {
			resolver := resolvers[token.Driver]

			resolver.Resolve(token)
		}

		if token.Value == "" {
			continue
		}

		field.Value = token.Value

		return nil
	}

	return errors.New(fmt.Sprintf("no envar value for field %s", field.Name))
}
