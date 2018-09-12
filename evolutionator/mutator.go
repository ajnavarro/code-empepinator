package evolutionator

import (
	"math/rand"

	"github.com/robertkrimen/otto/ast"
)

type Mutator struct {
	rand *rand.Rand

	ids map[string]*ast.Identifier
}

func NewMutator(rand *rand.Rand) *Mutator {
	return &Mutator{
		rand: rand,
		ids:  make(map[string]*ast.Identifier),
	}
}

func (e *Mutator) Mutate(a *ast.Program) *ast.Program {
	for i, b := range a.Body {
		a.Body[i] = e.evalStatement(b)
	}

	return a
}

func (e *Mutator) evalStatement(stat ast.Statement) ast.Statement {
	switch t := stat.(type) {
	case *ast.FunctionStatement:
		t.Function.Body = e.evalStatement(t.Function.Body)
	case *ast.BlockStatement:
		for i, s := range t.List {
			t.List[i] = e.evalStatement(s)
		}
	case *ast.ReturnStatement:
		t.Argument = e.evalExpression(t.Argument)
	case *ast.VariableStatement:
		for i, ex := range t.List {
			t.List[i] = e.evalExpression(ex)
		}
	default:
		println(t)
		panic("STATEMENT NOT FOUND")
	}

	return stat
}

func (e *Mutator) evalExpression(expr ast.Expression) ast.Expression {
	switch t := expr.(type) {
	case *ast.BinaryExpression:
		t.Left = e.evalExpression(t.Left)
		t.Right = e.evalExpression(t.Right)
	case *ast.UnaryExpression:
		t.Operand = e.evalExpression(t.Operand)
	case *ast.VariableExpression:
		t.Initializer = e.evalExpression(t.Initializer)
	case *ast.Identifier:
		e.ids[t.Name] = t
	default:
		panic(t)
	}

	return expr
}

func (e *Mutator) hit(max int32) bool {
	return e.rand.Int31n(max) == 1
}
