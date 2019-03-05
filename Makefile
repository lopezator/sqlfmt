# MAINTAINER: David López <not4rent@gmail.com>

APP    =  sqlfmt
SHELL  := /bin/bash

VERSION ?= v1.1.0
OS      ?= linux darwin windows
GOPROXY ?= https://athens.azurefd.net

.PHONY: build
build:
	@for app in $(APP) ; do \
		for os in $(OS) ; do \
			ext=""; \
			if [ "$$os" == "windows" ]; then \
				ext=".exe"; \
			fi; \
			GOPROXY=$(GOPROXY) GOOS=$$os GOARCH=amd64 CGO_ENABLED=0 \
			go build -o ./bin/$$app-$(VERSION)-$$os-amd64$$ext \
			./cmd/$$app; \
		done; \
	done

.PHONY: sql-fmt
sql-fmt:
	@echo "Running sql-fmt..."
	@find . -name '*.sql' -not -path './vendor/*' -print0 | xargs -I '{}' -0 \
	bash -c '((sqlfmt < {}) > /dev/null || exit 1 && (sqlfmt < {}) | sponge {})'

.PHONY: sql-check
sql-check:
	@echo "Running sql-check..."
	@find . -name '*.sql' -not -path './vendor/*' -print0 | xargs -I '{}' -0 \
	bash -c 'diff {} <(sqlfmt < {})' > /dev/null || (echo "sql-check failed, run \"make sql-fmt\" and try again"; exit 1)