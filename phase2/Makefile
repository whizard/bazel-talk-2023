# HELP
# This will output the help for each task
# thanks to https://marmelab.com/blog/2016/02/29/auto-documented-makefile.html
.PHONY: help

help: ## Make targets
	@awk 'BEGIN {FS = ":.*?## "} /^[a-zA-Z_-]+:.*?## / {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}' $(MAKEFILE_LIST)

.DEFAULT_GOAL := help

.EXPORT_ALL_VARIABLES:

OSUPPER = $(shell uname -s 2>/dev/null | tr [:lower:] [:upper:])
DARWIN = $(strip $(findstring DARWIN, $(OSUPPER)))

.PHONY: update-project
update-project: ## Update bazel cache based on go.mod files
	bazel run //:gazelle -- update-repos -from_file=go.mod
	bazel run //:gazelle

.PHONY: build
build: update-project ## Build the binaries
	$(info Building $(OSUPPER) binaries)

ifneq ($(DARWIN),)
		bazel build //cmd/...
else # assume Linux
		bazel build //cmd/... --features=pure --features=static --platforms=@io_bazel_rules_go//go/toolchain:linux_amd64
endif

.PHONY: tests
tests: ## Run tests
ifneq ($(DARWIN),)
	bazel test ... --test_arg=-test.short
else
	bazel test ... --test_arg=-test.short --features=pure --features=static --platforms=@io_bazel_rules_go//go/toolchain:linux_amd64
endif

.PHONY: build-image
build-image: ## Builds container images
	bazel run cmd/myapp1/image

.PHONY: push-image
push-image: build-image ## Publish the images to the docker repo
	bazel run cmd/myapp1/image_push

.PHONY: tidy
tidy:	## Update go.mod and go.sum
	go mod tidy

.PHONY: prepare-commit
prepare-commit: tidy build tests ## Update dependencies, build and run tests
	go fmt ./... ## Format go code

.PHONY: clean
clean: ## Do a complete cleanup of the cache
	@echo 'cleaning up'
	@bazel clean --expunge
	

