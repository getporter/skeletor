package skeletor

import (
	"fmt"

	"github.com/pkg/errors"
	yaml "gopkg.in/yaml.v2"
)

type InstallAction struct {
	Steps []InstallStep `yaml:"install"`
}

type InstallStep struct {
	InstallArguments `yaml:"azure"`
}

type InstallArguments struct {
	Step `yaml:",inline"`

	Name       string                 `yaml:"name"`
	Parameters map[string]interface{} `yaml:"parameters"`
}

func (m *Mixin) Install() error {
	payload, err := m.getPayloadData()
	if err != nil {
		return err
	}

	var action InstallAction
	err = yaml.Unmarshal(payload, &action)
	if err != nil {
		return err
	}
	if len(action.Steps) != 1 {
		return errors.Errorf("expected a single step, but got %d", len(action.Steps))
	}
	step := action.Steps[0]

	fmt.Fprintf(m.Out, "Starting installation operations: %s\n", step.Name)
	fmt.Fprintf(m.Out, "Finished installation operations: %s\n", step.Name)

	var lines []string
	for _, output := range step.Outputs {
		//TODO populate the output
		v := "SOME VALUE"
		l := fmt.Sprintf("%s=%v", output.Name, v)
		lines = append(lines, l)
	}
	m.Context.WriteOutput(lines)
	return nil
}
