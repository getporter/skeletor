package skeletor

import (
	"get.porter.sh/porter/pkg/runtime"
)

const defaultClientVersion string = "v0.0.0"

type Mixin struct {
	runtime.RuntimeConfig
	ClientVersion string
	//add whatever other context/state is needed here
}

// New azure mixin client, initialized with useful defaults.
func New() *Mixin {
	return &Mixin{
		RuntimeConfig: runtime.NewConfig(),
		ClientVersion: defaultClientVersion,
	}
}
