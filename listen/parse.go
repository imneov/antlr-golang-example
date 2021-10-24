package main

import (
	"github.com/antlr/antlr4/runtime/Go/antlr"
	"github.com/imneov/antlr-golang-example/listen/parser"
	"strconv"
)

type calcListener struct {
	*parser.BaseCalcListener

	stack []Expr
}

func (l *calcListener) push(i Expr) {
	l.stack = append(l.stack, i)

}

func (l *calcListener) pop() Expr {
	if len(l.stack) < 1 {
		panic("stack is empty unable to pop")
	}

	// Get the last value from the stack.
	result := l.stack[len(l.stack)-1]

	// Remove the last element from the stack.
	l.stack = l.stack[:len(l.stack)-1]

	return result
}

func (l *calcListener) ExitMulDiv(c *parser.MulDivContext) {
	right, left := l.pop(), l.pop()
	l.push(&BinaryExpr{
		Op:  c.GetOp().GetTokenType(),
		LHS: left,
		RHS: right,
	})
}

func (l *calcListener) ExitAddSub(c *parser.AddSubContext) {
	right, left := l.pop(), l.pop()
	l.push(&BinaryExpr{
		Op:  c.GetOp().GetTokenType(),
		LHS: left,
		RHS: right,
	})
}

func (l *calcListener) ExitNumber(c *parser.NumberContext) {
	i, err := strconv.Atoi(c.GetText())
	if err != nil {
		panic(err.Error())
	}

	l.push(IntNode(i))
}

// calc takes a string expression and returns the evaluated result.
func parse(input string) Expr {
	// Setup the input
	is := antlr.NewInputStream(input)

	// Create the Lexer
	lexer := parser.NewCalcLexer(is)
	stream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)

	// Create the Parser
	p := parser.NewCalcParser(stream)

	// Finally parse the expression (by walking the tree)
	var listener calcListener
	antlr.ParseTreeWalkerDefault.Walk(&listener, p.Start())

	return listener.pop()

}
