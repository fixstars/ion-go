package ion

// #cgo CFLAGS: -I${SRCDIR}/ion-core
// #cgo LDFLAGS: -L${SRCDIR}/ion-core -lstdc++ -lion
// #include <stdio.h>
// #include <stdlib.h>
// #include "c_ion.h"
import "C"

import (
	"errors"
)

type Param struct {
	p C.ion_param_t
}

func NewParam(key, value string) (*Param, error) {
	p := Param{}
	ret := C.ion_param_create(&p.p, C.CString(key), C.CString(value))
	if ret != 0 {
		return nil, errors.New("ion_param_create")
	}
	return &p, nil
}

func DeleteParam(p *Param) error {
	ret := C.ion_param_destroy(p.p)
	if ret != 0 {
		return errors.New("ion_param_destroy")
	}
	return nil
}
