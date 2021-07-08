package z3

import (
	"unsafe"
)

// #include <stdlib.h>
// #include "go-z3.h"
import "C"

// Config is used to set configuration for Z3. This should be created with
// MkConfig and closed with Close when you're done using it.
//
// Config structures are used to set parameters for Z3 behavior. See the
// Z3 docs for information on available parameters. They can be set with
// SetParamValue.
//
type Config struct {
	raw C.Z3_config
}

// MkConfig allocates a new configuration object.
func MkConfig() *Config {
	return &Config{
		raw: C.Z3_mk_config(),
	}
}

// Close frees the memory associated with this configuration
func (c *Config) Close() error {
	C.Z3_del_config(c.raw)
	return nil
}

/*
  SetParamValue sets the parameters for a Config.

  auto_config (bool) (default: true)
  debug_ref_count (bool) (default: false)
  dot_proof_file (string) (default: proof.dot)
  dump_models (bool) (default: false)
  model (bool) (default: true)
  model_validate (bool) (default: false)
  proof (bool) (default: false)
  rlimit (unsigned int) (default: 0)
  smtlib2_compliant (bool) (default: false)
  stats (bool) (default: false)
  timeout (unsigned int) (default: 4294967295)
  trace (bool) (default: false)
  trace_file_name (string) (default: z3.log)
  type_check (bool) (default: true)
  unicode (bool)
  unsat_core (bool) (default: false)
  well_sorted_check (bool) (default: false)
 */
func (c *Config) SetParamValue(k, v string) {
	ck := C.CString(k)
	cv := C.CString(v)

	// We free the strings since they're not actually stored
	defer C.free(unsafe.Pointer(ck))
	defer C.free(unsafe.Pointer(cv))

	C.Z3_set_param_value(c.raw, ck, cv)
}

// Z3Value returns the raw internal pointer value. This should only be
// used if you really understand what you're doing. It may be invalid after
// Close is called.
func (c *Config) Z3Value() C.Z3_config {
	return c.raw
}
