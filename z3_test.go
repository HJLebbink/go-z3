package z3

import (
	"fmt"
	"testing"
)

func CreateSolver() (solver *Solver, ctx *Context) {
	config := MkConfig()
	ctx = MkContext(config)
	err := config.Close()
	if err != nil {
		return nil, nil
	}
	//defer ctx.Close()

	solver = ctx.MkSolver()
	//defer solver.Close()
	return
}

func CreateSolverTactic(tactic_name string) (solver *Solver, ctx *Context, tactic *Tactic) {
	config := MkConfig()
	ctx = MkContext(config)
	err := config.Close()
	if err != nil {
		return nil, nil, nil
	}
	//defer ctx.Close()
	tactic = ctx.MkTactic(tactic_name)

	solver = ctx.MkSolverFromTactic(tactic)
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



/*
(declare-fun x () Int)
(assert (>= x 2))
(assert (>= x 3))
(apply (then ctx-solver-simplify propagate-values (par-then (repeat (or-else split-clause skip)) propagate-ineqs)))
yields:
(goals (goal (>= x 3) :precision precise :depth 3))
 */

func Test_Simplify1(t *testing.T) {

	config := MkConfig()
	var ctx *Context = MkContext(config)
	config.Close()

	var ctx_solver_simplify *Tactic = ctx.MkTactic("ctx-solver-simplify")
	var propagate_values *Tactic = ctx.MkTactic("propagate-values")
	var split_clause *Tactic = ctx.MkTactic("split-clause")
	var propagate_ineqs *Tactic = ctx.MkTactic("propagate-ineqs")
	var skip *Tactic = ctx.MkTactic("skip")
	var tactic *Tactic = ctx.AndThen(ctx_solver_simplify, ctx.AndThen(propagate_values, ctx.AndThen(ctx.Repeat(ctx.OrElse(split_clause, skip), 10), propagate_ineqs)))


	var x *AST = ctx.Const(ctx.Symbol("x"), ctx.IntSort())

	//(x > 6) OR (x < 12) -> TRUE
	//(x > 6) AND (x < 12) -> (x > 6) AND (x < 12)
	//(x < 6) OR (x > 12) -> (x < 6) OR (x > 12)
	//(x < 6) AND (x > 12) -> FALSE
	//(x > 6) OR (x < 6) -> x != 6

	var int6 = ctx.Int(6, ctx.IntSort())
	var int12 = ctx.Int(12, ctx.IntSort())

	//(x > 6) OR (x > 12) -> (x > 12)

	var y *AST = x.Gt(int6).Or(x.Gt(int12))
	fmt.Printf("original = %v\n", y.String())

	var goal *Goal = ctx.MkGoal(true, false, false)
	//defer goal.Close()

	goal.Assert(y)

	var r *ApplyResult = tactic.Apply(goal)
	fmt.Printf("ApplyResult = %v\n", r.String())


	var solver = ctx.MkSolverFromTactic(tactic)
	solver.Assert(y)

}




// see https://stackoverflow.com/questions/11507360/t-1-or-t-2-t-1
// see https://stackoverflow.com/questions/38511917/z3-c-api-set-parameter-for-tactic
func Test_single_int_range_simplification(t *testing.T) {

	config := MkConfig()
	var ctx *Context = MkContext(config)
	config.Close()

	ctx.UpdateParamValue("debug_ref_count","true")
	defer ctx.Close()

	var x *AST = ctx.Const(ctx.Symbol("x"), ctx.IntSort())

	//(x > 6) OR (x < 12) -> TRUE
	//(x > 6) AND (x < 12) -> (x > 6) AND (x < 12)
	//(x < 6) OR (x > 12) -> (x < 6) OR (x > 12)
	//(x < 6) AND (x > 12) -> FALSE
	//(x > 6) OR (x < 6) -> x != 6

	var int6 = ctx.Int(6, ctx.IntSort())
	var int12 = ctx.Int(12, ctx.IntSort())

	//(x > 6) OR (x > 12) -> (x > 12)

	var y *AST = x.Gt(int6).Or(x.Gt(int12))
	fmt.Printf("original = %v\n", y.String())

	if false {
		fmt.Printf("%v\n", y.SimplifyGetHelp())
		var params *Params = ctx.MkParams()
		params.SetBool("arith_lhs", true)
		params.SetBool("eq2ineq", false)
		params.SetBool("local_ctx", true)
		params.SetBool("rewrite_patterns", true)
		fmt.Printf("params = %v\n", params.String())
		fmt.Printf("simplify = %v\n", y.SimplifyEx(params).String())
	}

	if true {

		var goal *Goal = ctx.MkGoal(true, false, false)
		//fmt.Printf("goal = %v\n", goal.String())
		//defer goal.Close()

		var params2 *Params = ctx.MkParams()
		//params2.SetBool("arith.propagate_eqs", true)

		//var tactic *Tactic = ctx.MkTactic("simplify")

		//simplify, propagate-values, ctx-simplify

		var tactic *Tactic = ctx.MkTactic("lia").With(params2)
		//fmt.Printf("tactic = %v\n", tactic.String())
		//defer tactic.Close()

		goal.Assert(y)
		//fmt.Printf("goal = %v\n", goal.String())
		//var params3 *Params = ctx.MkParams()
		//var r *ApplyResult = tactic.TacticApplyEx(goal, params3)
		var r *ApplyResult = tactic.Apply(goal)
		fmt.Printf("ApplyResult = %v\n", r.String())
	}
	if false {
		solver := ctx.MkSolverForLogic("LIA") // for logics see: http://smtlib.cs.uiowa.edu/logics.shtml
		solver.Assert(y)
		solver.Check()
		//fmt.Printf("%v\n", solver.String())

		if satisfiable := solver.Check(); satisfiable != True {
			fmt.Printf("unsatisfiable!")
		} else {
			fmt.Printf("%v\n", solver.String())
		}

		//fmt.Printf("%v\n", z.SimplifyGetHelp())

		//(x > 6) AND (x > 12) -> (x > 6)
		//solver.Assert(x.Gt(int6))
		//solver.Assert(x.Gt(int12))
		//var y *AST = x.Simplify()

		//	fmt.Printf("%s\n", y.String())
	}
}
