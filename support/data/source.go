package envarData

import (
	"fmt"
	"strings"
)

// SourceToken is a struct representing a single resolvable source and its eventual value.
// It is instantiated directly from an element of the reflected tag string.
// On the property 'SqlDatabase string `envar:"foo:BAR"`', for example, Driver would be foo, Key would be BAR.
type SourceToken struct {
	Driver     string
	Key        string
	IsResolved bool
	Value      string
}

func (token *SourceToken) Resolve(value string) {
	token.Value = value
	token.IsResolved = true
}

// SourceTokenRegistry is a collection of SourceToken instances, arranged in a 2d map,
// indexed by the Driver and Key properties.
// This ensures that each Driver/Key pair is only registered and instantiated once.
type SourceTokenRegistry struct {
	Tokens map[string]map[string]*SourceToken
}

// RegisterChain takes an input string, and will return an array of SourceToken instances for that string.
// Critically, it will return an existing instance if one has already been created (based on the Driver & Key pair).
func (registry *SourceTokenRegistry) RegisterChain(input string) []*SourceToken {
	inputs := strings.Split(input, ",")
	var outputs []*SourceToken

	for _, s := range inputs {
		outputs = append(outputs, registry.Register(s))
	}

	return outputs
}

func SplitDriverAndKey(input string) (string, string) {
	output := strings.Split(input, ":")

	if len(output) != 2 {
		panic(fmt.Sprintf("Encountered syntactically invalid envar token %s", input))
	}

	return output[0], output[1]
}

// Register registers and/or resolves a SourceToken instance out of an input string
func (registry *SourceTokenRegistry) Register(input string) *SourceToken {
	driver, key := SplitDriverAndKey(input)

	//log.Printf("Key Pair from [%s] %s : %s", input, driver, key)

	if registry.Tokens == nil {
		registry.Tokens = map[string]map[string]*SourceToken{}
	}

	if registry.Tokens[driver] == nil {
		registry.Tokens[driver] = map[string]*SourceToken{}
	}

	if registry.Tokens[driver][key] == nil {
		registry.Tokens[driver][key] = &SourceToken{driver, key, false, ""}
	}

	return registry.Tokens[driver][key]
}
