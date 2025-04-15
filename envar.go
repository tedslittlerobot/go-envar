package envar

import (
	envarData "github.com/tedslittlerobot/go-envar/support/data"
	envarResolvers "github.com/tedslittlerobot/go-envar/support/resolvers"
)

type Envar struct {
	Resolvers    envarData.ResolverRegistry
	SourceTokens envarData.SourceTokenRegistry
}

func (e *Envar) Apply(v interface{}) {
	e.SetupDefaultResolvers()

	reflection := envarData.CreateReflection(v)

	// get list of fields
	fields := reflection.MakeFields(&e.SourceTokens)

	for _, field := range fields {
		e.Resolvers.ResolveFieldValue(field)
	}

	reflection.SetFieldValues(fields)
}

func (e *Envar) SetupDefaultResolvers() {
	if e.Resolvers.DisableDefaultResolvers {
		return
	}

	_, exists := e.Resolvers.Resolvers["env"]
	if !exists {
		e.Resolvers.Resolvers["env"] = envarResolvers.EnvironmentVariableResolver{}
	}

	_, exists = e.Resolvers.Resolvers["default"]
	if !exists {
		e.Resolvers.Resolvers["default"] = envarResolvers.RawValueResolver{}
	}
}
