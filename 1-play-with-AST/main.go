package main

import (
	"bytes"
	"go/ast"
	"go/parser"
	"go/printer"
	"go/token"

	"github.com/fatih/astrewrite"
)

func main() {
	src := `package main

type Foo struct{}`

	Rename("foo.go", "Bar", src)
}

func Rename(filename, renameTo string, src interface{}) ([]byte, error) {
	fset := token.NewFileSet()
	file, err := parser.ParseFile(fset, filename, src, parser.ParseComments)
	if err != nil {
		return nil, err
	}

	rewriteFunc := func(n ast.Node) (ast.Node, bool) {
		x, ok := n.(*ast.TypeSpec)
		if !ok {
			return n, true
		}
		// change struct type name to renameTo
		x.Name.Name = renameTo
		return x, true
	}

	rewritten := astrewrite.Walk(file, rewriteFunc)

	var buf bytes.Buffer
	printer.Fprint(&buf, fset, rewritten)

	return buf.Bytes(), nil
}
