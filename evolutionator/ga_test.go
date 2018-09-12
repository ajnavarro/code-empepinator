package evolutionator

import (
	"testing"

	"github.com/robertkrimen/otto/ast"
	"github.com/sanity-io/litter"
	"github.com/stretchr/testify/assert"
)

func TestEvolutionator(t *testing.T) {
	var ast *ast.Program

	res, err := Optimize(ast)

	assert.NoError(t, err)
	litter.Dump(res)
}
