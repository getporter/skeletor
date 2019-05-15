# A Porter Mixin Skeleton

This repository contains the skeleton structure of a Porter Mixin. You can clone this repository and use it as a starting point to build new mixins. The structure of this project matches closely with existing Porter Mixins like [Porter Azure](https://github.com/deislabs/porter-azure), [Porter Helm](https://github.com/deislabs/porter-helm), and [Porter Terraform](https://github.com/deislabs/porter-terraform).

## Project Structure

In the `cmd/skeletor` directory, you will find a cli built using [spf13/cobra](https://github.com/spf13/cobra). The CLI contains a go file for each basic capability a Mixin should implement:

* build
* schema
* verison
* install
* upgrade
* delete

Each of these command implementations have a corresponding Mixin implemention in the `pkg/skeletor` directory. Each of the commands above is wired into an empty implementation in `pkg/skeletor` that needs to be completed. In order to build a new Mixin, you need to complete these implementations with the relevant technology. For example, to build a [Cloud Formation](https://aws.amazon.com/cloudformation/) mixin, you might implement the methods in `pkg/skeletor` using the [AWS Go SDK](https://docs.aws.amazon.com/sdk-for-go/api/service/cloudformation/).

## Provided capabilities

This skeleton mixin project brings some free capabilities:

### File System Access and Context

Porter provides a [Context](https://github.com/deislabs/porter/tree/master/pkg/context) package that has helpful mechanisms for accessing the File System using [spf13/afero](https://github.com/spf13/afero). This makes it easy to provide mock File System implementations during testing. The Context package also provides a mechanism to encapsualte stdin, stdout and stderr so that they can easily be passed from `cmd/skeletor` code to implementing `pkg/skeletor` code.  

### Template and Static Asset Handling

The project already includes [Packr V2](https://github.com/gobuffalo/packr/tree/master/v2) for dealing with static files, such as templates or other content that is best modeled outside of a Go file. You can see an example of this in `pkg/skeletor/schema.go`.

### Basic Schema

The project provides an implementation of the `skeletor schema` command that is mostly functional. To fully implement this for your mixin, you simply need to provide a valid JSON schema. For reference, consult `pkg/skeletor/schema/skeletor.json`.

### Basic Tests

The project provides some very basic test skeletons that you can use as a starting point for building tests for your mixin.

### Makefile

The project also includes a Makefile that will can be used to both build and install the mixin. The Makefile also includes a TODO `publish` target that shows how you might publish the mixin and generate an mixin feed for easily sharing your mixin.