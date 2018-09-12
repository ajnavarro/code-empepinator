package evolutionator

import (
	"fmt"
	"math/rand"

	"github.com/ajnavarro/code-empepinator/sandbox"

	"github.com/MaxHalford/eaopt"
	"github.com/robertkrimen/otto/ast"
)

type Pair struct {
	Input  []float64
	Output float64
}

var _ eaopt.Genome = jsGenome{}

// See https://github.com/MaxHalford/eaopt#implementing-the-genome-interface
type jsGenome struct {
	name  string
	ast   *ast.Program
	pairs []*Pair
}

func (g jsGenome) Evaluate() (float64, error) {
	j := sandbox.NewJavascript(g.name)
	var out []float64
	for _, p := range g.pairs {
		v, err := j.ExecuteAST(g.ast, p.Input...)
		if err != nil {
			fmt.Println("ERROR EXECUTING AST", err.Error())
			return 0, nil
		}
		out = append(out, p.Output-v)
		fmt.Println("DIFF PAIR", out)
	}

	// TODO

	return 1, nil
}

func (g jsGenome) Mutate(rng *rand.Rand) {
	NewMutator(rng).Mutate(g.ast)
}

func (g jsGenome) Crossover(genome eaopt.Genome, rng *rand.Rand) {
	// Assumes the first body statement is a function
	body := g.ast.Body[0].(*ast.FunctionStatement).Function.Body.(*ast.BlockStatement)
	mine := body.List
	other := genome.(jsGenome).ast.Body[0].(*ast.FunctionStatement).Function.Body.(*ast.BlockStatement).List

	result := make([]ast.Statement, 0)

	var i int
	for i = 0; i < len(mine) && i < len(other); i++ {
		switch rng.Intn(3) {
		case 0: // take statement from mine
			// fmt.Println("Case 0")
			result = append(result, mine[i])
		case 1: // take statement from other
			// fmt.Println("Case 1")
			result = append(result, other[i])
		case 2: // take both statements
			// fmt.Println("Case 2")
			result = append(result, mine[i])
			result = append(result, other[i])
		}
	}

	// In case there are left over statements in mine or other, add them
	j := i
	for j < len(mine) {
		// fmt.Println("leftover fom mine")
		result = append(result, mine[j])
		j++
	}

	j = i
	for j < len(other) {
		// fmt.Println("leftover fom other")
		result = append(result, other[j])
		j++
	}

	body.List = result
}

func (g jsGenome) Clone() eaopt.Genome {
	if g.ast == nil {
		return jsGenome{}
	}

	ast := *g.ast
	return jsGenome{
		name:  g.name,
		ast:   &ast,
		pairs: g.pairs,
	}
}
