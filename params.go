package z3

// #include "go-z3.h"
import "C"


type Params struct {
	rawCtx		C.Z3_context
	rawParams 	C.Z3_params
}


// NewConfig allocates a new configuration object.
func (c *Context) NewParams() *Params {
	return &Params{
		rawCtx: c.raw,
		rawParams: C.Z3_mk_params(c.raw),
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

//-------------------------------------------------------------------
// Memory Management
//-------------------------------------------------------------------

// Close decreases the reference count for this params. If nothing else
// has manually increased the reference count, this will free the memory
// associated with it.
func (p *Params) Close() error {
	C.Z3_params_dec_ref(p.rawCtx, p.rawParams)
	return nil
}

// IncRef increases the reference count of this model. This is advanced,
// you probably don't need to use this.
func (p *Params) IncRef() {
	C.Z3_params_inc_ref(p.rawCtx, p.rawParams)
}

// DecRef decreases the reference count of this model. This is advanced,
// you probably don't need to use this.
//
// Close will decrease it automatically from the initial 1, so this should
// only be called with exact matching calls to IncRef.
func (p *Params) DecRef() {
	C.Z3_params_dec_ref(p.rawCtx, p.rawParams)
}
