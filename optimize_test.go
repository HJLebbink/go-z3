package z3

import (
	"fmt"
	"testing"
)

func TestOptimize(t *testing.T) {
	config := MkConfig()
	defer config.Close()

	ctx := MkContext(config)
	defer ctx.Close()

	// Create the "x xor y" constraint
	intTyp := ctx.IntSort()
	x := ctx.Const(ctx.Symbol("x"), intTyp)
	y := ctx.Const(ctx.Symbol("y"), intTyp)
	zero := ctx.Int(0, intTyp)
	ten := ctx.Int(10, intTyp)
	eleven := ctx.Int(11, intTyp)

	// Create the optimize
	o := ctx.MkOptimize()
	defer o.Close()

	// Assert constraints
	o.Add(ten.Ge(x).And(x.Ge(zero)))
	o.Add(ten.Ge(y).And(y.Ge(zero)))
	o.Add(x.Add(y).Le(eleven))

	handleX := o.Maximize(x)
	handleY := o.Maximize(y)

	// Optimize
	result := o.Check()
	if result != True {
		t.Fatalf("bad: %s", result)
	}

	if upperX := o.Upper(handleX).Int(); upperX != 10 {
		t.Fatalf("bad: %d", upperX)
	}
	if lowerX := o.Lower(handleX).Int(); lowerX != 10 {
		t.Fatalf("bad: %d", lowerX)
	}
	if upperY := o.Upper(handleY).Int(); upperY != 1 {
		t.Fatalf("bad: %d", upperY)
	}
	if lowerY := o.Lower(handleY).Int(); lowerY != 1 {
		t.Fatalf("bad: %d", lowerY)
	}

	// Get the model
	m := o.Model()
	defer m.Close()
	fmt.Printf("\nModel:\n%s", m.String())
}