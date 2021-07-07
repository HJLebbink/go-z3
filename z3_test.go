package z3

import (
	"fmt"
	"testing"
)

func CreateSolver() (solver *Solver, ctx *Context) {
	config := NewConfig()
	ctx = NewContext(config)
	err := config.Close()
	if err != nil {
		return nil, nil
	}
	//defer ctx.Close()

	solver = ctx.NewSolver()
	//defer solver.Close()
	return
}

// This example is a basic mathematical example
func Test_BasicMath(t *testing.T) {
	var s, ctx = CreateSolver()

	// Logic:
	// x + y + z > 4
	// x + y < 2
	// z > 0
	// x != y != z
	// x, y, z != 0
	// x + y = -3

	// Vars
	x := ctx.Const(ctx.Symbol("x"), ctx.IntSort())
	y := ctx.Const(ctx.Symbol("y"), ctx.IntSort())
	z := ctx.Const(ctx.Symbol("z"), ctx.IntSort())
	{
		zero := ctx.Int(0, ctx.IntSort()) // To save repeats

		// x + y + z > 4
		s.Assert(x.Add(y, z).Gt(ctx.Int(4, ctx.IntSort())))

		// x + y < 2
		s.Assert(x.Add(y).Lt(ctx.Int(2, ctx.IntSort())))

		// z > 0
		s.Assert(z.Gt(zero))

		// x != y != z
		s.Assert(x.Distinct(y, z))

		// x, y, z != 0
		s.Assert(x.Eq(zero).Not())
		s.Assert(y.Eq(zero).Not())
		s.Assert(z.Eq(zero).Not())

		// x + y = -3
		s.Assert(x.Add(y).Eq(ctx.Int(-3, ctx.IntSort())))

		if v := s.Check(); v != True {
			fmt.Println("Unsolveable")
			return
		}
	}

	// Get the resulting model:
	m := s.Model()
	assignments := m.Assignments()
	m.Close()

	{
		var xs = assignments["x"].String()
		var ys = assignments["y"].String()
		var zs = assignments["z"].String()

		if xs != "(- 2)" {
			t.Fatalf("")
		}
		if ys != "(- 1)" {
			t.Fatalf("")
		}
		if zs != "8" {
			t.Fatalf("")
		}
		// Output:
		// x = (- 2)
		// y = (- 1)
		// z = 8
		if false {
			fmt.Printf("x = %s\n", xs)
			fmt.Printf("y = %s\n", ys)
			fmt.Printf("z = %s\n", zs)
		}
	}
}

// From C examples: demorgan
func Test_Demorgan(t *testing.T) {
	// Create the context
	var s, ctx = CreateSolver()

	// Create a couple variables
	x := ctx.Const(ctx.Symbol("x"), ctx.BoolSort())
	y := ctx.Const(ctx.Symbol("y"), ctx.BoolSort())

	// Final goal: !(x && y) == (!x || !y)
	// Built incrementally so its clearer

	// !(x && y)
	not_x_and_y := x.And(y).Not()

	// (!x || !y)
	not_x_or_not_y := x.Not().Or(y.Not())

	// Conjecture and negated
	conj := not_x_and_y.Iff(not_x_or_not_y)
	negConj := conj.Not()

	// Assert the constraints
	s.Assert(negConj)

	var v LBool = s.Check()

	if v != False {
		t.Fatalf("")
	}
	// Output:
	// DeMorgan is valid
}

// From C examples: find_model_example2
func Test_FindModel2(t *testing.T) {
	var s, ctx = CreateSolver()

	// Create a couple variables
	x := ctx.Const(ctx.Symbol("x"), ctx.IntSort())
	y := ctx.Const(ctx.Symbol("y"), ctx.IntSort())

	// Create a couple integers
	v1 := ctx.Int(1, ctx.IntSort())
	v2 := ctx.Int(2, ctx.IntSort())

	// y + 1
	y_plus_one := y.Add(v1)

	// x < y + 1 && x > 2
	c1 := x.Lt(y_plus_one)
	c2 := x.Gt(v2)

	// Assert the constraints
	s.Assert(c1)
	s.Assert(c2)

	{
		if v := s.Check(); v != True {
			t.Logf("unsatisfied!")
			return
		}

		// Get the resulting model:
		m := s.Model()
		assignments := m.Assignments()
		m.Close()

		var xs = assignments["x"].Int()
		var ys = assignments["y"].Int()

		if xs != 3 {
			t.Fatalf("")
		}
		if ys != 3 {
			t.Fatalf("")
		}
	}

	// Create some new assertions
	//
	// !(x == y)
	c3 := x.Eq(y).Not()
	s.Assert(c3)

	{
		// Solve
		if v := s.Check(); v != True {
			t.Logf("unsatisfied!")
			return
		}

		// Get the resulting model:
		m := s.Model()
		assignments := m.Assignments()
		m.Close()

		var xs = assignments["x"].Int()
		var ys = assignments["y"].Int()

		if xs != 3 {
			t.Fatalf("")
		}
		if ys != 4 {
			t.Fatalf("")
		}
	}
}

// see https://stackoverflow.com/questions/11507360/t-1-or-t-2-t-1
func Test_single_int_range_simplification(t *testing.T) {
	var solver, ctx = CreateSolver()
	var x *AST = ctx.Const(ctx.Symbol("x"), ctx.IntSort())


	//(x > 6) OR (x < 12) -> TRUE

	//(x > 6) AND (x < 12) -> (x > 6) AND (x < 12)

	//(x < 6) OR (x > 12) -> (x < 6) OR (x > 12)

	//(x < 6) AND (x > 12) -> FALSE

	//(x > 6) OR (x < 6) -> x != 6


	var int6 = ctx.Int(6, ctx.IntSort())
	var int12 = ctx.Int(12, ctx.IntSort())

	//(x > 6) OR (x > 12) -> (x > 12)
	//solver.Assert(x.Or(x.Gt(int6), x.Gt(int12)))

	//(x > 6) AND (x > 12) -> (x > 6)
	solver.Assert(x.Gt(int6))
	solver.Assert(x.Gt(int12))
	//var y *AST = x.Simplify()

//	s

//	fmt.Printf("%s\n", y.String())
}
