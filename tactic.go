package z3

// #include "go-z3.h"
import "C"

type Tactic struct {
	rawCtx   C.Z3_context
	rawTactic C.Z3_tactic
}


// Return a tactic associated with the given name.
func (c *Context) MkTactic(name string) *Tactic {
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

// Close decreases the reference count for this tactic. If nothing else
// has manually increased the reference count, this will free the memory
// associated with it.
func (t *Tactic) Close() error {
	C.Z3_tactic_dec_ref(t.rawCtx, t.rawTactic)
	return nil
}

// Z3_tactic_apply
func (t *Tactic) Apply(g *Goal) *ApplyResult {
	rawApplyResult := C.Z3_tactic_apply(t.rawCtx, t.rawTactic, g.rawGoal)
	C.Z3_apply_result_inc_ref(t.rawCtx, rawApplyResult)
	return &ApplyResult{
		rawCtx: t.rawCtx,
		rawApplyResult: rawApplyResult,
	}
}

/*
Legal parameters are:
  arith.auto_config_simplex (bool) (default: false)
  arith.bprop_on_pivoted_rows (bool) (default: true)
  arith.branch_cut_ratio (unsigned int) (default: 2)
  arith.dump_lemmas (bool) (default: false)
  arith.eager_eq_axioms (bool) (default: true)
  arith.enable_hnf (bool) (default: true)
  arith.greatest_error_pivot (bool) (default: false)
  arith.ignore_int (bool) (default: false)
  arith.int_eq_branch (bool) (default: false)
  arith.min (bool) (default: false)
  arith.nl (bool) (default: true)
  arith.nl.branching (bool) (default: true)
  arith.nl.delay (unsigned int) (default: 500)
  arith.nl.expp (bool) (default: false)
  arith.nl.gr_q (unsigned int) (default: 10)
  arith.nl.grobner (bool) (default: true)
  arith.nl.grobner_cnfl_to_report (unsigned int) (default: 1)
  arith.nl.grobner_eqs_growth (unsigned int) (default: 10)
  arith.nl.grobner_expr_degree_growth (unsigned int) (default: 2)
  arith.nl.grobner_expr_size_growth (unsigned int) (default: 2)
  arith.nl.grobner_frequency (unsigned int) (default: 4)
  arith.nl.grobner_max_simplified (unsigned int) (default: 10000)
  arith.nl.grobner_subs_fixed (unsigned int) (default: 2)
  arith.nl.horner (bool) (default: true)
  arith.nl.horner_frequency (unsigned int) (default: 4)
  arith.nl.horner_row_length_limit (unsigned int) (default: 10)
  arith.nl.horner_subs_fixed (unsigned int) (default: 2)
  arith.nl.nra (bool) (default: true)
  arith.nl.order (bool) (default: true)
  arith.nl.rounds (unsigned int) (default: 1024)
  arith.nl.tangents (bool) (default: true)
  arith.print_ext_var_names (bool) (default: false)
  arith.print_stats (bool) (default: false)
  arith.propagate_eqs (bool) (default: true)
  arith.propagation_mode (unsigned int) (default: 1)
  arith.random_initial_value (bool) (default: false)
  arith.reflect (bool) (default: true)
  arith.rep_freq (unsigned int) (default: 0)
  arith.simplex_strategy (unsigned int) (default: 0)
  arith.solver (unsigned int) (default: 6)
  array.extensional (bool) (default: true)
  array.weak (bool) (default: false)
  auto_config (bool) (default: true)
  bv.delay (bool) (default: true)
  bv.enable_int2bv (bool) (default: true)
  bv.eq_axioms (bool) (default: true)
  bv.reflect (bool) (default: true)
  bv.watch_diseq (bool) (default: false)
  candidate_models (bool) (default: false)
  case_split (unsigned int) (default: 1)
  clause_proof (bool) (default: false)
  core.extend_nonlocal_patterns (bool) (default: false)
  core.extend_patterns (bool) (default: false)
  core.extend_patterns.max_distance (unsigned int) (default: 4294967295)
  core.minimize (bool) (default: false)
  core.validate (bool) (default: false)
  cube_depth (unsigned int) (default: 1)
  dack (unsigned int) (default: 1)
  dack.eq (bool) (default: false)
  dack.factor (double) (default: 0.1)
  dack.gc (unsigned int) (default: 2000)
  dack.gc_inv_decay (double) (default: 0.8)
  dack.threshold (unsigned int) (default: 10)
  delay_units (bool) (default: false)
  delay_units_threshold (unsigned int) (default: 32)
  dt_lazy_splits (unsigned int) (default: 1)
  ematching (bool) (default: true)
  induction (bool) (default: false)
  lemma_gc_strategy (unsigned int) (default: 0)
  logic (symbol) (default: )
  macro_finder (bool) (default: false)
  max_conflicts (unsigned int) (default: 4294967295)
  mbqi (bool) (default: true)
  mbqi.force_template (unsigned int) (default: 10)
  mbqi.id (string) (default: )
  mbqi.max_cexs (unsigned int) (default: 1)
  mbqi.max_cexs_incr (unsigned int) (default: 0)
  mbqi.max_iterations (unsigned int) (default: 1000)
  mbqi.trace (bool) (default: false)
  pb.conflict_frequency (unsigned int) (default: 1000)
  pb.learn_complements (bool) (default: true)
  phase_caching_off (unsigned int) (default: 100)
  phase_caching_on (unsigned int) (default: 400)
  phase_selection (unsigned int) (default: 3)
  pull_nested_quantifiers (bool) (default: false)
  q.lift_ite (unsigned int) (default: 0)
  qi.cost (string) (default: (+ weight generation))
  qi.eager_threshold (double) (default: 10.0)
  qi.lazy_threshold (double) (default: 20.0)
  qi.max_instances (unsigned int) (default: 4294967295)
  qi.max_multi_patterns (unsigned int) (default: 0)
  qi.profile (bool) (default: false)
  qi.profile_freq (unsigned int) (default: 4294967295)
  qi.quick_checker (unsigned int) (default: 0)
  quasi_macros (bool) (default: false)
  random_seed (unsigned int) (default: 0)
  refine_inj_axioms (bool) (default: true)
  relevancy (unsigned int) (default: 2)
  restart.max (unsigned int) (default: 4294967295)
  restart_factor (double) (default: 1.1)
  restart_strategy (unsigned int) (default: 1)
  restricted_quasi_macros (bool) (default: false)
  seq.split_w_len (bool) (default: true)
  seq.validate (bool) (default: false)
  str.aggressive_length_testing (bool) (default: false)
  str.aggressive_unroll_testing (bool) (default: true)
  str.aggressive_value_testing (bool) (default: false)
  str.fast_length_tester_cache (bool) (default: false)
  str.fast_value_tester_cache (bool) (default: true)
  str.fixed_length_naive_cex (bool) (default: true)
  str.fixed_length_refinement (bool) (default: false)
  str.overlap_priority (double) (default: -0.1)
  str.regex_automata_difficulty_threshold (unsigned int) (default: 1000)
  str.regex_automata_failed_automaton_threshold (unsigned int) (default: 10)
  str.regex_automata_failed_intersection_threshold (unsigned int) (default: 10)
  str.regex_automata_intersection_difficulty_threshold (unsigned int) (default: 1000)
  str.regex_automata_length_attempt_threshold (unsigned int) (default: 10)
  str.string_constant_cache (bool) (default: true)
  str.strong_arrangements (bool) (default: true)
  string_solver (symbol) (default: seq)
  theory_aware_branching (bool) (default: false)
  theory_case_split (bool) (default: false)
  threads (unsigned int) (default: 1)
  threads.cube_frequency (unsigned int) (default: 2)
  threads.max_conflicts (unsigned int) (default: 400)
*/
// Z3_tactic_apply_ex
func (t *Tactic) ApplyEx(g *Goal, p *Params) *ApplyResult {
	rawApplyResult := C.Z3_tactic_apply_ex(t.rawCtx, t.rawTactic, g.rawGoal, p.rawParams)
	C.Z3_apply_result_inc_ref(t.rawCtx, rawApplyResult)
	return &ApplyResult{
		rawCtx: t.rawCtx,
		rawApplyResult: rawApplyResult,
	}
}