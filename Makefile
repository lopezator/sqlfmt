# MAINTAINER: David López <not4rent@gmail.com>

APP    =  sqlfmt
SHELL  := /bin/bash

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