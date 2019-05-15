package skeletor

import (
	"io/ioutil"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	yaml "gopkg.in/yaml.v2"
)

func TestMixin_UnmarshalInstallStep(t *testing.T) {
	b, err := ioutil.ReadFile("testdata/install-input.yaml")
	require.NoError(t, err)

	var step InstallStep
	err = yaml.Unmarshal(b, &step)
	require.NoError(t, err)

	assert.Equal(t, "Summon Minion", step.Description)
	assert.NotEmpty(t, step.Outputs)
	assert.Equal(t, Output{"VICTORY", "VICTORY_STATUS"}, step.Outputs[0])

	assert.Equal(t, "man-e-faces", step.Name)
	assert.Equal(t, map[string]interface{}{"species": "human"}, step.Parameters)
}

func TestMixin_UnmarshalInstallAction(t *testing.T) {
	b, err := ioutil.ReadFile("testdata/install-input-with-action.yaml")
	require.NoError(t, err)

	var action InstallAction
	err = yaml.Unmarshal(b, &action)
	require.NoError(t, err)

	require.Equal(t, 1, len(action.Steps))
	step := action.Steps[0]

	assert.Equal(t, "Summon Minion", step.Description)
	assert.NotEmpty(t, step.Outputs)
	assert.Equal(t, Output{"VICTORY", "VICTORY_STATUS"}, step.Outputs[0])

	assert.Equal(t, "mysql", step.Type)
	assert.Equal(t, "man-e-faces", step.Name)

	assert.Equal(t, map[string]interface{}{"species": "human"}, step.Parameters)
}
