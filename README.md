# A Porter Mixin Skeleton

This repository contains the skeleton structure of a Porter Mixin. You can clone
this repository and use it as a starting point to build new mixins. The
structure of this project matches closely with existing Porter [Mixins](https://porter.sh/mixins).

1. Create a new repository in GitHub [using this repository as a
   template](https://help.github.com/en/articles/creating-a-repository-from-a-template).
1. We recommend using Go 1.12 without go modules.
1. Rename the `cmd/skeletor` and `pkg/skeletor` directories to `cmd/YOURMIXIN` and
   `pkg/YOURMIXIN`.
1. Find the text `github.com/deislabs/porter-skeletor/pkg/skeletor` in the repository and change it to 
    `github.com/YOURNAME/YOURREPO/pkg/YOURMIXIN`.
1. Find any remaining `skeletor` text in the repository and replace it with `YOURMIXIN`.
1. In `pkg/YOURMIXIN/version.go` replace `YOURNAME` with the name you would like displayed as the mixin
   author. This value is displayed as the author of your mixin when `porter mixins list` is run.
1. Run `dep ensure`. Check-in `Gopkg.lock` and `vendor`.
1. Run `make build xbuild test` to try out all the make targets and
   verify that everything executes without failing.
1. Run `make install` to install your mixin into the Porter home directory. If
   you don't already have Porter installed, [install](https://porter.sh/install) it first.
1. Now your mixin is installed, you are ready start customizing and iterating on
   your mixin!

## Customize your mixin

This mixin is ready to wrap an existing command-line tool. The shortest path
would be to edit `build.go` to add the instructions to download the tool
and you are all set. It will look and feel like the [gcloud](https://porter.sh/mixins/gcloud) 
or [aws](https://porter.sh/mixins/aws) mixins, both of which are built on top of the exec mixin.

Edit the `Build` function in `pkg/skeletor/build.go`.
Here you can add any Dockerfile lines that you require to download and install
additional tools, configuration files, etc necessary for your mixin. The Build
function should write the Dockerfile lines to `m.Out` which is a pipe from the
mixin back to porter.

Search for `TODO` in the code and follow the instructions to customize the mixin.

Here is an example from the aws mixin, where it downloads the latest version of
of the aws binary and installs it:

https://github.com/deislabs/porter-aws/blob/001c19bfe06d248143353a55f07a42c913579481/pkg/aws/build.go#L7

This is enough to have a working mixin. Run `make build install` and then test
it out with a bundle.

That will get you started but make sure to read the mixin developer
documentation for how to create a full featured mixin:

* [Mixin Architecture](https://porter.sh/mixin-architecture/)
* [Mixin Commands](https://porter.sh/mixin-commands/)
* [Distributing Mixins](https://porter.sh/mixin-distribution/)

## Project Structure

In the `cmd/skeletor` directory, you will find a cli built using [spf13/cobra](https://github.com/spf13/cobra). The CLI contains a go file for each basic capability a Mixin should implement:

* build
* schema
* version
* install
* upgrade
* invoke
* uninstall

Each of these command implementations have a corresponding Mixin implementation in the `pkg/skeletor` directory. Each of the commands above is wired into an empty implementation in `pkg/skeletor` that needs to be completed. In order to build a new Mixin, you need to complete these implementations with the relevant technology. For example, to build a [Cloud Formation](https://aws.amazon.com/cloudformation/) mixin, you might implement the methods in `pkg/skeletor` using the [AWS Go SDK](https://docs.aws.amazon.com/sdk-for-go/api/service/cloudformation/).

## Provided capabilities

This skeleton mixin project brings some free capabilities:

### File System Access and Context

Porter provides a [Context](https://github.com/deislabs/porter/tree/master/pkg/context) package that has helpful mechanisms for accessing the File System using [spf13/afero](https://github.com/spf13/afero). This makes it easy to provide mock File System implementations during testing. The Context package also provides a mechanism to encapsualte stdin, stdout and stderr so that they can easily be passed from `cmd/skeletor` code to implementing `pkg/skeletor` code.  

### Template and Static Asset Handling

The project already includes [Packr V2](https://github.com/gobuffalo/packr/tree/master/v2) for dealing with static files, such as templates or other content that is best modeled outside of a Go file. You can see an example of this in `pkg/skeletor/schema.go`.

### Basic Schema

The project provides an implementation of the `skeletor schema` command that is mostly functional. To fully implement this for your mixin, you simply need to provide a valid JSON schema. For reference, consult `pkg/skeletor/schema/schema.json`.

### Basic Tests

The project provides some very basic test skeletons that you can use as a starting point for building tests for your mixin.

### Makefile

The project also includes a Makefile that will can be used to both build and install the mixin. The Makefile also includes a TODO `publish` target that shows how you might publish the mixin and generate an mixin feed for easily sharing your mixin.