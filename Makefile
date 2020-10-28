all: count fix-fmt check test

install: ## install dependencies
	@go get golang.org/x/lint/golint \
		honnef.co/go/tools/cmd/staticcheck \
		github.com/client9/misspell/cmd/misspell \
		golang.org/x/tools/go/analysis/passes/shadow/cmd/shadow \
		github.com/jstemmer/go-junit-report
	go install .

test: ## Run unit tests
	go install .
	go generate ./...
	diff ./examples/tmpl/core.gen.go.expected ./examples/tmpl/core.gen.go
	diff ./examples/tmpl/types.gen.json.expected ./examples/tmpl/types.gen.json
	diff ./examples/tmpl/stdin.gen.txt.expected ./examples/tmpl/dummy.stdin.gen.txt
	./tools/script/test.sh

fix-fmt: ## use fmt -w
	./tools/script/fix-fmt.sh

check: ## check code syntax
	./tools/script/code-checks.sh

bench: ## run benchmarks
	./tools/script/bench.sh

count: ## count lines and contributions
	./tools/script/count-code-lines.sh