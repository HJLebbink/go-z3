package z3

import (
	"fmt"
	"testing"
)

func TestParams_SetBool(t *testing.T) {
	var config = NewConfig()
	defer config.Close()
	var ctx = NewContext(config)
	defer ctx.Close()
	var params = ctx.NewParams()
	defer params.Close()

	params.SetBool(ctx.Symbol("ctx-solver-simplify"), true)
	fmt.Printf("%v\n", params.String())
}