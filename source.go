package envar

import (
	"log"
	"strings"
)

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

type SourceTokenRegistry struct {
	Tokens map[string]map[string]*SourceToken
}

func SplitDriverAndKey(input string) (string, string) {
	output := strings.Split(input, ":")

	if len(output) != 2 {
		log.Fatalf("Encountered syntactically invalid envar token %s", input)
	}

	return output[0], output[1]
}

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

func (registry *SourceTokenRegistry) RegisterChain(input string) []*SourceToken {
	inputs := strings.Split(input, ",")
	var outputs []*SourceToken

	for _, s := range inputs {
		outputs = append(outputs, registry.Register(s))
	}

	return outputs
}
