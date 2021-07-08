package z3

// #include "go-z3.h"
import "C"

// Solver is a single solver tied to a specific Context within Z3.
//
// It is created via the MkSolver methods on Context. When a solver is
// no longer needed, the Close method must be called. This will remove the
// solver from the context and no more APIs on Solver may be called
// thereafter.
//
// Freeing the context (Context.Close) will NOT automatically close associated
// solvers. They must be managed separately.
type Solver struct {
	rawCtx    C.Z3_context
	rawSolver C.Z3_solver
}

// MkSolver creates a new solver.
func (c *Context) MkSolver() *Solver {
	rawSolver := C.Z3_mk_solver(c.rawCtx)
	C.Z3_solver_inc_ref(c.rawCtx, rawSolver)
	return &Solver{
		rawCtx:    c.rawCtx,
		rawSolver: rawSolver,
	}
}

// MkSolver creates a new solver for the provided logic
// for logics see: http://smtlib.cs.uiowa.edu/logics.shtml
//
// Maps to: Z3_mk_solver_for_logic
func (c *Context) MkSolverForLogic(name string) *Solver {
	rawSolver := C.Z3_mk_solver_for_logic(c.rawCtx, C.Z3_mk_string_symbol(c.rawCtx, C.CString(name)))
	C.Z3_solver_inc_ref(c.rawCtx, rawSolver)
	return &Solver{
		rawCtx:    c.rawCtx,
		rawSolver: rawSolver,
	}
}

// Create a new solver that is implemented using the given tactic.
//
// Maps to: Z3_mk_solver_from_tactic
func (c *Context) MkSolverFromTactic(t *Tactic) *Solver {
	rawSolver := C.Z3_mk_solver_from_tactic(c.rawCtx, t.rawTactic)
	C.Z3_solver_inc_ref(c.rawCtx, rawSolver)
	return &Solver{
		rawCtx:    c.rawCtx,
		rawSolver: rawSolver,
	}
}

// Close frees the memory associated with this.
func (s *Solver) Close() error {
	C.Z3_solver_dec_ref(s.rawCtx, s.rawSolver)
	return nil
}

// Remove all assertions from the solver.
//
// Maps to: Z3_solver_reset
func (s *Solver) Reset()  {
	C.Z3_solver_reset(s.rawCtx, s.rawSolver)
}

// Convert a solver into a string.
//
// Maps to: Z3_solver_to_string
func (s *Solver) String() string {
	return C.GoString(C.Z3_solver_to_string(s.rawCtx, s.rawSolver))
}

// Maps to: Z3_solver_set_params
func (s *Solver) SetParams(p *Params) {
	C.Z3_solver_set_params(s.rawCtx, s.rawSolver, p.rawParams)
}


// Assert asserts a constraint onto the Solver.
//
// Maps to: Z3_solver_assert
func (s *Solver) Assert(a *AST) {
	C.Z3_solver_assert(s.rawCtx, s.rawSolver, a.rawAST)
}

// Check checks if the currently set formula is consistent.
//
// Maps to: Z3_solver_check
func (s *Solver) Check() LBool {
	return LBool(C.Z3_solver_check(s.rawCtx, s.rawSolver))
}

// Model returns the last model from a Check.
//
// Maps to: Z3_solver_get_model
func (s *Solver) Model() *Model {
	m := &Model{
		rawCtx:   s.rawCtx,
		rawModel: C.Z3_solver_get_model(s.rawCtx, s.rawSolver),
	}
	C.Z3_model_inc_ref(m.rawCtx, m.rawModel)
	return m
}