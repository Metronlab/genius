#user overridable variables
all: lint count test

install:
	go install github.com/jstemmer/go-junit-report@latest

gen:
	@go generate $(PKG)

lint:
	go fmt $(PKG)
	golangci-lint run -v $(PKG)

count:
	@bash -c $(PWD)/scripts/count-code-lines.sh

test:
	go generate ./...
	diff ./examples/tmpl/core.gen.go.expected ./examples/tmpl/core.gen.go
	diff ./examples/tmpl/types.gen.json.expected ./examples/tmpl/types.gen.json
	diff ./examples/tmpl/stdin.gen.txt.expected ./examples/tmpl/dummy.stdin.gen.txt
	@RUN=$(RUN) PKG=$(PKG) TIMEOUT=$(TIMEOUT) bash -c $(PWD)/scripts/test.sh

