package sandbox

import (
	"strconv"
	"strings"

	"github.com/robertkrimen/otto"
	"github.com/robertkrimen/otto/ast"
	"github.com/robertkrimen/otto/parser"
)

type Javascript struct {
	name, function string

	vm *otto.Otto

	AST *ast.Program
}

func NewJavascript(name, function string) *Javascript {
	return &Javascript{
		vm:       otto.New(),
		name:     name,
		function: function,
	}
}

func (e *Javascript) Parse() error {
	ast, err := parser.ParseFile(nil, "", e.function, 0)
	if err != nil {
		return err
	}

	e.AST = ast
	return nil
}

func (e *Javascript) Execute(vals ...float64) (float64, error) {
	_, err := e.vm.Run(e.function)
	if err != nil {
		return 0, err
	}

	var strVals []string
	for _, v := range vals {
		strVals = append(strVals, strconv.FormatFloat(v, 'f', -1, 64))
	}

	params := strings.Join(strVals, ",")

	val, err := e.vm.Run(e.name + "(" + params + ")")
	if err != nil {
		return 0, err
	}

	err = e.Parse()
	if err != nil {
		return 0, err
	}

	return val.ToFloat()
}

func (e *Javascript) ExecuteAST(ast *ast.Program, vals ...float64) (float64, error) {
	_, err := e.vm.Run(ast)
	if err != nil {
		return 0, err
	}

	var strVals []string
	for _, v := range vals {
		strVals = append(strVals, strconv.FormatFloat(v, 'f', -1, 64))
	}

	params := strings.Join(strVals, ",")
	val, err := e.vm.Run(e.name + "(" + params + ")")
	if err != nil {
		return 0, err
	}

	return val.ToFloat()
}
