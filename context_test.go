package z3

import (
	"testing"
)

func TestContext(t *testing.T) {
	config := MkConfig()
	defer config.Close()

	ctx := MkContext(config)
	ctx.Close()
}
