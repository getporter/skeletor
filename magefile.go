//go:build mage

package main

import (
	"os"

	"get.porter.sh/magefiles/mixins"
	"get.porter.sh/magefiles/releases"
)

const (
	mixinName    = "skeletor"
	mixinPackage = "github.com/getporter/skeletor"
	mixinBin     = "bin/mixins/" + mixinName
)

var magefile = mixins.NewMagefile(mixinPackage, mixinName, mixinBin)

// ConfigureAgent sets up the CI server with mage and GO
func ConfigureAgent() {
	magefile.ConfigureAgent()
}

// Build the mixin
func Build() {
	magefile.Build()
}

// XBuildAll cross-compiles the mixin before a release
func XBuildAll() {
	magefile.XBuildAll()
}

// TestUnit runs the unit tests
func TestUnit() {
	magefile.TestUnit()
}

// Test runs all types of tests
func Test() {
	magefile.Test()
}

// Publish the mixin to GitHub
func Publish() {
	// You can test out publishing locally by overriding PORTER_RELEASE_REPOSITORY and PORTER_PACKAGES_REMOTE
	if _, overridden := os.LookupEnv(releases.ReleaseRepository); !overridden {
		os.Setenv(releases.ReleaseRepository, "github.com/YOURNAME/YOURREPO")
	}
	magefile.PublishBinaries()

	// TODO: uncomment out the lines below to publish a mixin feed
	// Set PORTER_PACKAGES_REMOTE to a repository that will contain your mixin feed, similar to github.com/getporter/packages
	//if _, overridden := os.LookupEnv(releases.PackagesRemote); !overridden {
	//	os.Setenv("PORTER_PACKAGES_REMOTE", "git@github.com:YOURNAME/YOUR_PACKAGES_REPOSITORY")
	//}
	//magefile.PublishMixinFeed()
}

// TestPublish publishes the project to the specified GitHub username.
// If your mixin is official hosted in a repository under your username, you will need to manually
// override PORTER_RELEASE_REPOSITORY and PORTER_PACKAGES_REMOTE to test out publishing safely.
func TestPublish(username string) {
	magefile.TestPublish(username)
}

// Install the mixin
func Install() {
	magefile.Install()
}

// Clean removes generated build files
func Clean() {
	magefile.Clean()
}
