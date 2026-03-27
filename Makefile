# ==============================================================================
# Configuration Variables
# ==============================================================================

NAME                    := eduid_ladok
VERSION                 ?= local
NEWTAG                  ?= $(VERSION)
CURRENT_BRANCH          := $(shell git rev-parse --abbrev-ref HEAD)

# Build Configuration
LDFLAGS                 := -ldflags "-w -s --extldflags '-static' -X main.version=$(VERSION)"
CGO_ENABLED_STATIC      := CGO_ENABLED=0
BUILD_OS                := linux
BUILD_ARCH              := amd64
BUILD_FLAGS             := -v

# Docker Configuration
DOCKER_TAG              := docker.sunet.se/eduid/$(NAME):$(VERSION)

# Release Guard Configuration
_RELEASE_MODE           ?=
RESERVED_TAGS           := latest testing dev

# ==============================================================================
# Phony Targets Declaration
# ==============================================================================

.PHONY: help build test \
	docker-build docker-push docker-tag docker-pull \
	start stop restart clean_docker_images \
	gosec staticcheck vulncheck \
	_check-reserved-tag \
	release release-prod check_current_branch get_release-tag \
	vscode

# ==============================================================================
# Help Target
# ==============================================================================

help: ## Show this help message
	$(info Usage: make [target] [VERSION=x.x.x])
	$(info )
	$(info Common Targets:)
	$(info   build                 - Build the binary)
	$(info   test                  - Run all tests)
	$(info   docker-build          - Build Docker image)
	$(info   docker-push           - Push Docker image)
	$(info   start                 - Start services with docker-compose)
	$(info   stop                  - Stop services)
	$(info )
	$(info Release Targets:)
	$(info   release               - Create semver tag, build & push Docker image (BUMP=major|minor|patch))
	$(info   release-prod          - Promote a release to prod (:latest))
	$(info   get_release-tag       - Show current release version from latest git tag)
	$(info )
	$(info Environment Variables:)
	$(info   VERSION               - Docker image version (default: local))
	$(info   NEWTAG                - Target tag for docker-tag operations (default: VERSION))
	$(info   BUMP                  - Version bump type: major, minor, or patch (default: patch))
	$(info   FORCE                 - Override branch check (default: false))
	@:

# ==============================================================================
# Reserved Tag Guard
# ==============================================================================
# Prevents reserved Docker tags from being used outside of release targets.
# Reserved: semver (vX.Y.Z), latest, testing, demo, dev.
# Only 'make release', 'make release-prod', and 'make release-demo' may use them.

_check-reserved-tag:
ifneq ($(_RELEASE_MODE),1)
	@for val in "$(VERSION)" "$(NEWTAG)"; do \
		if echo "$$val" | grep -qE '^v[0-9]+\.[0-9]+\.[0-9]+$$'; then \
			echo "Error: '$$val' is a reserved semver tag. Use 'make release', 'make release-prod', or 'make release-demo' instead."; exit 1; \
		fi; \
		for reserved in $(RESERVED_TAGS); do \
			if [ "$$val" = "$$reserved" ]; then \
				echo "Error: '$$val' is a reserved tag. Use 'make release', 'make release-prod', or 'make release-demo' instead."; exit 1; \
			fi; \
		done; \
	done
endif

# ==============================================================================
# Build Targets
# ==============================================================================

build: ## Build static binary
	$(info Building static binary)
	$(CGO_ENABLED_STATIC) GOOS=$(BUILD_OS) GOARCH=$(BUILD_ARCH) go build \
		$(BUILD_FLAGS) -o ./bin/$(NAME) $(LDFLAGS) ./cmd/main.go
	$(info Done)

# ==============================================================================
# Testing Targets
# ==============================================================================

test: ## Run all tests
	$(info Running tests)
	go test -v -cover ./...

# ==============================================================================
# Code Quality & Security
# ==============================================================================

gosec: ## Run gosec security scanner
	$(info Running gosec)
	gosec -color -tests ./...

staticcheck: ## Run staticcheck linter
	$(info Running staticcheck)
	staticcheck ./...

vulncheck: ## Run vulnerability checker
	$(info Running vulncheck)
	govulncheck -scan package ./...

# ==============================================================================
# Docker Compose Operations
# ==============================================================================

start: ## Start services with docker-compose
	$(info Starting services)
	docker compose -f docker-compose.yml up -d --remove-orphans

stop: ## Stop services
	$(info Stopping services)
	docker compose -f docker-compose.yml rm -s -f

restart: stop start ## Restart services

# ==============================================================================
# Docker Build / Push / Tag Targets
# ==============================================================================

docker-build: _check-reserved-tag ## Build Docker image
	$(info Building Docker image with tag: $(VERSION))
	docker build --tag $(DOCKER_TAG) --file Dockerfile .

docker-push: _check-reserved-tag ## Push Docker image
	$(info Pushing Docker image)
	docker push $(DOCKER_TAG)

docker-tag: _check-reserved-tag ## Tag Docker image with NEWTAG
	$(info Tagging $(DOCKER_TAG) -> docker.sunet.se/eduid/$(NAME):$(NEWTAG))
	docker tag $(DOCKER_TAG) docker.sunet.se/eduid/$(NAME):$(NEWTAG)

docker-pull: _check-reserved-tag ## Pull Docker image
	$(info Pulling Docker image)
	docker pull $(DOCKER_TAG)

clean_docker_images: ## Clean Docker images
	$(info Cleaning Docker images)
	docker rmi $(DOCKER_TAG) -f

# ==============================================================================
# Development Tools
# ==============================================================================

vscode: ## Set up VS Code devcontainer environment
	$(info Installing Go tools)
	go install github.com/securego/gosec/v2/cmd/gosec@latest && \
	go install golang.org/x/vuln/cmd/govulncheck@latest && \
	go install honnef.co/go/tools/cmd/staticcheck@latest && \
	go install golang.org/x/tools/cmd/deadcode@latest
	$(info Done)

# ==============================================================================
# Release Management
# ==============================================================================

BUMP                    ?= patch
FORCE                   ?=

check_current_branch: ## Verify current branch is main
	$(info Current branch: $(CURRENT_BRANCH))
ifeq ($(CURRENT_BRANCH),main)
	$(info On main branch)
else
ifneq ($(FORCE),true)
	$(error Not on main branch — use FORCE=true to override)
else
	$(warning Not on main branch — continuing because FORCE=true)
endif
endif

get_release-tag: ## Show current release version from latest git tag
	@git tag -l "v*" --sort=-v:refname | grep -E '^v[0-9]+\.[0-9]+\.[0-9]+$$' | head -n1 || echo "v0.0.0"

#### Release target
# Creates a vX.Y.Z tag by bumping the latest existing tag.
# Usage:
#   make release                       # defaults to patch bump
#   make release BUMP=minor            # minor bump
#   make release BUMP=major            # major bump
#   make release FORCE=true            # release from any branch
#   make release BUMP=minor FORCE=true # combine options
release: check_current_branch ## Create and push a git tag (BUMP=major|minor|patch)
	@echo "$(BUMP)" | grep -qE '^(major|minor|patch)$$' || \
		{ echo "Error: BUMP must be major, minor, or patch (got: $(BUMP))"; exit 1; }
	@if [ "$(FORCE)" != "true" ] && ! git diff --quiet HEAD 2>/dev/null; then \
		echo "Error: working tree is dirty — commit or stash changes first (use FORCE=true to override)"; exit 1; \
	fi
	@LATEST=$$(git tag -l "v*" --sort=-v:refname | grep -E '^v[0-9]+\.[0-9]+\.[0-9]+$$' | head -n1); \
	if [ -z "$$LATEST" ]; then \
		echo "No existing version tags found, starting at v0.0.0"; \
		LATEST="v0.0.0"; \
	fi; \
	CURRENT=$$(echo "$$LATEST" | sed 's/^v//'); \
	MAJOR=$$(echo "$$CURRENT" | cut -d. -f1); \
	MINOR=$$(echo "$$CURRENT" | cut -d. -f2); \
	PATCH=$$(echo "$$CURRENT" | cut -d. -f3); \
	case "$(BUMP)" in \
		major) MAJOR=$$((MAJOR + 1)); MINOR=0; PATCH=0 ;; \
		minor) MINOR=$$((MINOR + 1)); PATCH=0 ;; \
		patch) PATCH=$$((PATCH + 1)) ;; \
	esac; \
	NEW_TAG="v$${MAJOR}.$${MINOR}.$${PATCH}"; \
	echo ""; \
	echo "Bumping $$LATEST -> $$NEW_TAG ($(BUMP))"; \
	echo ""; \
	git tag -a "$$NEW_TAG" -m "Release $$NEW_TAG"; \
	git push origin "$$NEW_TAG"; \
	echo ""; \
	echo "==> Release $$NEW_TAG created and pushed"; \
	echo ""; \
	echo "Building and pushing Docker images for $$NEW_TAG..."; \
	echo ""; \
	$(MAKE) docker-build VERSION=$$NEW_TAG _RELEASE_MODE=1 && \
	$(MAKE) docker-push VERSION=$$NEW_TAG _RELEASE_MODE=1 && \
	$(MAKE) docker-tag VERSION=$$NEW_TAG NEWTAG=dev _RELEASE_MODE=1 && \
	$(MAKE) docker-push VERSION=dev _RELEASE_MODE=1; \
	echo ""; \
	echo "==> Docker images built and pushed for $$NEW_TAG (:dev)"; \
	echo ""

#### Prod promotion
# Promotes a version to prod by locally pulling :vX.Y.Z images
# and re-tagging/pushing as :latest. No rebuild.
# Usage:
#   make release-prod              # promotes latest vX.Y.Z tag to prod
#   make release-prod TAG=v1.2.3   # promotes v1.2.3 to prod
release-prod: ## Promote a release tag to prod (:latest)
	@set -e; \
	if [ -n "$(TAG)" ]; then \
		SRC_TAG=$$(echo "$(TAG)" | sed 's#^refs/tags/##'); \
	else \
		SRC_TAG=$$(git tag -l "v*" --sort=-v:refname | grep -E '^v[0-9]+\.[0-9]+\.[0-9]+$$' | head -n1); \
		if [ -z "$$SRC_TAG" ]; then \
			echo "Error: no version tags found. Run 'make release' first."; exit 1; \
		fi; \
	fi; \
	echo "$$SRC_TAG" | grep -qE '^v[0-9]+\.[0-9]+\.[0-9]+$$' || \
		{ echo "Error: TAG must match vX.Y.Z (got: $$SRC_TAG)"; exit 1; }; \
	echo ""; \
	echo "Promoting $$SRC_TAG -> prod (:latest)"; \
	echo ""; \
	$(MAKE) docker-pull VERSION=$$SRC_TAG _RELEASE_MODE=1 && \
	$(MAKE) docker-tag VERSION=$$SRC_TAG NEWTAG=latest _RELEASE_MODE=1 && \
	$(MAKE) docker-push VERSION=latest _RELEASE_MODE=1; \
	echo ""; \
	echo "==> Prod promotion complete for $$SRC_TAG (:latest)"; \
	echo ""

