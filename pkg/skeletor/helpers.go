package skeletor

import (
	"testing"

	"get.porter.sh/porter/pkg/portercontext"
)

type TestMixin struct {
	*Mixin
	TestContext *portercontext.TestContext
}

// NewTestMixin initializes a mixin test client, with the output buffered, and an in-memory file system.
func NewTestMixin(t *testing.T) *TestMixin {
	c := portercontext.NewTestContext(t)
	m := &TestMixin{
		Mixin: &Mixin{
			Context: c.Context,
		},
		TestContext: c,
	}

	return m
}
