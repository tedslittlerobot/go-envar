package envar

type Config struct {
	Resolvers            map[string]ResolverInterface
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

func (config Config) GetDefaultResolvers() map[string]ResolverInterface {
	return map[string]ResolverInterface{
		"env":     EnvironmentVariableResolver{},
		"default": RawValueResolver{},
	}
}
