package evolutionator

import (
	"math/rand"
	"testing"

	"github.com/sanity-io/litter"

	"github.com/ajnavarro/code-empepinator/sandbox"
	"github.com/stretchr/testify/require"
)

func TestAST(t *testing.T) {
	require := require.New(t)
	m := NewMutator(rand.New(rand.NewSource(42)))

	content := `function multiply(p1,p2) {
		var a = p1;
		var b = p2;
		var result = p1 *p2;

		return result;
	}`
	executor := sandbox.NewJavascript("multiply")

	result, err := executor.Execute(content, 2, 2)
	require.NoError(err)
	require.Equal(float64(4), result)

	prevAst := *executor.AST
	ast := m.Mutate(executor.AST)

	require.NotEqual(prevAst, ast)
	litter.Dump(ast)
}
