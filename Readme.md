# GO Hamcrest
[![Code Coverage](https://codecov.io/gh/scarabsoft/go-hamcrest/branch/main/graph/badge.svg)](https://codecov.io/gh/scarabsoft/go-hamcrest)
[![Go Report Card](https://goreportcard.com/badge/github.com/scarabsoft/go-hamcrest)](https://goreportcard.com/report/github.com/scarabsoft/go-hamcrest)
[![Go Reference](https://pkg.go.dev/badge/github.com/scarabsoft/go-hamcrest.svg)](https://pkg.go.dev/github.com/scarabsoft/go-hamcrest)
[![GitHub license](https://img.shields.io/github/license/scarabsoft/go-hamcrest.svg)](https://github.com/scarabsoft/go-hamcrest/blob/main/LICENSE)

I am about to do my first steps with Go, so please don't expect this to be idiomatic. 
Coming from the Java Enterprise world(T_x). I am used to working with the hamcrest matcher framework, 
which makes it easy (at least for me) to write readable assertions. Maybe someone else finds this useful too.

**Contributions and Feedback are welcome.**


## Assertions / Requirements

```go
assert := hamcrest.NewAssertion(t)

require := hamcrest.NewRequirement(t)
```

Can be used interchangeably - the only difference is that a requirement stops the test case in case of a failing matcher,
whereas an assertion continues with the test execution.

## Matcher 

### General
 
```go
assert := hamcrest.NewAssertion(t)

assert.That(actual, is.EqualTo(10))
assert.That(actual, is.NotEqualTo(10))

assert.That(actual, is.True())
assert.That(actual, is.Ok())
assert.That(actual, is.False())

assert.That(actual, is.Nil())
assert.That(actual, is.NotNil())
assert.That(actual, is.PointingTo(expectedPtr))
```

### Numeric 

```go
assert := hamcrest.NewAssertion(t)

assert.That(actual, is.GreaterThan(10))
assert.That(actual, is.GreaterThanEqual(10))
assert.That(actual, is.LessThan(10))
assert.That(actual, is.LessThanEqual(10))

assert.That(actual, is.Between(10, 100))
assert.That(actual, is.BetweenOrEqual(10, 100))

assert.That(actual, is.NotBetween(10, 100))
assert.That(actual, is.NotBetweenOrEqual(10, 100))

assert.That(actual, is.CloseTo(9.5, 0.001))
assert.That(actual, is.NotCloseTo(9.5, 0.001))

```

### Strings
```go
assert := hamcrest.NewAssertion(t)

assert.That(acutal, has.Prefix("abc"))
assert.That(acutal, has.NotPrefix("abc"))

assert.That(acutal, has.Suffix("xyz"))
assert.That(acutal, has.NotSuffix("xyz"))

assert.That(actual, is.MatchingPattern(pattern))
assert.That(actual, is.NotMatchingPattern(pattern))

assert.That(actual, has.Length(3))

assert.That(actual, is.Empty())
assert.That(actual, is.NotEmpty())
```

### Collections

```go
assert := hamcrest.NewAssertion(t)

assert.That(actual, is.Empty())
assert.That(actual, is.NotEmpty())
assert.That(actual, has.Item(10))
assert.That(actual, has.NotItem(10))
assert.That(actual, has.Items(10,20))
assert.That(actual, has.NotItems(10,20))

assert.That(actual, has.Key("someKey"))
assert.That(actual, has.NotKey("someKey"))
assert.That(actual, has.Keys("K1", "K2"))
assert.That(actual, has.NotKeys("K1", "K2"))
```

### Boolean
```go

assert := hamcrest.NewAssertion(t)

assert.That(actual, is.EqualTo(true))
assert.That(actual, is.NotEqualTo(false))

assert.That(actual, is.True())
assert.That(actual, is.False())
assert.That(actual, is.Ok())
assert.That(actual, is.NotOk())

```

### Time
```go
assert := hamcrest.NewAssertion(t)

assert.That(actual, is.Before(expected))
assert.That(actual, is.NotBefore(expected))
assert.That(actual, is.BeforeOrEqual(expected))

assert.That(actual, is.After(expected))
assert.That(actual, is.NotAfter(expected))
assert.That(actual, is.AfterOrEqual(expected))

```

### Type

```go
assert := hamcrest.NewAssertion(t)

assert.That(actual, has.SameTypeAs(expected))
assert.That(actual, has.NotSameTypeAs(expected))
```