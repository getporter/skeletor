package skeletor

import (
	"bytes"
	"context"
	"io/ioutil"
	"os"
	"path"
	"testing"

	"get.porter.sh/porter/pkg/test"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestMain(m *testing.M) {
	test.TestMainWithMockedCommandHandlers(m)
}

// TODO: Add test cases for supported actions, we recommend install, update, uninstall and one custom action
func TestMixin_Execute(t *testing.T) {
	testcases := []struct {
		name        string // Test case name
		file        string // Path to th test input yaml
		wantOutput  string // Name of output that you expect to be created
		wantCommand string // Full command that you expect to be called based on the input YAML
	}{
		{"action", "testdata/step-input.yaml", "VICTORY",
			"skeletor man-e-faces --species human"},
	}

	defer os.Unsetenv(test.ExpectedCommandEnv)
	for _, tc := range testcases {
		t.Run(tc.name, func(t *testing.T) {
			ctx := context.Background()
			m := NewTestMixin(t)

			m.Setenv(test.ExpectedCommandEnv, tc.wantCommand)
			mixinInputB, err := ioutil.ReadFile(tc.file)
			require.NoError(t, err)

			m.In = bytes.NewBuffer(mixinInputB)

			err = m.Execute(ctx)
			require.NoError(t, err, "execute failed")

			if tc.wantOutput == "" {
				outputs, _ := m.FileSystem.ReadDir("/cnab/app/porter/outputs")
				assert.Empty(t, outputs, "expected no outputs to be created")
			} else {
				wantPath := path.Join("/cnab/app/porter/outputs", tc.wantOutput)
				exists, _ := m.FileSystem.Exists(wantPath)
				assert.True(t, exists, "output file was not created %s", wantPath)
			}
		})
	}
}
