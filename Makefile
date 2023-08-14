PKGS := github.com/emorydu/errors
SRCDIRS := $(shell go list -f '{{.Dir}}' $(PKGS))
GO := go

ROOT_PACKAGE := github.com/emorydu/errors
ifeq ($(origin ROOT_DIR), undefined)
ROOT_DIR := $(shell pwd)
endif



check: test vet gofmt misspell unconvert staticcheck ineffassign unparam

test:
	$(GO) test $(PKGS)

vet: | test
	$(GO) vet $(PKGS)

staticcheck:
	$(GO) get honnef.co/go/tools/cmd/staticcheck
	staticcheck -checks all $(PKGS)

misspell:
	$(GO) get github.com/client9/misspell/cmd/misspell
	misspell \
		-locale GB \
		-error \
		*.md *.go

unconvert:
	$(GO) get github.com/mdempsky/unconvert
	unconvert -v $(PKGS)

ineffassign:
	$(GO) get github.com/gordonklaus/ineffassign
	find $(SRCDIRS) -name '*.go' | xargs ineffassign

pedantic: check errcheck

unparam:
	$(GO) get mvdan.cc/unparam
	unparam ./...

errcheck:
	$(GO) get github.com/kisielk/errcheck
	errcheck $(PKGS)

gofmt:
	@echo Checking code is gofmted
	@test -z "$(shell gofmt -s -l -d -e $(SRCDIRS) | tee /dev/stderr)"

.PHONY: copyright.verify
copyright.verify:
ifeq (,$(shell which addlicense 2>/dev/null))
	@echo "===========> Installing addlicense"
	@$(GO) get -u github.com/marmotedu/addlicense
endif


## verify-copyright: Verify the boilerplate headers for all files.
.PHONY: verify-copyright
verify-copyright: copyright.verify
	@echo "===========> Verifying the boilerplate headers for all files"
	@addlicense --check -f $(ROOT_DIR)/boilerplate.txt $(ROOT_DIR) --skip-dirs=third_party

## add-copyright: Ensures source code files have copyright license headers.
.PHONY: add-copyright
add-copyright: copyright.verify
	@addlicense -v -f $(ROOT_DIR)/boilerplate.txt $(ROOT_DIR) --skip-dirs=third_party


## help: Show this help info.
.PHONY: help
help: Makefile
	@echo -e "\nUsage: make <TARGETS> ...\n\nTargets:"
	@sed -n 's/^##//p' $< | column -t -s ':' |  sed -e 's/^/ /'
	@echo "$$USAGE_OPTIONS"
