package envar

import (
	envarData "github.com/tedslittlerobot/go-envar/support/data"
	envarResolvers "github.com/tedslittlerobot/go-envar/support/resolvers"
)

type Envar struct {
	Resolvers    envarData.ResolverRegistry
	SourceTokens envarData.SourceTokenRegistry
}

func Make() Envar {
	e := Envar{
		Resolvers: envarData.ResolverRegistry{
			Resolvers: map[string]envarData.ResolverInterface{},
		},
		SourceTokens: envarData.SourceTokenRegistry{},
	}

	e.Resolvers.Resolvers["env"] = envarResolvers.EnvironmentVariableResolver{}
	e.Resolvers.Resolvers["default"] = envarResolvers.RawValueResolver{}

	return e
}

func (e *Envar) Apply(v interface{}) {
	reflection := envarData.CreateReflection(v)

	// get list of fields
	fields := reflection.MakeFields(&e.SourceTokens)

	for _, field := range fields {
		e.Resolvers.ResolveFieldValue(field)
	}

	reflection.SetFieldValues(fields)
}
