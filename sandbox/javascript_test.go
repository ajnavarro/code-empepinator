package sandbox

import (
	"math"
	"testing"

	"github.com/sanity-io/litter"

	"github.com/stretchr/testify/require"
)

func TestAST(t *testing.T) {
	require := require.New(t)

	content := `
function multiply(a,b) {
	return a*b;
}`
	executor := NewJavascript("multiply")

	result, err := executor.Execute(content, 2, 2)
	require.NoError(err)
	require.Equal(float64(4), result)

	litter.Dump(executor.AST)

	result, err = executor.ExecuteAST(executor.AST, 3, 2)
	require.NoError(err)
	require.Equal(float64(6), result)

}

func TestExecutor(t *testing.T) {
	require := require.New(t)

	content := `
function multiply(a,b) {
	return a*b;
}`
	executor := NewJavascript("multiply")

	result, err := executor.Execute(content, 2, 2)

	require.NoError(err)
	require.Equal(float64(4), result)
}

func TestExecutorBadParameters(t *testing.T) {
	require := require.New(t)

	content := `
function multiply(a,b) {
	return a*b;
}`
	executor := NewJavascript("multiply")

	result, err := executor.Execute(content, 2)
	require.NoError(err)
	require.True(math.IsNaN(result))
}

func TestExecutorBadCode(t *testing.T) {
	require := require.New(t)

	content := `
function multiply(a,b) {
	BADCODE
}`
	executor := NewJavascript("multiply")

	_, err := executor.Execute(content, 2)
	require.Error(err)
}
