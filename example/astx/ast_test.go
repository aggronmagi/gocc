package astx

import (
	"fmt"
	"testing"

	"github.com/aggronmagi/gocc/example/astx/ast"
	"github.com/aggronmagi/gocc/example/astx/lexer"
	"github.com/aggronmagi/gocc/example/astx/parser"
)

func TestPass(t *testing.T) {
	sml, err := test([]byte("a b c d e f"))
	if err != nil {
		t.Fatal(err.Error())
	}
	fmt.Printf("output: %s\n", sml)
}

func TestFail(t *testing.T) {
	_, err := test([]byte("a b ; d e f"))
	if err == nil {
		t.Fatal("expected parse error")
	} else {
		fmt.Printf("Parsing failed as expected: %v\n", err)
	}
}

func test(src []byte) (astree ast.StmtList, err error) {
	fmt.Printf("input: %s\n", src)
	s := lexer.NewLexer(src)
	p := parser.NewParser()
	a, err := p.Parse(s)
	if err == nil {
		astree = a.(ast.StmtList)
	}
	return
}
