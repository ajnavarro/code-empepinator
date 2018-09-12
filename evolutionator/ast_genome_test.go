package evolutionator

import (
	"math/rand"
	"testing"

	"github.com/robertkrimen/otto/ast"
	"github.com/stretchr/testify/assert"

	"github.com/ajnavarro/code-empepinator/sandbox"
	"github.com/sanity-io/litter"
)

var jscode = `
function multiply(p1,p2) {
    var a = p1;
    var b = p2;
    var result = p1 *p2;

    return result;
}`

var jscode2 = `
function multiply(p1,p2) {
    var a = p1;
		var b = p2;
		var c = 1;
    var result = p1 * p2 * c;

    return result;
}`

func parse(t *testing.T, code string) jsGenome {
	executor := sandbox.NewJavascript("multiply")
	err := executor.Parse(code)
	assert.NoError(t, err)

	return jsGenome{ast: executor.AST}
}

func TestASTGenome(t *testing.T) {
	g := parse(t, jscode)
	litter.Dump(g.ast)
}

func TestCrossover(t *testing.T) {
	g := parse(t, jscode)
	g2 := parse(t, jscode2)

	r := rand.New(rand.NewSource(3))
	g.Crossover(g2, r)
	//litter.Dump(g.ast.Body)
	body := g.ast.Body[0].(*ast.FunctionStatement).Function.Body.(*ast.BlockStatement)
	assert.Equal(t, 6, len(body.List))
}

func TestMutate(t *testing.T) {
	g := parse(t, jscode)

	s1 := litter.Sdump(g.ast)

	r := rand.New(rand.NewSource(3))
	g.Mutate(r)

	s2 := litter.Sdump(g.ast)

	assert.NotEqual(t, s1, s2)
}
