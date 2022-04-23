package replacedemo

import (
	"go/ast"
	"go/parser"
	"go/token"
)

func Parse(file string) *ast.File {
	fSet := token.NewFileSet()
	f, err := parser.ParseFile(fSet, file, nil, 0)
	if err != nil {
		panic(err)
	}
	return f
}
