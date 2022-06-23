package skeletor

import (
	"get.porter.sh/porter/pkg/portercontext"
)

const defaultClientVersion string = "v0.0.0"

type Mixin struct {
	*portercontext.Context
	ClientVersion string
	//add whatever other context/state is needed here
}

// New azure mixin client, initialized with useful defaults.
func New() (*Mixin, error) {
	return &Mixin{
		Context:       portercontext.New(),
		ClientVersion: defaultClientVersion,
	}, nil

}
