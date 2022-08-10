build:
	@$(MAKE) yacc
	@$(MAKE) build-go

yacc:
	@goyacc calc.y

build-go:
	@go build -o llvmcalc ./cmd

example:
	@./llvmcalc '1 + 5'

llc:
	@llc -filetype=obj a.ll

lld-macos-arch64:
	@ld64.lld -arch arm64 -platform_version macos 12.5.0 0.0.0 -o a.out a.o

all:
	@$(MAKE) build
	@$(MAKE) example
	@$(MAKE) llc
	@$(MAKE) lld-macos-arch64
