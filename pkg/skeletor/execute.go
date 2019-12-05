package skeletor

import (
	"get.porter.sh/porter/pkg/exec/builder"
	yaml "gopkg.in/yaml.v2"
)

func (m *Mixin) loadAction() (*Action, error) {
	var action Action
	err := builder.LoadAction(m.Context, "", func(contents []byte) (interface{}, error) {
		err := yaml.Unmarshal(contents, &action)
		return &action, err
	})
	return &action, err
}

func (m *Mixin) Execute() error {
	action, err := m.loadAction()
	if err != nil {
		return err
	}

	_, err = builder.ExecuteSingleStepAction(m.Context, action)
	return err
}
