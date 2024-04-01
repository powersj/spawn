.DEFAULT_TARGET: all

# Used to set the version info during makefile and goreleaser builds
export version := 0.1.0-dev
export branch := $(shell git rev-parse --abbrev-ref HEAD)
export commit := $(shell git rev-parse --short=8 HEAD)
# If this path is updated, update .goreleaser.yaml as well
export internalPKG := github.com/powersj/spawn/internal
export LDFLAGS := -X $(internalPKG).Version=$(version) -X $(internalPKG).Branch=$(branch) -X $(internalPKG).Commit=$(commit) $(LDFLAGS)

.PHONY: all
all: clean lint test build

.PHONY: build
build:
	CGO_ENABLED=0 go build -o spawn -ldflags "$(LDFLAGS)"

.PHONY: clean
clean:
	rm -rf spawn coverage.out dist/

.PHONY: help
help:
	@echo 'Available Targets:'
	@echo '  all      remove artifacts, lint and test code, then build the binary'
	@echo '  build    build the spawn binary'
	@echo '  clean    delete the build and test artifacts'
	@echo '  help     print this output'
	@echo '  lint     run golangci-lint'
	@echo '  release  build packages and release to GitHub'
	@echo '  snapshot build local snapshot of packages'
	@echo '  test     run all unit tests'

.PHONY: lint
lint:
	golangci-lint run

.PHONY: release
release: clean
	goreleaser release

.PHONY: snapshot
snapshot: clean
	goreleaser release --snapshot

.PHONY: test
test:
	go test -race -short -cover -coverprofile=coverage.out ./...
