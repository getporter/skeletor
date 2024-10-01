package skeletor

import (
	"bytes"
	"context"
	"os"
	"testing"

	"get.porter.sh/porter/pkg/linter"

	"github.com/stretchr/testify/require"
)

func TestMixin_Lint(t *testing.T) {
	testcases := []struct {
		name        string         // Test case name
		file        string         // Path to the test input yaml
		wantResults linter.Results // Indicates the wanted lint result
	}{
		{"valid file", "testdata/actions-input.yaml", nil},
		{"invalid name", "testdata/actions-input-fail-lint.yaml", linter.Results{
			linter.Result{
				Level: linter.LevelError,
				Location: linter.Location{
					Action:          "install",
					Mixin:           "skeletor",
					StepNumber:      1,
					StepDescription: "Summon Minion",
				},
				Code:    CodeInvalidName,
				Title:   "Invalid name",
				Message: "The name cannot be 'invalid'",
			},
		}},
	}

	for _, tc := range testcases {
		t.Run(tc.name, func(t *testing.T) {
			ctx := context.Background()
			m := NewTestMixin(t)
			mixinInputB, err := os.ReadFile(tc.file)
			require.NoError(t, err)

			m.In = bytes.NewBuffer(mixinInputB)

			results, err := m.Lint(ctx)
			require.NoError(t, err, "lint failed")

			require.Len(t, results, len(tc.wantResults))
			for _, wantResult := range tc.wantResults {
				require.Contains(t, results, wantResult)
			}
		})
	}
}
