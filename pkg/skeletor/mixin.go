//go:generate packr2
package skeletor

import (
	"github.com/deislabs/porter/pkg/context"
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
