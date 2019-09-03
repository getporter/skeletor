package skeletor

import (
	"io/ioutil"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestMixin_GetSchema(t *testing.T) {
	m := NewTestMixin(t)

	gotSchema, err := m.GetSchema()
	require.NoError(t, err)

	wantSchema, err := ioutil.ReadFile("schema/schema.json")
	require.NoError(t, err)

	assert.Equal(t, string(wantSchema), gotSchema)
}

func TestMixin_PrintSchema(t *testing.T) {
	m := NewTestMixin(t)

	err := m.PrintSchema()
	require.NoError(t, err)

	gotSchema := m.TestContext.GetOutput()

	wantSchema, err := ioutil.ReadFile("schema/schema.json")
	require.NoError(t, err)

	assert.Equal(t, string(wantSchema), gotSchema)
}
