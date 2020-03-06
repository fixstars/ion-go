package ion

// #cgo CFLAGS: -I${SRCDIR}/ion-core
// #cgo LDFLAGS: -L${SRCDIR}/ion-core -lstdc++ -lion
// #include <stdio.h>
// #include <stdlib.h>
// #include "c_ion.h"
import "C"

import (
	_ "errors"
)

type Port struct {
	p C.ion_port_t
}

// func NewPort(key string, t) (*Port, error) {
// 	p := Port{}
// 	ret := C.ion_port_create(&p.p, key, t)
// 	if ret != 0 {
// 		return nil, errors.New("ion_port_create")
// 	}
// 	return &n, nil
// }
//
// func DeletePort(p *Port) error {
// 	ret := C.ion_port_destroy(p.p)
// 	if ret != 0 {
// 		return errors.New("ion_port_destroy")
// 	}
// 	return nil
// }
