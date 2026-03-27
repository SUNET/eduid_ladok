NAME                    := goladok3
MODULE                  := github.com/SUNET/$(NAME)
CURRENT_BRANCH          := $(shell git rev-parse --abbrev-ref HEAD)
BUMP                    ?= patch
MSG                     ?= $(BUMP) release
FORCE                   ?=

.PHONY: help test tidy release check-branch check-clean get-version

help: ## Show this help message
	$(info Usage: make [target] [BUMP=patch|minor|major] [MSG="commit message"])
	$(info )
	$(info Targets:)
	$(info   release       Run tests, bump version, tag, and push (default: patch))
	$(info   test          Run all tests)
	$(info   tidy          Run go mod tidy)
	$(info   get-version   Show current version from git tags)
	$(info )
	$(info Examples:)
	$(info   make release MSG="Add new endpoint")
	$(info   make release BUMP=minor MSG="Add OIDC support")
	$(info   make release BUMP=major MSG="Breaking API change")
	$(info   make release BUMP=patch MSG="Fix bug" FORCE=true)
	@:

tidy:
	go mod tidy

test:
	go test -v --cover .

check-branch:
ifeq ($(CURRENT_BRANCH),main)
else
ifneq ($(FORCE),true)
	$(error Not on main branch ($(CURRENT_BRANCH)) — use FORCE=true to override)
else
	$(warning Not on main branch ($(CURRENT_BRANCH)) — continuing because FORCE=true)
endif
endif

check-clean:
ifneq ($(FORCE),true)
	@if ! git diff --quiet HEAD 2>/dev/null; then \
		echo "Error: working tree is dirty — commit or stash changes first (use FORCE=true to override)"; exit 1; \
	fi
endif

get-version:
	@git tag -l "v*" --sort=-v:refname | grep -E '^v[0-9]+\.[0-9]+\.[0-9]+$$' | head -n1 || echo "v0.0.0"

release: check-branch check-clean tidy test
	@echo "$(BUMP)" | grep -qE '^(major|minor|patch)$$' || \
		{ echo "Error: BUMP must be major, minor, or patch (got: $(BUMP))"; exit 1; }
	@git fetch --tags
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
	echo "$$LATEST -> $$NEW_TAG ($(BUMP))"; \
	echo ""; \
	git tag -a "$$NEW_TAG" -m "$(NAME) release $$NEW_TAG: $(MSG)"; \
	git push origin "$$NEW_TAG"; \
	git push origin $(CURRENT_BRANCH); \
	echo ""; \
	echo "==> Released $$NEW_TAG"; \
	echo ""; \
	curl -sf "https://proxy.golang.org/$(MODULE)/@v/$${NEW_TAG}.info" > /dev/null && \
		echo "==> Go module proxy indexed $$NEW_TAG" || \
		echo "==> Warning: proxy indexing failed (will be indexed on first fetch)"; \
	echo ""
