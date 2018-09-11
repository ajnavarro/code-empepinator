package sandbox

import (
	"math"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestExecutor(t *testing.T) {
	require := require.New(t)

	content := `
function multiply(a,b) {
	return a*b;
}`
	executor := NewJavascript("multiply", content)

	result, err := executor.Execute(2, 2)

	require.NoError(err)
	require.Equal(float64(4), result)
}

func TestExecutorBadParameters(t *testing.T) {
	require := require.New(t)

	content := `
function multiply(a,b) {
	return a*b;
}`
	executor := NewJavascript("multiply", content)

	result, err := executor.Execute(2)
	require.NoError(err)
	require.True(math.IsNaN(result))
}

func TestExecutorBadCode(t *testing.T) {
	require := require.New(t)

	content := `
function multiply(a,b) {
	BADCODE
}`
	executor := NewJavascript("multiply", content)

	_, err := executor.Execute(2)
	require.Error(err)
}
