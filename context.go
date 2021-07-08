package z3

// #include "go-z3.h"
import "C"

// Context is what handles most of the interactions with Z3.
type Context struct {
	rawCtx C.Z3_context
}

func MkContext(c *Config) *Context {
	return &Context{
		rawCtx: C.Z3_mk_context(c.Z3Value()),
	}
}

/*
Legal parameters are:
  auto_config (bool) (default: true)
  debug_ref_count (bool) (default: false)
  dot_proof_file (string) (default: proof.dot)
  dump_models (bool) (default: false)
  model (bool) (default: true)
  model_validate (bool) (default: false)
  proof (bool) (default: false)
  rlimit (unsigned int) (default: 0)
  smtlib2_compliant (bool) (default: false)
  stats (bool) (default: false)
  timeout (unsigned int) (default: 4294967295)
  trace (bool) (default: false)
  trace_file_name (string) (default: z3.log)
  type_check (bool) (default: true)
  unicode (bool)
  unsat_core (bool) (default: false)
  well_sorted_check (bool) (default: false)
*/
// maps to: Z3_update_param_value
func (c *Context) UpdateParamValue(id, value string) {
	C.Z3_update_param_value(c.rawCtx, C.CString(id), C.CString(value))
}

// Close frees the memory associated with this context.
func (c *Context) Close() error {
	// Clear context
	C.Z3_del_context(c.rawCtx)

	// Clear error handling
	errorHandlerMapLock.Lock()
	delete(errorHandlerMap, c.rawCtx)
	errorHandlerMapLock.Unlock()

	return nil
}

//Legal parameters are:
//  max_depth (unsigned int)
//  max_memory (unsigned int)
//  max_steps (unsigned int)
//  propagate_eq (bool)
//
// Maps to: Z3_tactic_using_params
func (t *Tactic) With(p *Params) *Tactic {
	rawTactic := C.Z3_tactic_using_params(t.rawCtx, t.rawTactic, p.rawParams)
	C.Z3_tactic_inc_ref(t.rawCtx, rawTactic)
	return &Tactic{
		rawCtx: t.rawCtx,
		rawTactic: rawTactic,
	}
}

func (c *Context) AndThen(t1, t2 *Tactic) *Tactic {
	rawTactic := C.Z3_tactic_and_then(c.rawCtx, t1.rawTactic, t2.rawTactic)
	C.Z3_tactic_inc_ref(c.rawCtx, rawTactic)
	return &Tactic{
		rawCtx: c.rawCtx,
		rawTactic: rawTactic,
	}
}
func (c *Context) ParAndThen(t1, t2 *Tactic) *Tactic {
	rawTactic := C.Z3_tactic_par_and_then(c.rawCtx, t1.rawTactic, t2.rawTactic)
	C.Z3_tactic_inc_ref(c.rawCtx, rawTactic)
	return &Tactic{
		rawCtx: c.rawCtx,
		rawTactic: rawTactic,
	}

}
func (c *Context) OrElse(t1, t2 *Tactic) *Tactic {
	rawTactic := C.Z3_tactic_or_else(c.rawCtx, t1.rawTactic, t2.rawTactic)
	C.Z3_tactic_inc_ref(c.rawCtx, rawTactic)
	return &Tactic{
		rawCtx: c.rawCtx,
		rawTactic: rawTactic,
	}

}
func (c *Context) Repeat(t *Tactic, max uint) *Tactic {
	rawTactic := C.Z3_tactic_repeat(c.rawCtx, t.rawTactic, C.uint(max))
	C.Z3_tactic_inc_ref(c.rawCtx, rawTactic)
	return &Tactic{
		rawCtx: c.rawCtx,
		rawTactic: rawTactic,
	}

}