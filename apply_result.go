package z3

// #include "go-z3.h"
import "C"

type ApplyResult struct {
	rawCtx C.Z3_context
	rawApplyResult C.Z3_apply_result
}



// String returns a human-friendly string version of the apply_result.
func (t *ApplyResult) String() string {
	return C.GoString(C.Z3_apply_result_to_string(t.rawCtx, t.rawApplyResult))
}

// Close decreases the reference count for this tactic. If nothing else
// has manually increased the reference count, this will free the memory
// associated with it.
func (a *ApplyResult) Close() error {
	C.Z3_apply_result_dec_ref(a.rawCtx, a.rawApplyResult)
	return nil
}