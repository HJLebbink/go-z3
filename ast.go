package z3

// #include <stdlib.h>
// #include "go-z3.h"
import "C"

// AST represents an AST value in Z3.
//
// AST memory management is automatically managed by the Context it
// is contained within. When the Context is freed, so are the AST nodes.
type AST struct {
	rawCtx C.Z3_context
	rawAST C.Z3_ast
}

// String returns a human-friendly string version of the AST.
func (a *AST) String() string {
	return C.GoString(C.Z3_ast_to_string(a.rawCtx, a.rawAST))
}

// DeclName returns the name of a declaration. The AST value must be a
// func declaration for this to work.
func (a *AST) DeclName() *Symbol {
	return &Symbol{
		rawCtx: a.rawCtx,
		rawSymbol: C.Z3_get_decl_name(
			a.rawCtx, C.Z3_to_func_decl(a.rawCtx, a.rawAST)),
	}
}

//-------------------------------------------------------------------
// Var, Literal Creation
//-------------------------------------------------------------------

// Const declares a variable. It is called "Const" since internally
// this is equivalent to create a function that always returns a constant
// value. From an initial user perspective this may be confusing but go-z3
// is following identical naming convention.
func (c *Context) Const(s *Symbol, typ *Sort) *AST {
	return &AST{
		rawCtx: c.rawCtx,
		rawAST: C.Z3_mk_const(c.rawCtx, s.rawSymbol, typ.rawSort),
	}
}

// Int creates an integer type.
//
// Maps: Z3_mk_int
func (c *Context) Int(v int, typ *Sort) *AST {
	return &AST{
		rawCtx: c.rawCtx,
		rawAST: C.Z3_mk_int(c.rawCtx, C.int(v), typ.rawSort),
	}
}

// Real creates a real type.
//
// Maps: Z3_mk_real
func (c *Context) Real(num int, den int, typ *Sort) *AST {
	return &AST{
		rawCtx: c.rawCtx,
		rawAST: C.Z3_mk_real(c.rawCtx, C.int(num), C.int(den)),
	}
}

// Float creates an float type.
//
// Maps: Z3_mk_real
func (c *Context) Float(v float64) *AST {
	//TODO: test if this could work
	return &AST{
		rawCtx: c.rawCtx,
		rawAST: C.Z3_mk_real(c.rawCtx, C.int(v), C.int(1)),
	}
}

// Str creates an string type.
//
// Maps: Z3_mk_string
func (c *Context) Str(str string) *AST {
	//TODO: test if this could work
	return &AST{
		rawCtx: c.rawCtx,
		rawAST: C.Z3_mk_string(c.rawCtx, C.CString(str)),
	}
}



// RealSeq returns the seq type number.
func (c *Context) RealSet(reals ...float64) *AST {
	set := &AST{
		rawCtx: c.rawCtx,
		rawAST: C.Z3_mk_empty_set(
			c.rawCtx,
			c.RealSort().rawSort,
		),
	}
	for _, content := range reals {
		C.Z3_mk_set_add(
			c.rawCtx,
			set.rawAST,
			c.Float(content).rawAST,
		)
	}
	return set
}

// StringSet returns the seq type string.
func (c *Context) StringSet(strings ...string) *AST {
	set := &AST{
		rawCtx: c.rawCtx,
		rawAST: C.Z3_mk_empty_set(
			c.rawCtx,
			c.StringSort().rawSort,
		),
	}
	for _, content := range strings {
		C.Z3_mk_set_add(
			c.rawCtx,
			set.rawAST,
			c.Str(content).rawAST,
		)
	}
	return set
}


// True creates the value "true".
//
// Maps: Z3_mk_true
func (c *Context) True() *AST {
	return &AST{
		rawCtx: c.rawCtx,
		rawAST: C.Z3_mk_true(c.rawCtx),
	}
}

// False creates the value "false".
//
// Maps: Z3_mk_false
func (c *Context) False() *AST {
	return &AST{
		rawCtx: c.rawCtx,
		rawAST: C.Z3_mk_false(c.rawCtx),
	}
}

//-------------------------------------------------------------------
// Value Readers
//-------------------------------------------------------------------

// Int gets the integer value of this AST. The value must be able to fit
// into a machine integer.
func (a *AST) Int() int {
	var dst C.int
	C.Z3_get_numeral_int(a.rawCtx, a.rawAST, &dst)
	return int(dst)
}

// Provides an interface to the AST simplifier used by Z3.
//
// Maps: Z3_simplify
func (a *AST) Simplify() *AST {
	return &AST{
		rawCtx: a.rawCtx,
		rawAST: C.Z3_simplify(a.rawCtx, a.rawAST),
	}
}

// Provides an interface to the AST simplifier used by Z3.
//
// Maps: Z3_simplify
func (a *AST) SimplifyEx(p *Params) *AST {
	return &AST{
		rawCtx: a.rawCtx,
		rawAST: C.Z3_simplify_ex(a.rawCtx, a.rawAST, p.rawParams),
	}
}

/* Return a string describing all available parameters:
	algebraic_number_evaluator (bool) simplify/evaluate expressions containing (algebraic) irrational numbers. (default: true)
	arith_ineq_lhs (bool) rewrite inequalities so that right-hand-side is a constant. (default: false)
	arith_lhs (bool) all monomials are moved to the left-hand-side, and the right-hand-side is just a constant. (default: false)
	bit2bool (bool) try to convert bit-vector terms of size 1 into Boolean terms (default: true)
	blast_distinct (bool) expand a distinct predicate into a quadratic number of disequalities (default: false)
	blast_distinct_threshold (unsigned int) when blast_distinct is true, only distinct expressions with less than this number of arguments are blasted (default: 4294967295)
	blast_eq_value (bool) blast (some) Bit-vector equalities into bits (default: false)
	blast_select_store (bool) eagerly replace all (select (store ..) ..) term by an if-then-else term (default: false)
	bv_extract_prop (bool) attempt to partially propagate extraction inwards (default: false)
	bv_ineq_consistency_test_max (unsigned int) max size of conjunctions on which to perform consistency test based on inequalities on bitvectors. (default: 0)
	bv_ite2id (bool) rewrite ite that can be simplified to identity (default: false)
	bv_le_extra (bool) additional bu_(u/s)le simplifications (default: false)
	bv_not_simpl (bool) apply simplifications for bvnot (default: false)
	bv_sort_ac (bool) sort the arguments of all AC operators (default: false)
	cache_all (bool) cache all intermediate results. (default: false)
	elim_and (bool) conjunctions are rewritten using negation and disjunctions (default: false)
	elim_ite (bool) eliminate ite in favor of and/or (default: true)
	elim_rem (bool) replace (rem x y) with (ite (>= y 0) (mod x y) (- (mod x y))). (default: false)
	elim_sign_ext (bool) expand sign-ext operator using concat and extract (default: true)
	elim_to_real (bool) eliminate to_real from arithmetic predicates that contain only integers. (default: false)
	eq2ineq (bool) expand equalities into two inequalities (default: false)
	expand_nested_stores (bool) replace nested stores by a lambda expression (default: false)
	expand_power (bool) expand (^ t k) into (* t ... t) if  1 < k <= max_degree. (default: false)
	expand_select_store (bool) conservatively replace a (select (store ...) ...) term by an if-then-else term (default: false)
	expand_store_eq (bool) reduce (store ...) = (store ...) with a common base into selects (default: false)
	expand_tan (bool) replace (tan x) with (/ (sin x) (cos x)). (default: false)
	flat (bool) create nary applications for and,or,+,*,bvadd,bvmul,bvand,bvor,bvxor (default: true)
	gcd_rounding (bool) use gcd rounding on integer arithmetic atoms. (default: false)
	hi_div0 (bool) use the 'hardware interpretation' for division by zero (for bit-vector terms) (default: true)
	hoist_ite (bool) hoist shared summands under ite expressions (default: false)
	hoist_mul (bool) hoist multiplication over summation to minimize number of multiplications (default: false)
	ignore_patterns_on_ground_qbody (bool) ignores patterns on quantifiers that don't mention their bound variables. (default: true)
	ite_extra_rules (bool) extra ite simplifications, these additional simplifications may reduce size locally but increase globally (default: false)
	local_ctx (bool) perform local (i.e., cheap) context simplifications (default: false)
	local_ctx_limit (unsigned int) limit for applying local context simplifier (default: 4294967295)
	max_degree (unsigned int) max degree of algebraic numbers (and power operators) processed by simplifier. (default: 64)
	max_memory (unsigned int) maximum amount of memory in megabytes (default: 4294967295)
	max_steps (unsigned int) maximum number of steps (default: 4294967295)
	mul2concat (bool) replace multiplication by a power of two into a concatenation (default: false)
	mul_to_power (bool) collpase (* t ... t) into (^ t k), it is ignored if expand_power is true. (default: false)
	pull_cheap_ite (bool) pull if-then-else terms when cheap. (default: false)
	push_ite_arith (bool) push if-then-else over arithmetic terms. (default: false)
	push_ite_bv (bool) push if-then-else over bit-vector terms. (default: false)
	push_to_real (bool) distribute to_real over * and +. (default: true)
	rewrite_patterns (bool) rewrite patterns. (default: false)
	som (bool) put polynomials in sum-of-monomials form (default: false)
	som_blowup (unsigned int) maximum increase of monomials generated when putting a polynomial in sum-of-monomials normal form (default: 10)
	sort_store (bool) sort nested stores when the indices are known to be different (default: false)
	sort_sums (bool) sort the arguments of + application. (default: false)
	split_concat_eq (bool) split equalities of the form (= (concat t1 t2) t3) (default: false)
*/
// Maps: Z3_simplify_get_help
func (a *AST) SimplifyGetHelp() string {
	return C.GoString(C.Z3_simplify_get_help(a.rawCtx))
}

