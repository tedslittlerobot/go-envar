package envar

type RawValueResolver struct{}

func (resolver RawValueResolver) Resolve(token *SourceToken) {
	token.Resolve(token.Key)
}

func (resolver RawValueResolver) PreLoad(tokens []*SourceToken) {}
