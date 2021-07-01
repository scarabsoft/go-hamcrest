package examples

import (
	"errors"
	"github.com/scarabsoft/go-hamcrest"
	"github.com/scarabsoft/go-hamcrest/is"
	"testing"
)

func TestPlayground(t *testing.T) {
	assert := hamcrest.NewAssertion(t)

	var err error = nil
	assert.NoError(err)

	err2 := errors.New("Some Error")
	assert.Error(err2)

	assert.That(true, is.True())
}
