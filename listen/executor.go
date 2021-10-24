package main

import (
	"github.com/imneov/antlr-golang-example/listen/parser"
)

// calc takes a string expression and returns the evaluated result.
func eval(expr Expr) Node {
	switch expr := expr.(type) {
	case *BinaryExpr:
		return evalBinary(expr)
	case IntNode:
		return expr
	}
	return UNDEFINED_RESULT
}

// evalBinary eval simple types.
func evalBinary(expr *BinaryExpr) Node {
	lhs := eval(expr.LHS)
	rhs := eval(expr.RHS)
	switch lhs := lhs.(type) {
	case IntNode:
		switch rhs := rhs.(type) {
		case IntNode:
			return evalBinaryInt(expr.Op, lhs, rhs)
		}
	}
	return UNDEFINED_RESULT
}

func evalBinaryInt(op int, lhs, rhs IntNode) Node {
	switch op {
	case parser.CalcParserADD:
		return lhs + rhs
	case parser.CalcParserSUB:
		return lhs - rhs
	case parser.CalcParserMUL:
		return lhs * rhs
	case parser.CalcParserDIV:
		if rhs == 0 {
			return UNDEFINED_RESULT
		}
		return lhs / rhs
	}
	return UNDEFINED_RESULT
}
