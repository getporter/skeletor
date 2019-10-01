package skeletor

import (
	"github.com/deislabs/porter-skeletor/pkg"
	"github.com/deislabs/porter/pkg/mixin"
	"github.com/deislabs/porter/pkg/porter/version"
)

func (m *Mixin) PrintVersion(opts version.Options) error {
	metadata := mixin.Metadata{
		Name: "skeletor",
		VersionInfo: mixin.VersionInfo{
			Version: pkg.Version,
			Commit:  pkg.Commit,
			Author:  "YOURNAME",
		},
	}
	return version.PrintVersion(m.Context, opts, metadata)
}
