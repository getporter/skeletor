package azure

import (
	"fmt"
	"strings"

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

	Type          string                 `yaml:"type"`
	Template      string                 `yaml:"template"`
	Name          string                 `yaml:"name"`
	ResourceGroup string                 `yaml:"resourceGroup"`
	Parameters    map[string]interface{} `yaml:"parameters"`
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

	// Get the arm deployer
	deployer, err := m.getARMDeployer()
	if err != nil {
		return err
	}
	// Get the Template based on the arguments (type)
	t, err := deployer.FindTemplate(step.Type, step.Template)
	if err != nil {
		return err
	}
	fmt.Fprintln(m.Out, "Starting deployment operations...")
	// call Deployer.Deploy(...)
	outputs, err := deployer.Deploy(
		step.Name,
		step.ResourceGroup,
		step.Parameters["location"].(string),
		t,
		step.Parameters, //arm params
		nil,             //Tags not supported right now
	)
	if err != nil {
		return err
	}
	fmt.Fprintln(m.Out, "Finished deployment operations...")
	// ARM does some stupid stuff with output keys, turn them
	// all into upper case for better matching
	for k, v := range outputs {
		newKey := strings.ToUpper(k)
		outputs[newKey] = v
	}

	var lines []string
	for _, output := range step.Outputs {
		// ToUpper the key because of the case weirdness with ARM outputs
		v, ok := outputs[strings.ToUpper(output.Key)]
		if !ok {
			return fmt.Errorf("couldn't find output key")
		}
		l := fmt.Sprintf("%s=%v", output.Name, v)
		lines = append(lines, l)
	}
	m.Context.WriteOutput(lines)
	return nil
}
