package skeletor

import (
	"fmt"

	"github.com/deislabs/porter-skeletor/pkg"
)

func (m *Mixin) PrintVersion() {
	fmt.Fprintf(m.Out, "Skeletor mixin %s (%s)\n", pkg.Version, pkg.Commit)
}
