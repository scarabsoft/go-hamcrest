package is

import (
	"fmt"
	"github.com/scarabsoft/go-hamcrest/internal"
	"github.com/scarabsoft/go-hamcrest/matcher"
	"time"
)

func Before(expected interface{}) matcher.Matcher {
	return timeCheck(expected, internal.IsLessThan, "want value before %s; got %s")
}

func NotBefore(expected interface{}) matcher.Matcher {
	return timeCheck(expected, func(actual, expected interface{}) bool {
		return internal.IsGreaterThanEqual(actual, expected)
	}, "want value not before %s; got %s")
}

func BeforeOrEqual(expected interface{}) matcher.Matcher {
	return timeCheck(expected, internal.IsLessThanEqual, "want value before or equal to %s; got %s")
}

func After(expected interface{}) matcher.Matcher {
	return timeCheck(expected, internal.IsGreaterThan, "want value after %s; got %s")
}

func NotAfter(expected interface{}) matcher.Matcher {
	return timeCheck(expected, internal.IsLessThanEqual, "want value not after %s; got %s")
}

func AfterOrEqual(expected interface{}) matcher.Matcher {
	return timeCheck(expected, internal.IsGreaterThanEqual, "want value after or equal to %s; got %s")
}

func failIfNotTimeType(name string, value interface{}, ) matcher.Function {
	return func() matcher.MatchResult {

		_, ok := value.(time.Time)
		if !ok {
			return matcher.Failed(fmt.Sprintf("want %s to be type of time.Time; got %s", name, internal.FormatTypeWithValue(value)))
		}

		return matcher.Next()
	}
}

func timeCheck(expected interface{}, fn func(interface{}, interface{}) bool, format string) matcher.Matcher {
	return matcher.New(
		func(actual interface{}, chain matcher.Chain) matcher.Chain {
			return chain.
				Add(
					matcher.FailIfIsNil("actual", actual),
					matcher.FailIfIsNil("expected", expected),
					failIfNotTimeType("actual", actual),
					failIfNotTimeType("expected", expected),
					func() matcher.MatchResult {
						actualTime := actual.(time.Time)
						expectedTime := expected.(time.Time)

						if !fn(actualTime.UnixNano(), expectedTime.UnixNano()) {
							return matcher.Failed(
								fmt.Sprintf(format,
									expectedTime.Format(time.RFC3339Nano),
									actualTime.Format(time.RFC3339Nano),
								),
							)
						}
						return matcher.Matched()
					})
		},
	)
}
