package main

import (
	"fmt"
	"go/ast"
	"go/token"
)

func checkCyclomaticComplexity(fset *token.FileSet, node ast.Node) {

	ast.Inspect(node, func(n ast.Node) bool {

		funcDecl, ok := n.(*ast.FuncDecl)
		if !ok {
			return true
		}
		complexity := 1
		ast.Inspect(funcDecl.Body, func(inner ast.Node) bool {
			if _, ok := inner.(*ast.IfStmt); ok {
				complexity++
			}
			if _, ok := inner.(*ast.ForStmt); ok {
				complexity++
			}
			if _, ok := inner.(*ast.RangeStmt); ok {
				complexity++
			}
			if _, ok := inner.(*ast.CaseClause); ok {
				complexity++
			}
			return true
		})

		fmt.Println(complexity)
		return true
	})
}
