
.PHONY: build

default: build

build:
	go build -o toml2cli .

examples: build
	./toml2cli --in-file=examples/example.toml --out-file=examples/urfave.go
