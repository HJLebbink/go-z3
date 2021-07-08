package z3

import (
	"testing"
)

func TestConfig(t *testing.T) {
	c := MkConfig()
	c.SetParamValue("proof", "true")
	c.Close()
}
