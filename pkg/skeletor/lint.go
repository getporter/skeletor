package skeletor

import (
	"context"
	"fmt"

	"get.porter.sh/porter/pkg/encoding"
	"get.porter.sh/porter/pkg/exec/builder"
	"get.porter.sh/porter/pkg/linter"

	"gopkg.in/yaml.v2"
)

const mixinName = "skeletor"

const (
	// CodeInvalidName is the linter code for when the name of the step is invalid.
	CodeInvalidName linter.Code = "skeletor-100"
)

func (m *Mixin) Lint(ctx context.Context) (linter.Results, error) {
	var input BuildInput
	err := builder.LoadAction(ctx, m.RuntimeConfig, "", func(contents []byte) (interface{}, error) {
		err := yaml.Unmarshal(contents, &input)
		return &input, err
	})
	if err != nil {
		return nil, err
	}

	results := make(linter.Results, 0)

	for _, action := range input.Actions {
		for stepNumber, step := range action.Steps {
			// TODO: Replace with your own linting logic
			if step.Name != "invalid" {
				continue
			}

			result := linter.Result{
				Level: linter.LevelError,
				Code:  CodeInvalidName,
				Location: linter.Location{
					Action:          action.Name,
					Mixin:           mixinName,
					StepNumber:      stepNumber + 1, // We index from 1 for natural counting, 1st, 2nd, etc.
					StepDescription: step.Description,
				},
				Title:   "Invalid name",
				Message: "The name cannot be 'invalid'",
				URL:     "",
			}
			results = append(results, result)
		}
	}
	return results, nil
}

func (m *Mixin) PrintLintResults(ctx context.Context) error {
	results, err := m.Lint(ctx)
	if err != nil {
		return err
	}

	b, err := encoding.MarshalJson(results)
	if err != nil {
		return fmt.Errorf("could not marshal lint results %#v: %w", results, err)
	}

	// Print the results as json to stdout for Porter to read
	resultsJson := string(b)
	fmt.Fprintln(m.Config.Out, resultsJson)

	return nil
}
