package evolutionator

import (
	"math/rand"

	"github.com/MaxHalford/eaopt"
	"github.com/robertkrimen/otto/ast"
)

var _ eaopt.Genome = jsGenome{}

// See https://github.com/MaxHalford/eaopt#implementing-the-genome-interface
type jsGenome struct {
	ast *ast.Program
}

func (g jsGenome) Evaluate() (float64, error) {
	// TODO
	return 1, nil
}

func (g jsGenome) Mutate(rng *rand.Rand) {
	NewMutator(rng).Mutate(g.ast)
}

func (g jsGenome) Crossover(genome eaopt.Genome, rng *rand.Rand) {
	// TODO
}

func (g jsGenome) Clone() eaopt.Genome {
	ast := *g.ast
	return jsGenome{&ast}
}
