package z3

import (
	"testing"
)

func TestSymbol(t *testing.T) {
	config := MkConfig()
	defer config.Close()

	ctx := MkContext(config)
	defer ctx.Close()

	// String symbol
	x := ctx.Symbol("x")
	if v := x.String(); v != "x" {
		t.Fatalf("bad: %q", v)
	}

	// Int symbol
	y := ctx.SymbolInt(42)
	if v := y.String(); v != "42" {
		t.Fatalf("bad: %q", v)
	}
}
