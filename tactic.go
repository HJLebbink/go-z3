package z3

// #include "go-z3.h"
import "C"

type Tactic struct {
	rawCtx   C.Z3_context
	rawTactic C.Z3_tactic
}


// Return a tactic associated with the given name.
func (c *Context) NewTactic(name string) *Tactic {
	rawTactic := C.Z3_mk_tactic(c.rawCtx, C.CString(name))
	C.Z3_tactic_inc_ref(c.rawCtx, rawTactic)
	return &Tactic{
		rawCtx: c.rawCtx,
		rawTactic: rawTactic,
	}
}

func (c *Context) GetTacticNames() string {
	return "qfaufbv, qfauflia, qfbv, qfidl, qflia, qflra, qfnia, qfnra, qfufbv, qfufbv_ackr, qfufnra, qfuf, ufnia, uflra, auflia, auflira, aufnira, lra, lia, lira, ackermannize_bv, simplify, propagate-values, ctx-simplify"
}

// String returns a human-friendly string version of the tactic.
func (t *Tactic) String() string {
	//return C.GoString(C.Z3_tactic_to_string(t.rawCtx, t.rawTactic))
	return "TODO"
}

func (t *Tactic) With(p *Params) *Tactic {
	return &Tactic{
		rawCtx: t.rawCtx,
		rawTactic: C.Z3_tactic_using_params(t.rawCtx, t.rawTactic, p.rawParams),
	}
}

// Z3_tactic_apply
func (t *Tactic) TacticApply(g *Goal) *ApplyResult {
	rawApplyResult := C.Z3_tactic_apply(t.rawCtx, t.rawTactic, g.rawGoal)
	C.Z3_apply_result_inc_ref(t.rawCtx, rawApplyResult)
	return &ApplyResult{
		rawCtx: t.rawCtx,
		rawApplyResult: rawApplyResult,
	}
}

// Z3_tactic_apply_ex
func (t *Tactic) TacticApplyEx(g *Goal, p *Params) *ApplyResult {
	rawApplyResult := C.Z3_tactic_apply_ex(t.rawCtx, t.rawTactic, g.rawGoal, p.rawParams)
	C.Z3_apply_result_inc_ref(t.rawCtx, rawApplyResult)
	return &ApplyResult{
		rawCtx: t.rawCtx,
		rawApplyResult: rawApplyResult,
	}
}

// Close decreases the reference count for this tactic. If nothing else
// has manually increased the reference count, this will free the memory
// associated with it.
func (t *Tactic) Close() error {
	C.Z3_tactic_dec_ref(t.rawCtx, t.rawTactic)
	return nil
}