package is

import (
	"fmt"
	"github.com/scarabsoft/go-hamcrest/internal"
	"github.com/scarabsoft/go-hamcrest/matcher"
	"time"
)

func Before(given interface{}) matcher.Matcher {
	return timeCheck(given, internal.IsLessThan, "want value before %s; got %s")
}

func NotBefore(given interface{}) matcher.Matcher {
	return timeCheck(given, func(actual, given interface{}) bool {
		return internal.IsGreaterThanEqual(actual, given)
	}, "want value not before %s; got %s")
}

func BeforeOrEqual(given interface{}) matcher.Matcher {
	return timeCheck(given, internal.IsLessThanEqual, "want value before or equal to %s; got %s")
}

func After(given interface{}) matcher.Matcher {
	return timeCheck(given, internal.IsGreaterThan, "want value after %s; got %s")
}

func NotAfter(given interface{}) matcher.Matcher {
	return timeCheck(given, internal.IsLessThanEqual, "want value not after %s; got %s")
}

func AfterOrEqual(given interface{}) matcher.Matcher {
	return timeCheck(given, internal.IsGreaterThanEqual, "want value after or equal to %s; got %s")
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

func timeCheck(given interface{}, fn func(interface{}, interface{}) bool, format string) matcher.Matcher {
	return matcher.New(
		func(actual interface{}, chain matcher.Chain) matcher.Chain {
			return chain.
				Add(
					matcher.FailIfIsNil("actual", actual),
					matcher.FailIfIsNil("given", given),
					failIfNotTimeType("actual", actual),
					failIfNotTimeType("given", given),
					func() matcher.MatchResult {
						actualTime := actual.(time.Time)
						givenTime := given.(time.Time)

						if !fn(actualTime.UnixNano(), givenTime.UnixNano()) {
							return matcher.Failed(
								fmt.Sprintf(format,
									givenTime.Format(time.RFC3339Nano),
									actualTime.Format(time.RFC3339Nano),
								),
							)
						}
						return matcher.Matched()
					})
		},
	)
}
