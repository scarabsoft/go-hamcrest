package hamcrest

import (
	"github.com/scarabsoft/go-hamcrest/has"
	"github.com/scarabsoft/go-hamcrest/is"
	"testing"
	"time"
)

func TestPerformanceBug_HasKey(t *testing.T) {
	assert := NewAssertion(t)

	start := time.Now()

	iteration := 100 * 1000
	data := make(map[uint64]struct{})
	for i := 0; i < iteration; i++ {
		data[uint64(i)] = struct{}{}
	}

	for i := 0; i < iteration; i++ {
		assert.That(data, has.Key(uint64(i)))
	}

	assert.That(time.Since(start), is.LessThan(1*time.Second))
}

func TestMessageF(t *testing.T) {
	assert := NewAssertion(t)

	msg := MessageF("%s %d", "hello World", 1234)

	assert.That(msg.String(), is.EqualTo("hello World 1234"))
}
