APP    =  sqlfmt
OS     =  amd64-linux-gnu amd64-darwin amd64-windows
SHELL  := /bin/bash

# Overridable by CI
COMMIT_SHORT     ?= $(shell git rev-parse --verify --short HEAD)
VERSION          ?= $(COMMIT_SHORT)
VERSION_NOPREFIX ?= $(shell echo $(VERSION) | sed -e 's/^[[v]]*//')

# Build vars
CGO_ENABLED = 1
GOOS :=
GOARCH :=
CMAKE_SYSTEM_NAME :=
CC :=
CMAKE_C_COMPILER :=
CXX :=
CMAKE_CXX_COMPILER :=
LDFLAGS :=
EXTRA_CMAKE_FLAGS :=
SUFFIX :=
EXT :=

#
# Common methodology based targets
#
.PHONY: build
build:
	@for app in $(APP) ; do \
		for os in $(OS) ; do \
			if [ "$$os" == "amd64-linux-gnu" ]; then \
				GOOS=linux; \
				GOARCH=amd64; \
				CMAKE_SYSTEM_NAME=Linux; \
				CC=x86_64-unknown-linux-gnu-cc; \
				CXX=x86_64-unknown-linux-gnu-c++; \
				LDFLAGS="-static-libgcc -static-libstdc++"; \
				SUFFIX=-linux-2.6.32-gnu-amd64; \
			fi; \
			if [ "$$os" == "amd64-darwin" ]; then \
				GOOS=darwin; \
				GOARCH=amd64; \
				CMAKE_SYSTEM_NAME=Darwin; \
				CC=x86_64-apple-darwin13-cc; \
				CXX=x86_64-apple-darwin13-c++; \
				EXTRA_CMAKE_FLAGS=-DCMAKE_INSTALL_NAME_TOOL=x86_64-apple-darwin13-install_name_tool; \
				SUFFIX=-darwin-10.9-amd64; \
			fi; \
			if [ "$$os" == "amd64-windows" ]; then \
				EXT=".exe"; \
				GOOS=windows; \
				GOARCH=amd64; \
				CMAKE_SYSTEM_NAME=Windows; \
				CC=x86_64-w64-mingw32-cc; \
				CXX=x86_64-w64-mingw32-c++; \
				LDFLAGS=-static; \
				SUFFIX=-windows-6.2-amd64; \
			fi; \
			CGO_ENABLED=$(CGO_ENABLED) GOOS=$$GOOS GOARCH=$$GOARCH CMAKE_SYSTEM_NAME=$$CMAKE_SYSTEM_NAME \
			CMAKE_C_COMPILER=$CC CC=$$CC CMAKE_CXX_COMPILER=$CXX CXX=$$CXX EXTRA_CMAKE_FLAGS=$$EXTRA_CMAKE_FLAGS SUFFIX=$$SUFFIX \
			go build \
				-a -x -v -ldflags "$(LDFLAGS)  \
					-X main.Version=$(VERSION_NOPREFIX) \
					-X main.GitRev=$(COMMIT_SHORT) \
				" \
				-o ./bin/$$app-$(VERSION_NOPREFIX)-$$os$$EXT \
				./cmd/$$app; \
		done; \
	done;