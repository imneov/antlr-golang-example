package main

var (
	UNDEFINED_RESULT = &DefaultNode{typ: Undefined}
	NULL_RESULT      = &DefaultNode{typ: Undefined}
)

// Type node type
type Type int

const (
	// Undefine is Not a value
	// This isn't explicitly representable in JSON except by omitting the value.
	Undefined Type = iota
	// Int is json number, a discrete Int
	Int
)

//Expr
type Expr interface {
	expr()
}

func (*BinaryExpr) expr() {}

//BinaryExpr
type BinaryExpr struct {
	Op  int
	LHS Expr
	RHS Expr
}

//Node
type Node interface {
	node()
}

type IntNode int64

func (r IntNode) Type() Type { return Int }
func (IntNode) node()        {}
func (IntNode) expr()        {}

//DefaultNode interface
type DefaultNode struct {
	// Type is the json type
	typ Type
	// raw is the raw json
	raw string
}

func (r DefaultNode) Type() Type { return r.typ }
func (*DefaultNode) node()       {}
func (*DefaultNode) expr()       {}

func print(node Node) interface{} {
	switch node := node.(type) {
	case IntNode:
		return node
	}
	return nil
}
