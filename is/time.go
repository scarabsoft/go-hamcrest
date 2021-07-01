package is

import "github.com/scarabsoft/go-hamcrest/matcher"

func Before(given interface{}) matcher.Matcher {
	return matcher.New(
		func(actual interface{}, chain matcher.Chain) matcher.Chain {
			panic("Implement me")
		},
	)
}

func BeforeOrEqual(given interface{}) matcher.Matcher {
	return matcher.New(
		func(actual interface{}, chain matcher.Chain) matcher.Chain {
			panic("Implement me")
		},
	)
}

func NotBefore(given interface{}) matcher.Matcher {
	return matcher.New(
		func(actual interface{}, chain matcher.Chain) matcher.Chain {
			panic("Implement me")
		},
	)
}

func After(given interface{}) matcher.Matcher {
	return matcher.New(
		func(actual interface{}, chain matcher.Chain) matcher.Chain {
			panic("Implement me")
		},
	)
}

func AfterOrEqual(given interface{}) matcher.Matcher {
	return matcher.New(
		func(actual interface{}, chain matcher.Chain) matcher.Chain {
			panic("Implement me")
		},
	)
}

func NotAfter(given interface{}) matcher.Matcher {
	return matcher.New(
		func(actual interface{}, chain matcher.Chain) matcher.Chain {
			panic("Implement me")
		},
	)
}