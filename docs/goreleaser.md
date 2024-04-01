# goreleaser

goreleaser is used to create releases of the `spawn` project

## Set up

1. To install `goreleaser`, run `make deps`
2. A GitHub token is required to publish releases, but not build local snapshots
  Get a GitHub token and store it in $HOME/.config/goreleaser/github_token

## Usage

`goreleaser` **MUST** be executed from `make` to ensure the version environment
variables are passed correctly.

### Snapshot

A snapshot will run the build and release process, but will not publish the
artifacts to GitHub. Snapshots also do not require a tag and clean git repo so
these are great for development and ensuring the process works:

```shell
make snapshot
```

### Release

To release a new version of `spawn` to GitHub do the following:

1. Check out main and ensure your repo is clean
2. Create the release tag (e.g. `git tag v1.0.0`) that matches the version
  specified in the Makefile
3. Run `make release`
4. Update the version in the Makefile for the next version
5. Commit and push the updated version

If correctly set up this will create packages under `dist` using the tagged
version, upload the packages to GitHub, and create a new GitHub release.

## Builds

The configuration will build the following files:

```text
spawn-{version}-darwin-arm64.tar.gz
spawn-{version}-linux-arm64.tar.gz
spawn-{version}-linux-x86_64.tar.gz
spawn-{version}-windows-x86_64.zip
```

Each archive will contain the following files:

* spawn
* README.md
* LICENSE
