//go:generate packr2
package skeletor

import (
	"get.porter.sh/porter/pkg/context"
)

// REMOVE THIS CODE BELOW IF YOUR MIXIN DON\'T REQUIRE A CLIENT VERSION
// Change the default client version that fit your need
const defaultClientVersion string = "v0.0.0"

type Mixin struct {
	*context.Context
	ClientVersion string
	//add whatever other context/state is needed here
}

// New azure mixin client, initialized with useful defaults.
func New() (*Mixin, error) {
	return &Mixin{
		Context: context.New(),

		// REMOVE THIS CODE BELOW IF YOUR MIXIN DON\'T REQUIRE A CLIENT VERSION
		ClientVersion: defaultClientVersion,

		//add whatever is needed here
	}, nil

}
