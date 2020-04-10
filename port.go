package ion

// #cgo pkg-config: ion-core
// #include <ion/c_ion.h>
import "C"

import (
	"errors"
	"runtime"
)

type Port struct {
	p C.ion_port_t
}

func NewPort(key string, ty Type) (*Port, error) {
	p := Port{}
	t := C.ion_type_t{
		code:  C.ion_type_code_t(ty.code),
		bits:  C.uchar(ty.bits),
		lanes: C.uchar(ty.lanes),
	}
	ret := C.ion_port_create(&p.p, C.CString(key), t)
	if ret != 0 {
		return nil, errors.New("ion_port_create")
	}

	runtime.SetFinalizer(&p, func(p *Port) {
		DeletePort(p)
	})

	return &p, nil
}

func DeletePort(p *Port) error {
	ret := C.ion_port_destroy(p.p)
	if ret != 0 {
		return errors.New("ion_port_destroy")
	}
	return nil
}
