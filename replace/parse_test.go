package replacedemo

import (
	"go/ast"
	"go/token"
	"testing"
)

func TestReplace(t *testing.T) {
	f := Parse("/home/zhan/project/goPro/replace/document/demo.go")
	fSet := token.NewFileSet()
	ast.Print(fSet, f)
}
