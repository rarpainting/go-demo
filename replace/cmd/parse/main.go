package main

import (
	"go/ast"
	"go/token"
	"replacedemo"
)

type ParamType int

const (
	PT_APPID ParamType = iota + 1
	PT_CTX
)

func main() {
	f := replacedemo.Parse("/home/zhan/project/goPro/replace/document/demo.go")
	fSet := token.NewFileSet()
	ast.Print(fSet, f)
	for _, decl := range f.Decls {
		// 获取方法
		switch typeDecl := decl.(type) {
		case *ast.FuncDecl:
			specPType, specPStr := ParamType(0), ""
			// 检查有没有参数
			for _, param := range typeDecl.Type.Params.List {
				if len(param.Names) < 1 { // 没有参数名 撤
					continue
				}

				// 检查参数类型/参数名
				// TODO: 应该先检查参数类型 参数名不检查只赋值
				switch {
				case param.Names[0].Name == "appId":
					paramType, ok := param.Type.(*ast.Ident)
					if !ok {
						continue
					}
					if paramType.Name != "int" {
						continue
					}
					specPType, specPStr = PT_APPID, param.Names[0].Name

				case param.Names[0].Name == "ctx" || param.Names[0].Name == "context":
					paramType, ok := param.Type.(*ast.SelectorExpr)
					if !ok {
						continue
					}
					xType, ok := paramType.X.(*ast.Ident)
					if !ok {
						continue
					}
					if xType.Name != "context" && paramType.Sel.Name != "Context" {
						continue
					}
					specPType, specPStr = PT_CTX, param.Names[0].Name
				}
			}

			if specPType == 0 || len(specPStr) == 0 {
				continue
			}

			logLine := []int{} // 需要改动的位置
			// 先检查是否有log
			for _, line := range typeDecl.Body.List {
				switch typeLine := line.(type) {
				case *ast.ExprStmt: // 表达式
					typeX, ok := typeLine.X.(*ast.CallExpr)
					if !ok {
						continue
					}
					seFunc, ok := typeX.Fun.(*ast.SelectorExpr)
					if !ok {
						continue
					}
					pak, okPak := seFunc.X.(*ast.Ident)
					if okPak || pak.Name != "log" || seFunc.Sel.Name != "Println" {
						continue
					}

					logLine = append(logLine, 0) // TODO: 将需要添加的位置补充
				}
			}

			// 没有需要改动的位置 爬
			if len(logLine) == 0 {
				continue
			}
		}
	}
}
