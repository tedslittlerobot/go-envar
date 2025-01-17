package envar

type Config struct {
	Resolvers map[string]ResolverInterface
}

func ApplyDefaultsToConfig(config Config) Config {
	if len(config.Resolvers) == 0 || config.Resolvers == nil {
		config.Resolvers = map[string]ResolverInterface{
			"default": PlainValueResolver{},
		}
	}

	return config
}
