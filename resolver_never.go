package envar

type NeverResolver struct{}

func (resolver NeverResolver) Resolve(token *SourceToken) {
	token.Resolve("")
}

func (resolver NeverResolver) PreLoad(tokens []*SourceToken) {}
