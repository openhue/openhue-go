# Go Command Path
GO ?= `which go`

GREEN = "\\033[0\;32m"
BOLD = "\\033[1m"
RESET = "\\033[0m"

SPEC_URL = "https://api.redocly.com/registry/bundle/openhue/openhue/v2/openapi.yaml?branch=main"

.PHONY: help
help:
	@grep -h -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "make \033[36m%-10s\033[0m %s\n", $$1, $$2}'

.PHONY: generate
generate: _check-oapi-codegen-installation ## Generates the ./openhue.gen.go client | Usage: make generate [spec=/path/to/openhue.yaml]
ifdef spec
	@echo "Code generation from $(spec)"
	@oapi-codegen --package=openhue -generate=client,types -o ./openhue.gen.go "$(spec)"
else
	@echo "Code generation from $(SPEC_URL)"
	@oapi-codegen --package=openhue -generate=client,types -o ./openhue.gen.go "$(SPEC_URL)"
endif
	@echo "\n${GREEN}${BOLD}./openhue.gen.go successfully generated 🚀${RESET}"

.PHONY: tidy
tidy: ## Tidy makes sure go.mod matches the source code in the module
	@$(GO) mod tidy
	@echo "\n${GREEN}${BOLD}go.mod successfully cleaned 🧽${RESET}"

.PHONY: test
test: ## Run the tests
	@$(GO) test -cover ./...
	@echo "\n${GREEN}${BOLD}all tests successfully passed ✅ ${RESET}"

#
# Private targets
#

.PHONY: _check-oapi-codegen-installation
_check-oapi-codegen-installation:
	@command -v oapi-codegen >/dev/null 2>&1 || { echo >&2 "oapi-codegen is not installed (https://github.com/oapi-codegen/oapi-codegen)"; exit 1; }