//go:generate packr2
package skeletor

import (
	"bufio"
	"io/ioutil"

	"github.com/deislabs/porter/pkg/context"

	"github.com/pkg/errors"
)

type Mixin struct {
	*context.Context
	//add whatever other context/state is needed here
}

// New azure mixin client, initialized with useful defaults.
func New() (*Mixin, error) {
	return &Mixin{
		Context: context.New(),
	}, nil

}

func (m *Mixin) getPayloadData() ([]byte, error) {
	reader := bufio.NewReader(m.In)
	data, err := ioutil.ReadAll(reader)
	return data, errors.Wrap(err, "could not read the payload from STDIN")
}
