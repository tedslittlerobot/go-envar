package envar

import "github.com/tedslittlerobot/go-envar/support/resolvers"

type Config struct {
	Resolvers            map[string]envarResolvers.ResolverInterface
	WithDefaultResolvers bool
}

func (config *Config) ApplyDefaults() {
	if config.WithDefaultResolvers || len(config.Resolvers) == 0 {
		resolvers := config.GetDefaultResolvers()

		for key, resolver := range config.Resolvers {
			resolvers[key] = resolver
		}

		config.Resolvers = resolvers
	}
}

func (config *Config) GetDefaultResolvers() map[string]envarResolvers.ResolverInterface {
	return map[string]envarResolvers.ResolverInterface{
		"env":     envarResolvers.EnvironmentVariableResolver{},
		"default": envarResolvers.RawValueResolver{},
	}
}
