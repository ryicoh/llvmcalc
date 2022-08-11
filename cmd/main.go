package main

import (
	"os"
	"strings"

	"github.com/ryicoh/llvmcalc"
)

func main() {
	l := new(llvmcalc.Lexer)
	l.Scanner.Init(strings.NewReader(os.Args[1]))
	llvmcalc.Parse(l)

	asm := llvmcalc.Codegen(l)
	os.WriteFile("a.ll", []byte(asm), 0600)
}
