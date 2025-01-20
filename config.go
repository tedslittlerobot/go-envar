package envar

import "github.com/tedslittlerobot/go-envar/support/resolvers"

type Config struct {
	Resolvers            map[string]envarResolvers.ResolverInterface
	WithDefaultResolvers bool
}

func ApplyDefaultsToConfig(config Config) Config {
	if config.WithDefaultResolvers {
		resolvers := config.GetDefaultResolvers()

		for key, resolver := range config.Resolvers {
			resolvers[key] = resolver
		}

		config.Resolvers = resolvers
	}

	return config
}

func (config Config) GetDefaultResolvers() map[string]envarResolvers.ResolverInterface {
	return map[string]envarResolvers.ResolverInterface{
		"env":     envarResolvers.EnvironmentVariableResolver{},
		"default": envarResolvers.RawValueResolver{},
	}
}
