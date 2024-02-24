PROJECT_NAME := minitmpl
VERSION := $(shell cat VERSION)
BUILD_DIR := build
DIST_DIR := dist
LDFLAGS := "-w -s"
VVERSION := v$(VERSION)

.PHONY: build
build:
	@echo "Building for this platform ..."
	@mkdir -p $(BUILD_DIR)
	@CGO_ENABLED=0 go build -ldflags=$(LDFLAGS) -o $(BUILD_DIR)/$(PROJECT_NAME) main.go
	@echo "Build complete!"

.PHONY: install
install: build
	@install build/$(PROJECT_NAME) /usr/local/bin
	@echo "Binary installed at /usr/local/bin/$(PROJECT_NAME)"

.PHONY: release
release:
	@echo "Check if the git working directory is clean"
	@if [ -n "$$(git status --porcelain)" ]; then \
		echo "Error: The working directory is not clean. Please commit or stash your changes."; \
		exit 1; \
	fi

	@echo "Check if the version tag already exists and does not point to HEAD"
	@if git rev-parse $(VVERSION) >/dev/null 2>&1; then \
		if ! git describe --tags --exact-match HEAD >/dev/null 2>&1; then \
			echo "Error: Version $(VVERSION) is already tagged on a different commit."; \
			exit 1; \
		fi; \
	fi

	@echo "Tag the latest commit with the version from VERSION file if not already tagged"
	@if ! git describe --tags --exact-match HEAD >/dev/null 2>&1; then \
		echo "Latest commit not tagged. Tagging with version from VERSION file..."; \
		git tag -a $(VVERSION) -m "Release $(VVERSION)"; \
		git push origin $(VVERSION); \
	fi

	@echo "Set GITEA_TOKEN variable and run goreleaser"
	@$(eval GITEA_TOKEN=$(shell pass www/behzadan.ir/git/reza/tokens/dt06-goreleaser))
	@goreleaser release


.PHONY: run
run:
	go run main.go $(filter-out $@,$(MAKECMDGOALS))
%:
	@:

clean:
	@echo "Cleaning up..."
	@rm -rf ${BUILD_DIR}
	@rm -rf ${DIST_DIR}
