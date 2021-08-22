package evaluator

import (
	"testing"

	"github.com/penthious/writing_an_interpreter_in_go/ast"
	"github.com/penthious/writing_an_interpreter_in_go/lexer"
	"github.com/penthious/writing_an_interpreter_in_go/object"
	"github.com/penthious/writing_an_interpreter_in_go/parser"
)

func TestDefineMacros(t *testing.T) {
	input := `
let number = 1;
let function = fn(x, y) { x + y; };
let mymacro = macro(x, y) { x + y; };
`

	env := object.NewEnvironment()
	program := testParseProgram(input)

	DefineMacros(program, env)

	if len(program.Statements) != 2 {
		t.Fatalf("program.Statements want=2, got=%d", len(program.Statements))
	}

	_, ok := env.Get("number")
	if ok {
		t.Fatalf("number should not be defined")
	}

	_, ok = env.Get("function")
	if ok {
		t.Fatalf("function should not be defined")
	}
	obj, ok := env.Get("mymacro")
	if !ok {
		t.Fatalf("macro not in environment")
	}

	macro, ok := obj.(*object.Macro)
	if !ok {
		t.Fatalf("object is not Macro, got=%T (%+v)", obj, obj)
	}

	if len(macro.Parameters) != 2 {
		t.Fatalf("macro.Parameters want=2, got=%d", len(macro.Parameters))
	}

	if macro.Parameters[0].String() != "x" {
		t.Fatalf("macro.Parameters[0] want=x, got=%s", macro.Parameters[0])
	}

	if macro.Parameters[1].String() != "y" {
		t.Fatalf("macro.Parameters[0] want=y, got=%s", macro.Parameters[1])
	}

	wantBody := "(x + y)"

	if macro.Body.String() != wantBody {
		t.Fatalf("macro.Body want=%s, got=%s", wantBody, macro.Parameters[1])
	}

}

func TestExpandMacro(t *testing.T) {
	tests := []struct {
		input string
		want  string
	}{
		{
			`
			let infixExpression = macro() { quote(1 + 2); };

			infixExpression();
			`,
			`(1 + 2)`,
		},
		{
			`
			let reverse = macro(a, b) { quote(unquote(b) - unquote(a)); };

			reverse(2 + 2, 10 - 5);
			`,
			`(10 - 5) - (2 + 2)`,
		},
		{
			`
			let unless = macro(condition, consequence, alternative) {
				quote(if (!(unquote(condition))) {
					unquote(consequence);
				} else {
					unquote(alternative);
				});
			};

			unless(10 > 5, puts("not greater"), puts("greater"));
			`,
			`if (!(10 > 5)) { puts("not greater") } else { puts("greater") }`,
		},
	}

	for _, tt := range tests {
		want := testParseProgram(tt.want)
		program := testParseProgram(tt.input)

		env := object.NewEnvironment()
		DefineMacros(program, env)
		expanded := ExpandMacros(program, env)

		if expanded.String() != want.String() {
			t.Errorf("not equal, want=%q, got=%q", expanded.String(), want.String())
		}
	}
}

func testParseProgram(input string) *ast.Program {
	l := lexer.New(input)
	p := parser.New(l)
	return p.ParseProgram()
}
