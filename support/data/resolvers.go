package envarData

import (
	"fmt"
)

type ResolverInterface interface {
	PreLoad([]*SourceToken)
	Resolve(*SourceToken)
}

type ResolverRegistry struct {
	Resolvers map[string]ResolverInterface
}

func (r *ResolverRegistry) GetResolver(name string) ResolverInterface {
	resolver, exists := r.Resolvers[name]

	if exists {
		return resolver
	}

	panic(fmt.Sprintf("no resolver found for %s", name))
}

func (r *ResolverRegistry) AddResolver(name string, resolverInterface ResolverInterface) *ResolverRegistry {
	r.Resolvers[name] = resolverInterface

	return r
}

func (r *ResolverRegistry) AddResolvers(resolvers map[string]ResolverInterface) *ResolverRegistry {
	for name, resolverInterface := range resolvers {
		r.Resolvers[name] = resolverInterface
	}

	return r
}

// ResolveFieldValue takes a field struct, and resolves it according to the specified resolver map, keyed by Driver name
func (r *ResolverRegistry) ResolveFieldValue(field *Field) {
	for _, token := range field.Sources {
		if !token.IsResolved {
			r.GetResolver(token.Driver).Resolve(token)
		}

		if token.Value == "" {
			continue
		}

		field.Value = token.Value

		return
	}

	panic(fmt.Sprintf("no envar value for field %s", field.Name))
}
