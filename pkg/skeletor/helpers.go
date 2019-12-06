package skeletor

import (
	"testing"

	"get.porter.sh/porter/pkg/context"
)

type TestMixin struct {
	*Mixin
	TestContext *context.TestContext
}

// NewTestMixin initializes a mixin test client, with the output buffered, and an in-memory file system.
func NewTestMixin(t *testing.T) *TestMixin {
	c := context.NewTestContext(t)
	m := &TestMixin{
		Mixin: &Mixin{
			Context: c.Context,
		},
		TestContext: c,
	}

	return m
}
