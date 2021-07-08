package z3

// #include "go-z3.h"
import "C"

// Sort represents a sort in Z3.
type Sort struct {
	rawCtx  C.Z3_context
	rawSort C.Z3_sort
}

// BoolSort returns the boolean type.
func (c *Context) BoolSort() *Sort {
	return &Sort{
		rawCtx:  c.rawCtx,
		rawSort: C.Z3_mk_bool_sort(c.rawCtx),
	}
}

// IntSort returns the int type.
func (c *Context) IntSort() *Sort {
	return &Sort{
		rawCtx:  c.rawCtx,
		rawSort: C.Z3_mk_int_sort(c.rawCtx),
	}
}

// RealSort returns the boolean type.
func (c *Context) RealSort() *Sort {
	return &Sort{
		rawCtx:  c.rawCtx,
		rawSort: C.Z3_mk_real_sort(c.rawCtx),
	}
}

// StringSort returns the string type.
func (c *Context) StringSort() *Sort {
	return &Sort{
		rawCtx:  c.rawCtx,
		rawSort: C.Z3_mk_string_sort(c.rawCtx),
	}
}

// SeqSort returns the seq type.
func (c *Context) SeqSort(sort Sort) *Sort{
	return &Sort{
		rawCtx:  c.rawCtx,
		rawSort: C.Z3_mk_seq_sort(c.rawCtx, sort.rawSort),
	}
}