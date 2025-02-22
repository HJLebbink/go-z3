package z3

import (
	"testing"
)

func TestSolver(t *testing.T) {
	config := MkConfig()
	defer config.Close()
	ctx := MkContext(config)
	defer ctx.Close()

	// Create the "x xor y" constraint
	boolTyp := ctx.BoolSort()
	x := ctx.Const(ctx.Symbol("x"), boolTyp)
	y := ctx.Const(ctx.Symbol("y"), boolTyp)
	x_xor_y := x.Xor(y)
	ast := x_xor_y
	t.Logf("\nAST:\n%s", ast.String())

	// Create the solver
	s := ctx.MkSolver()
	defer s.Close()

	// Assert constraints
	s.Assert(x_xor_y)

	// Solve
	result := s.Check()
	if result != True {
		t.Fatalf("bad: %v", result)
	}

	// Get the model
	m := s.Model()
	defer m.Close()
	t.Logf("\nModel:\n%s", m.String())
}

func TestRealSolver(t *testing.T) {
	config := MkConfig()
	defer config.Close()

	ctx := MkContext(config)
	defer ctx.Close()

	x := ctx.Const(ctx.Symbol("x"), ctx.RealSort())
	y := ctx.Const(ctx.Symbol("y"), ctx.RealSort())
	ast := x.Div(y).Eq(ctx.Real(1, 1, ctx.RealSort()))
	t.Logf("\nAST:\n%s", ast.String())

	// Create the solver
	s := ctx.MkSolver()
	defer s.Close()

	// Assert constraints
	s.Assert(ast)

	// Solve
	result := s.Check()
	if result != True {
		t.Fatalf("bad: %v", result)
	}

}