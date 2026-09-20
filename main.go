package main

import (
	"fmt"
	antlr4code "mj/ANTLR4_Code"
	"mj/compiler"
	"os"
	"strings"
)

func main() {
	t, _ := os.ReadFile("tests.mj")
	s := string(t)
	s = strings.ReplaceAll(s, "\r\n", "\n")
	Tree, _, _ := antlr4code.GenerateTheAST(s)

	instractions := compiler.Compile(Tree)
	fmt.Println(instractions.PatchAll())
}
