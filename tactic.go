package z3

// #include "go-z3.h"
import "C"

type Tactic struct {
	rawCtx   C.Z3_context
	rawTactic C.Z3_tactic
}


func (c *Context) NewTactic() *Tactic {
	return &Tactic{
		rawCtx: c.raw,
		rawTactic: C.Z3_mk_tactic(c.raw),
	}
}

// String returns a human-friendly string version of the model.
func (t *Tactic) String() string {
	return C.GoString(C.Z3_tactic_to_string(t.rawCtx, t.rawTactic))
}

//-------------------------------------------------------------------
// Memory Management
//-------------------------------------------------------------------

// Close decreases the reference count for this tactic. If nothing else
// has manually increased the reference count, this will free the memory
// associated with it.
func (t *Tactic) Close() error {
	C.Z3_tactic_dec_ref(t.rawCtx, t.rawTactic)
	return nil
}

// IncRef increases the reference count of this tactic. This is advanced,
// you probably don't need to use this.
func (t *Tactic) IncRef() {
	C.Z3_model_inc_ref(t.rawCtx, t.rawTactic)
}

// DecRef decreases the reference count of this tactic. This is advanced,
// you probably don't need to use this.
//
// Close will decrease it automatically from the initial 1, so this should
// only be called with exact matching calls to IncRef.
func (t *Tactic) DecRef() {
	C.Z3_model_dec_ref(t.rawCtx, t.rawTactic)
}
