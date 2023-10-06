# healthcrm

![Linting and Tests](https://github.com/savannahghi/healthcrm/actions/workflows/ci.yaml/badge.svg)
[![Coverage Status](https://coveralls.io/repos/github/savannahghi/healthcrm/badge.svg?branch=main)](https://coveralls.io/github/savannahghi/healthcrm?branch=main)

#  healthcrm Library
healthcrm is a library that abstracts the interaction with SIL's health CRM service.

### Installing it
healthcrm is compatible with modern Go releases in module mode, with Go installed:

```bash
go get -u github.com/savannahghi/healthcrm

```
will resolve and add the package to the current development module, along with its dependencies.

Alternatively the same can be achieved if you use import in a package:

```go
import "github.com/savannahghi/healthcrm"

```
and run `go get` without parameters.

The package name is `healthcrm`


### Developing

The default branch library is `main`

We try to follow semantic versioning ( <https://semver.org/> ). For that reason,
every major, minor and point release should be _tagged_.

```
git tag -m "v0.0.1" "v0.0.1"
git push --tags
```

Continuous integration tests *must* pass on Github CI. Our coverage threshold
is >=80% i.e you *must* keep coverage above 80%.


## Environment variables

In order to run tests, you need to have an `env.sh` file similar to this one:

```bash
# Application settings
export HEALTH_CRM_AUTH_SERVER_ENDPOINT=""
export HEALTH_CRM_CLIENT_ID=""
export HEALTH_CRM_CLIENT_SECRET=""
export HEALTH_CRM_GRANT_TYPE=""
export HEALTH_CRM_USERNAME=""
export HEALTH_CRM_PASSWORD=""
```

This file *must not* be committed to version control.

It is important to _export_ the environment variables. If they are not exported,
they will not be visible to child processes e.g `go test ./...`.

These environment variables should also be set up on github CI environment variable section.

## Contributing ##
I would like to cover the entire GitHub API and contributions are of course always welcome. The
calling pattern is pretty well established, so adding new methods is relatively
straightforward. See [`CONTRIBUTING.md`](CONTRIBUTING.md) for details.

## Versioning ##

In general, enumutils follows [semver](https://semver.org/) as closely as we
can for tagging releases of the package. For self-contained libraries, the
application of semantic versioning is relatively straightforward and generally
understood. We've adopted the following
versioning policy:

* We increment the **major version** with any incompatible change to
	non-preview functionality, including changes to the exported Go API surface
	or behavior of the API.
* We increment the **minor version** with any backwards-compatible changes to
	functionality, as well as any changes to preview functionality in the GitHub
	API. GitHub makes no guarantee about the stability of preview functionality,
	so neither do we consider it a stable part of the go-github API.
* We increment the **patch version** with any backwards-compatible bug fixes.

## License ##

This library is distributed under the MIT license found in the [LICENSE](./LICENSE)
file.