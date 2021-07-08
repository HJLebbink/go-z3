package z3

import (
	"fmt"
	"testing"
)



//(x > 6) OR (x > 12) -> (x > 12)
//(x > 6) OR (x < 12) -> TRUE
//(x > 6) AND (x < 12) -> (x > 6) AND (x < 12)
//(x < 6) OR (x > 12) -> (x < 6) OR (x > 12)
//(x < 6) AND (x > 12) -> FALSE
//(x > 6) OR (x < 6) -> x != 6



func Test_Simplify1(t *testing.T) {

	config := MkConfig()
	var ctx *Context = MkContext(config)
	config.Close()

	var s *SnellerSimplifier = NewSnellerSimplifier(ctx)

	var a *AST = ctx.Const(ctx.Symbol("a"), ctx.IntSort())
	var int6 = ctx.Int(6, ctx.IntSort())
	var int12 = ctx.Int(12, ctx.IntSort())

	{
		//(a > 6) OR (a > 12) -> (a > 6)
		var y *AST = a.Gt(int6).Or(a.Gt(int12))
		fmt.Printf("%v simplifies to %v\n", y, s.Simplify(y))
	}
	{
		//(a > 6) OR (a < 12) -> TRUE
		var y *AST = a.Gt(int6).Or(a.Lt(int12))
		fmt.Printf("%v simplifies to %v\n", y, s.Simplify(y))
	}
	{
		//(a > 6) AND (a < 12) -> (a > 6) AND (a < 12)
		var y *AST = a.Gt(int6).And(a.Lt(int12))
		fmt.Printf("%v simplifies to %v\n", y, s.Simplify(y))
	}
	{
		//(a < 6) OR (a > 12) -> (a < 6) OR (a > 12)
		var y *AST = a.Lt(int6).Or(a.Gt(int12))
		fmt.Printf("%v simplifies to %v\n", y, s.Simplify(y))
	}
	{
		//(a < 6) AND (a > 12) -> FALSE
		var y *AST = a.Lt(int6).And(a.Gt(int12))
		fmt.Printf("%v simplifies to %v\n", y, s.Simplify(y))
	}
	{
		//(a > 6) OR (a <= 6) -> TRUE
		var y *AST = a.Gt(int6).Or(a.Le(int6))
		fmt.Printf("%v simplifies to %v\n", y, s.Simplify(y))
	}
	{
		//(a >= 6) AND (a <= 6) -> a == 6
		var y *AST = a.Ge(int6).And(a.Le(int6))
		fmt.Printf("%v simplifies to %v\n", y, s.Simplify(y))
	}
	if false { // has a bug...
		//(a > 6) OR (a < 6) -> x != 6
		var y *AST = a.Gt(int6).Or(a.Lt(int6))
		fmt.Printf("%v simplifies to %v\n", y, s.Simplify(y))
	}


	var b *AST = ctx.Const(ctx.Symbol("b"), ctx.BoolSort())
	var c *AST = ctx.Const(ctx.Symbol("c"), ctx.BoolSort())
	{
		//(b AND c) AND (b OR c) -> ?
		var y *AST = b.And(c).And(b.Or(c))
		fmt.Printf("%v simplifies to %v\n", y, s.Simplify(y))
	}
}

