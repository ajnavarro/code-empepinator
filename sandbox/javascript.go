package sandbox

import (
	"strconv"
	"strings"

	"github.com/robertkrimen/otto"
	"github.com/robertkrimen/otto/ast"
	"github.com/robertkrimen/otto/parser"
)

type Javascript struct {
	name string
	vm   *otto.Otto
	AST  *ast.Program
}

func NewJavascript(name string) *Javascript {
	return &Javascript{
		vm:   otto.New(),
		name: name,
	}
}

func (e *Javascript) Parse(function string) error {
	ast, err := parser.ParseFile(nil, "", function, 0)
	if err != nil {
		return err
	}

	e.AST = ast
	return nil
}

func (e *Javascript) Execute(function string, vals ...float64) (float64, error) {
	_, err := e.vm.Run(function)
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

	if err := e.Parse(function); err != nil {
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
