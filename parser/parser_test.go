package parser

import (
	"testing"

	"github.com/cachesdev/monkey-interpreter/ast"
	"github.com/cachesdev/monkey-interpreter/lexer"
)

// LET
func TestLetStatements(t *testing.T) {
	input := `
  let x = 5;
  let y = 10;
  let foobar = 838383;
  `

	l := lexer.New(input)
	p := New(l)

	program := p.ParseProgram()
	checkParseErrors(t, p)
	if program == nil {
		t.Fatalf("ParseProgram() returned nill")
	}
	if len(program.Statements) != 3 {
		t.Fatalf("program.Statements no contiene 3 sentencias. recibido=%d", len(program.Statements))
	}

	tests := []struct {
		expectedIdentifier string
	}{
		{"x"},
		{"y"},
		{"foobar"},
	}

	for i, tt := range tests {
		stmt := program.Statements[i]
		if !testLetStatement(t, stmt, tt.expectedIdentifier) {
			return
		}
	}
}

func testLetStatement(t *testing.T, s ast.Statement, name string) bool {
	if s.TokenLiteral() != "let" {
		t.Errorf("s.TokenLiteral no es 'let'. recibido=%q", s.TokenLiteral())
		return false
	}

	letStmt, ok := s.(*ast.LetStatement)
	if !ok {
		t.Errorf("s no es *ast.LetStatement. recibido=%T", s)
		return false
	}

	if letStmt.Name.Value != name {
		t.Errorf("letStmt.Name.Value no es '%s'. recibido=%s", name, letStmt.Name.Value)
		return false
	}

	if letStmt.Name.TokenLiteral() != name {
		t.Errorf("letStmt.Name.TokenLiteral() no es '%s'. recibido=%s", name, letStmt.Name.TokenLiteral())
		return false
	}

	return true
}

// RETURN
func TestReturnStatements(t *testing.T) {
	input := `
  return 5;
  return 10;
  return 993322;
  `

	l := lexer.New(input)
	p := New(l) // program.New

	program := p.ParseProgram()
	checkParseErrors(t, p)

	if len(program.Statements) != 3 {
		t.Fatalf("program.Statements no contiene 3 sentencias. recibido=%d", len(program.Statements))
	}

	for _, stmt := range program.Statements {
		returnStmt, ok := stmt.(*ast.ReturnStatement)
		if !ok {
			t.Errorf("stmt no es *ast.ReturnStatement. recibido=%T", stmt)
			continue
		}

		if returnStmt.TokenLiteral() != "return" {
			t.Errorf("returnStmt.TokenLiteral no es 'return', recibido=%q", returnStmt.TokenLiteral())
		}
	}
}

func TestIdentifierExpression(t *testing.T) {
	input := `
  foobar;
  `

	l := lexer.New(input)
	p := New(l)
	program := p.ParseProgram()
	checkParseErrors(t, p)

	if len(program.Statements) != 1 {
		t.Fatalf("El programa no tiene setencias suficientes. recibido=%d", len(program.Statements))
	}

	stmt, ok := program.Statements[0].(*ast.ExpressionStatement)
	if !ok {
		t.Fatalf("program.Statements[0] no es ast.ExpressionStatement. recibido=%T", program.Statements[0])
	}

	ident, ok := stmt.Expression.(*ast.Identifier)
	if !ok {
		t.Fatalf("exp not *ast.Identifier. recibido=%T", stmt.Expression)
	}
	if ident.Value != "foobar" {
		t.Errorf("ident.Value no es %s. recibido=%s", "foobar", ident.Value)
	}
	if ident.TokenLiteral() != "foobar" {
		t.Errorf("ident.TokenLiteral no es %s. recibido=%s", "foobar", ident.TokenLiteral())
	}
}

func TestIntegerLiteralExpression(t *testing.T) {
	input := "5;"

	l := lexer.New(input)
	p := New(l)
	program := p.ParseProgram()
	checkParseErrors(t, p)

	if len(program.Statements) != 1 {
		t.Fatalf("El programa no tiene setencias suficientes. recibido=%d", len(program.Statements))
	}

	stmt, ok := program.Statements[0].(*ast.ExpressionStatement)
	if !ok {
		t.Fatalf("program.Statements[0] no es ast.ExpressionStatement. recibido=%d", len(program.Statements))
	}

	literal, ok := stmt.Expression.(*ast.IntegerLiteral)
	if !ok {
		t.Fatalf("exp no es *ast.IntegerLiteral. recibido=%T", stmt.Expression)
	}
	if literal.Value != 5 {
		t.Errorf("literal.Value no es %d. recibido=%d", 5, literal.Value)
	}
	if literal.TokenLiteral() != "5" {
		t.Errorf("literal.TokenLiteral no es %s. recibido=%s", "5", literal.TokenLiteral())
	}
}

func checkParseErrors(t *testing.T, p *Parser) {
	errors := p.Errors()
	if len(errors) == 0 {
		return
	}

	t.Errorf("parser tiene %d errores", len(errors))
	for _, msg := range errors {
		t.Errorf("parser error: %q", msg)
	}
	t.FailNow()
}
