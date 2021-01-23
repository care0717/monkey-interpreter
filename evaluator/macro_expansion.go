package evaluator

import (
	"github.com/care0717/monkey-interpreter/ast"
	"github.com/care0717/monkey-interpreter/object"
)

func DefineMacros(program *ast.Program, env object.Environment) {
	var definitions []int

	for i, statement := range program.Statements {
		if isMacroDefinition(statement) {
			addMacro(statement, env)
			definitions = append(definitions, i)
		}
	}

	for i := len(definitions) - 1; i >= 0; i-- {
		index := definitions[i]
		program.Statements = append(
			program.Statements[:index],
			program.Statements[index+1:]...)
	}
}

func isMacroDefinition(node ast.Statement) bool {
	letStatement, ok := node.(*ast.LetStatement)
	if !ok {
		return false
	}

	_, ok = letStatement.Value.(*ast.MacroLiteral)

	if !ok {
		return false
	}

	return true
}

func addMacro(stmt ast.Statement, env object.Environment) {
	letStatement, _ := stmt.(*ast.LetStatement)
	macroLiteral, _ := letStatement.Value.(*ast.MacroLiteral)

	macro := &object.Macro{
		Parameters: macroLiteral.Parameters,
		Body:       macroLiteral.Body,
		Env:        env,
	}
	env.Set(letStatement.Name.Value, macro)
}
