package z3

type SnellerSimplifier struct {
	ctx *Context
	tactic *Tactic
}

func NewSnellerSimplifier(ctx *Context) *SnellerSimplifier {
	var ctx_solver_simplify *Tactic = ctx.MkTactic("ctx-solver-simplify")
	var propagate_values *Tactic = ctx.MkTactic("propagate-values")
	var split_clause *Tactic = ctx.MkTactic("split-clause")
	var propagate_ineqs *Tactic = ctx.MkTactic("propagate-ineqs")
	var skip *Tactic = ctx.MkTactic("skip")
	var tactic *Tactic = ctx.AndThen(ctx_solver_simplify, ctx.AndThen(propagate_values, ctx.AndThen(ctx.Repeat(ctx.OrElse(split_clause, skip), 10), propagate_ineqs)))

	defer ctx_solver_simplify.Close()
	defer propagate_values.Close()
	defer split_clause.Close()
	defer propagate_ineqs.Close()
	defer skip.Close()

	return &SnellerSimplifier{
		ctx: ctx,
		tactic: tactic,
	}
}

func (s*SnellerSimplifier) Close() error {
	s.ctx.Close()
	s.tactic.Close()
	return nil
}


func (s *SnellerSimplifier) Simplify(x *AST) []*AST {
	var goal *Goal = s.ctx.MkGoal(true, false, false)
	goal.Assert(x)
	var apply_result *ApplyResult = s.tactic.Apply(goal)

	if (apply_result.GetNumSubgoals() > 0) {
		var subgoal *Goal = apply_result.GetSubgoal(0)

		if size := subgoal.GetGoalSize(); size > 0 {
			var result = make([]*AST, size)
			for i := 0; i<size; i++ {
				result[i] = subgoal.GetFormula(i)
			}
			return result
		} else {
			return []*AST{s.ctx.True()}
		}
	} else {
		return []*AST{s.ctx.False()}
	}
}

