BIN  := $(CURDIR)/bin
CMD  := $(CURDIR)/cmd
DEMO := $(CURDIR)/demo
DIST := $(CURDIR)/dist

GOARCH ?= $(shell go env GOARCH)
GOOS   ?= $(shell go env GOOS)
GOBIN  := $(HOME)/.local/bin
GO     := env GOARCH=$(GOARCH) GOBIN=$(GOBIN) GOOS=$(GOOS) go
GOEXE  != $(GO) env GOEXE

all: $(BIN)/color-make

clean:
	@ $(RM) --recursive --verbose $(BIN)
	@ $(RM) --recursive --verbose $(DIST)

.PHONY: dist
dist: $(DIST)/color-make-$(GOOS)-$(GOARCH)$(GOEXE)

install: $(GOBIN)/color-make

pretty:
	$(GO) fmt ./...
	$(GO) mod tidy
	$(GO) vet ./...
	prettier --write $(CURDIR)

release: $(BIN)/color-make-$(GOOS)-$(GOARCH)$(GOEXE)

ALWAYS:

$(BIN)/color-make: $(CMD)/color-make ALWAYS
	$(GO) build -o $@ $<

$(DIST)/color-make-$(GOOS)-$(GOARCH)$(GOEXE): $(CMD)/color-make ALWAYS
	$(GO) build -o $@ $<

$(GOBIN)/color-make: $(BIN)/color-make
	@ install -D --mode="u=rwx,go=rx" --no-target-directory --verbose $< $@
