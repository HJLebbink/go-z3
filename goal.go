package z3

// #include "go-z3.h"
import "C"

type Goal struct {
	rawCtx   C.Z3_context
	rawGoal C.Z3_goal
}

// Create a new goal
func (c *Context) MkGoal(models, unsat_cores, proofs bool) *Goal {
	rawGoal := C.Z3_mk_goal(c.rawCtx, C.bool(models), C.bool(unsat_cores), C.bool(proofs))
	C.Z3_goal_inc_ref(c.rawCtx, rawGoal)
	return &Goal{
		rawCtx: c.rawCtx,
		rawGoal: rawGoal,
	}
}

// String returns a human-friendly string version of the goal.
func (t *Goal) String() string {
	return C.GoString(C.Z3_goal_to_string(t.rawCtx, t.rawGoal))
}

// Z3_goal_assert
func (g *Goal) Assert(a *AST) {
	C.Z3_goal_assert(g.rawCtx, g.rawGoal, a.rawAST)
}

// Close decreases the reference count for this goal. If nothing else
// has manually increased the reference count, this will free the memory
// associated with it.
func (t *Goal) Close() error {
	C.Z3_goal_dec_ref(t.rawCtx, t.rawGoal)
	return nil
}



// Z3_goal_formula
func (a *Goal) GetFormula(i int) *AST {
	rawAST := C.Z3_goal_formula(a.rawCtx, a.rawGoal, C.uint(i))
	return &AST{
		rawCtx: a.rawCtx,
		rawAST: rawAST,
	}
}

// Return the number of formulas in the given goal.
//
// Maps to: Z3_goal_size
func (g *Goal) GetGoalSize() int {
	return int(C.Z3_goal_size(g.rawCtx, g.rawGoal))
}