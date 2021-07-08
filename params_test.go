package z3

import (
	"fmt"
	"testing"
)

func TestParams_SetBool(t *testing.T) {
	var config = MkConfig()
	defer config.Close()
	var ctx = MkContext(config)
	defer ctx.Close()
	var params = ctx.MkParams()
	defer params.Close()

	params.SetBool("ctx-solver-simplify", true)
	fmt.Printf("%v\n", params.String())
}