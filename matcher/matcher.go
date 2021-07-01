package matcher

import (
	"sync"
)

type ChainBuilder = func(actual interface{}, chain Chain) Chain

type matcherImpl struct {
	matched      bool
	cause        string
	once         sync.Once
	chainBuilder ChainBuilder
}

func New(chainBuilder ChainBuilder) Matcher {
	return &matcherImpl{
		chainBuilder: chainBuilder,
	}
}

func (b *matcherImpl) Matches(actual interface{}) bool {
	b.once.Do(func() {
		r := b.chainBuilder(actual, newMatcherChain()).Exec()
		b.matched = r.Matched
		b.cause = r.Cause
	})
	return b.matched
}

func (b *matcherImpl) Cause() string {
	return b.cause
}

type Matcher interface {
	Matches(actual interface{}) bool

	Cause() string
}
