package hamcrest

import (
	"fmt"
	"github.com/scarabsoft/go-hamcrest/matcher"
	"strings"
	"testing"
)

type MessageInterface interface {
	String() string
}

type messageImpl struct {
	format string
	args   []interface{}
}

func (m *messageImpl) String() string {
	if m.args == nil || len(m.args) == 0 {
		return m.format
	}
	return fmt.Sprintf(m.format, m.args)
}

func Message(m string) MessageInterface {
	return &messageImpl{m, nil}
}

func MessageF(format string, args []interface{}) MessageInterface {
	return &messageImpl{format, args}
}

type Assertion struct {
	t *testing.T
}

func NewAssertion(t *testing.T) *Assertion {
	return &Assertion{t}
}

func (a *Assertion) That(actual interface{}, matcher matcher.Matcher, messages ...MessageInterface) matcher.Matcher {
	a.t.Helper()
	if !matcher.Matches(actual) {
		a.t.Error(generateErrorMessage(matcher, messages...))
	}
	return matcher
}

type Requirement struct {
	t *testing.T
}

func NewRequirement(t *testing.T) *Requirement {
	return &Requirement{t}
}

func (r *Requirement) That(actual interface{}, matcher matcher.Matcher, messages ...MessageInterface) matcher.Matcher {
	r.t.Helper()
	if !matcher.Matches(actual) {
		r.t.Fatal(generateErrorMessage(matcher, messages...))
	}
	return matcher
}

func generateErrorMessage(matcher matcher.Matcher, messages ...MessageInterface) string {
	if len(messages) == 0 {
		return matcher.Cause()
	}

	ss := make([]string, len(messages))
	for _, msg := range messages {
		ss = append(ss, msg.String())
	}

	return matcher.Cause() + " -- " + strings.Join(ss, " ")
}
