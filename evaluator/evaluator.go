package evaluator

import (
	"monkey-study/ast"
	"monkey-study/object"
)

func Eval(node ast.Node) object.Object {
	switch node := node.(type) {
	// for statement
	case *ast.Program:
		return evalStatements(node.Statements)
	case *ast.ExpressionStatement:
		return Eval(node.Expression)
	// for expression
	case *ast.IntegerLiteral:
		return &object.Integer{Value: node.Value}
	}

	return nil
}

func evalStatements(stmts []ast.Statement) object.Object {
	var result object.Object

	for _, statement := range stmts {
		result = Eval(statement)
	}

	return result
}
