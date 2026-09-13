package main

import (
	"fmt"
	"go/ast"
	"go/token"
)

func checkUncheckedErrors(fset *token.FileSet, node ast.Node) {
	ast.Inspect(node, func(n ast.Node) bool {
		block, ok := n.(*ast.BlockStmt)
		if !ok {
			return true
		}

		stmts := block.List

		for i := 0; i < len(stmts); i++ {
			assign, ok := stmts[i].(*ast.AssignStmt)
			if !ok || assign.Tok != token.DEFINE {
				continue
			}

			// Step A: does this assignment declare a variable named "err"?
			hasErr := false
			for _, lhsExpr := range assign.Lhs {
				ident, ok := lhsExpr.(*ast.Ident)
				if ok && ident.Name == "err" {
					hasErr = true
				}
			}
			if !hasErr {
				continue // no `err` declared here, nothing to check
			}

			// Step B: (your existing isProperErrCheck logic, unchanged, declared HERE)
			isProperErrCheck := false
			if i+1 < len(stmts) {
				if ifStmt, ok := stmts[i+1].(*ast.IfStmt); ok {
					if binExpr, ok := ifStmt.Cond.(*ast.BinaryExpr); ok {
						if binExpr.Op == token.NEQ {
							if ident, ok := binExpr.X.(*ast.Ident); ok {
								if ident.Name == "err" {
									isProperErrCheck = true
								}
							}
						}
					}
				}
			}

			if !isProperErrCheck {
				fmt.Printf("Unchecked error at %s\n", fset.Position(assign.Pos()))
			}
		}

		return true
	})
}
