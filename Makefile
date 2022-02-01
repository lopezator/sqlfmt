# MAINTAINER: David López <not4rent@gmail.com>

SHELL  = /bin/bash

VERSION         = v1.3.0
BUILDER         = docker.elastic.co/beats-dev/golang-crossbuild:1.17.6-$$BUILDER_TAG-debian10
BUILD_TARGETS   = linux/amd64 windows/amd64 darwin/amd64 darwin/arm64

.PHONY: build
build:
	@for BUILD_TARGET in $(BUILD_TARGETS) ; do \
		EXT=""; \
		# Select the appropiate docker image for each BUILD_TARGET, more info here: https://github.com/elastic/golang-crossbuild \
		if [[ $$BUILD_TARGET == "linux/amd64" || $$BUILD_TARGET == "windows/amd64" ]]; then \
			BUILDER_TAG="main"; \
		elif [[ $$BUILD_TARGET == "darwin/amd64" ]]; then \
			BUILDER_TAG="darwin"; \
		elif [[ $$BUILD_TARGET == "darwin/arm64" ]]; then \
			BUILDER_TAG="darwin-arm64"; \
		fi; \
		PLATFORM=$$(sed "s/\//-/g" <<<$$BUILD_TARGET); \
		if [[ $$BUILD_TARGET == *"windows"* ]]; then \
			EXT=".exe"; \
		fi; \
		docker run -it --rm \
		-v $(GOPATH)/src/github.com/lopezator/sqlfmt:/go/src/github.com/lopezator/sqlfmt \
		-w /go/src/github.com/lopezator/sqlfmt \
		-e CGO_ENABLED=1 \
		$(BUILDER) \
		--build-cmd "go build -o ./bin/sqlfmt-$(VERSION)-$$PLATFORM$$EXT ./cmd/sqlfmt" \
		-p "$$BUILD_TARGET"; \
	done \

.PHONY: sql-fmt
sql-fmt:
	@echo "Running sql-fmt..."
	@find . -name '*.sql' -print0 | xargs -I '{}' -0 \
	bash -c '((sqlfmt < {}) > /dev/null || exit 1 && (sqlfmt < {}) | sponge {})'

.PHONY: sql-check
sql-check:
	@echo "Running sql-check..."
	@find . -name '*.sql' -print0 | xargs -I '{}' -0 \
	bash -c 'diff {} <(sqlfmt < {})' > /dev/null || (echo "sql-check failed, run \"make sql-fmt\" and try again"; exit 1)