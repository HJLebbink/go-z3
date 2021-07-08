package z3

// #include "go-z3.h"
import "C"


type Params struct {
	rawCtx		C.Z3_context
	rawParams 	C.Z3_params
}


// NewConfig allocates a new configuration object.
func (c *Context) NewParams() *Params {
	rawParams := C.Z3_mk_params(c.rawCtx)
	C.Z3_params_inc_ref(c.rawCtx, rawParams)
	return &Params{
		rawCtx: c.rawCtx,
		rawParams: rawParams,
	}
}


// Convert a parameter set into a string. This function is mainly used for printing the contents of a parameter set.
//
// Maps: Z3_params_to_string
func (p *Params) String() string {
	return C.GoString(C.Z3_params_to_string(p.rawCtx, p.rawParams))
}


// Maps: Z3_params_set_bool
func (p *Params) SetBool(k *Symbol, v bool) {
	C.Z3_params_set_bool(p.rawCtx, p.rawParams, k.rawSymbol, C.bool(v))
}

// Close decreases the reference count for this params. If nothing else
// has manually increased the reference count, this will free the memory
// associated with it.
func (p *Params) Close() error {
	C.Z3_params_dec_ref(p.rawCtx, p.rawParams)
	return nil
}