package envar

type PlainValueResolver struct{}

func (resolver PlainValueResolver) Resolve(token *SourceToken) {
	token.Resolve(token.Key)
}

func (resolver PlainValueResolver) PreLoad(tokens []*SourceToken) {}
