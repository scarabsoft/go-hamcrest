


## Assertions

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
assert.That(actual, is.PointingTo(givenPtr))
assert.That(actual, is.NotPointingTo(givenPtr))
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

assert.That(actual, is.GreaterThan(10))
assert.That(actual, is.GreaterThanEqual(10))
assert.That(actual, is.LessThan(10))
assert.That(actual, is.LessThanEqual(10))
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

### Panic

### Error
```go
assert := hamcrest.NewAssertion(t)

err := someFunction()

assert.Error(err)
assert.NoError(err)
```

### Time
```go
assert := hamcrest.NewAssertion(t)

assert.That(actual, is.Before(given))
assert.That(actual, is.NotBefore(given))

assert.That(actual, is.After(given))
assert.That(actual, is.NotAfter(given))

assert.That(actual, is.Between(min, max))
assert.That(actual, is.NotBetween(min, max))
```

### Type

```go
assert := hamcrest.NewAssertion(t)

assert.That(actual, has.SameTypeAs(given))
assert.That(actual, has.NotSameTypeAs(given))
```