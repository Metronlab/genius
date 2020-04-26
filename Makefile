all: count fix-fmt check tests

install: ## install dependencies
	@go get golang.org/x/lint/golint \
		honnef.co/go/tools/cmd/staticcheck \
		github.com/client9/misspell/cmd/misspell \
		golang.org/x/tools/go/analysis/passes/shadow/cmd/shadow \
		github.com/jstemmer/go-junit-report
	go install .

tests: ## Run unit tests
	go install .
	go generate ./...
	diff ./test/core.gen.go.expected ./test/core.gen.go
	./tools/script/test.sh

fix-fmt: ## use fmt -w
	./tools/script/fix-fmt.sh

check: ## check code syntax
	./tools/script/code-checks.sh

bench: ## run benchmarks
	./tools/script/bench.sh

count: ## count lines and contributions
	./tools/script/count-code-lines.sh