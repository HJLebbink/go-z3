Z3_REF ?= master

all: libz3.a test

clean:
	rm -rf vendor
	rm -f libz3.a

gofmt:
	@echo "Checking code with gofmt.."
	gofmt -s *.go >/dev/null

libz3.a: vendor/z3
	@echo "HJ1"
	cd vendor/z3 && python scripts/mk_make.py -x --staticlib
	@echo "HJ2"
	cd vendor/z3/build && make
	@echo "HJ3"
	cp vendor/z3/build/libz3.a .
	@echo "HJ4"

vendor/z3:
	@echo "HJ5"
	mkdir -p vendor
	@echo "HJ6"
	git clone https://github.com/Z3Prover/z3.git vendor/z3
	@echo "HJ7"
	cd vendor/z3 && git reset --hard && git clean -fdx
	@echo "HJ8"
	cd vendor/z3 && git checkout ${Z3_REF}
	@echo "HJ9"

test: gofmt
	@echo "HJ10"
	go test -v

.PHONY: all clean libz3.a test
