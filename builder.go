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

type Builder struct {
	b C.ion_builder_t
}

func NewBuilder() (*Builder, error) {
	b := Builder{}
	ret := C.ion_builder_create(&b.b)
	if ret != 0 {
		return nil, errors.New("ion_builder_create")
	}
	return &b, nil
}

func DeleteBuilder(b *Builder) error {
	ret := C.ion_builder_destroy(b.b)
	if ret != 0 {
		return errors.New("ion_builder_destroy")
	}
	return nil
}

func (b Builder) SetTarget(target string) error {
	ret := C.ion_builder_set_target(b.b, C.CString(target))
	if ret != 0 {
		return errors.New("ion_builder_set_target")
	}
	return nil
}

func (b Builder) WithBBModule(module_name string) error {
	ret := C.ion_builder_with_bb_module(b.b, C.CString(module_name))
	if ret != 0 {
		return errors.New("ion_builder_with_bb_module")
	}
	return nil
}

func (b Builder) Add(key string) (*Node, error) {
	n := Node{}
	ret := C.ion_builder_add_node(b.b, C.CString(key), &n.n)
	if ret != 0 {
		return nil, errors.New("ion_builder_add_node")
	}
	return &n, nil
}

func (b Builder) Compile(function_name string) error {
	ret := C.ion_builder_compile(b.b, C.CString(function_name))
	if ret != 0 {
		return errors.New("ion_builder_compile")
	}
	return nil
}

func (b Builder) Save(file_path string) error {
	ret := C.ion_builder_save(b.b, C.CString(file_path))
	if ret != 0 {
		return errors.New("ion_builder_save")
	}
	return nil
}

func (b Builder) Load(file_path string) error {
	ret := C.ion_builder_load(b.b, C.CString(file_path))
	if ret != 0 {
		return errors.New("ion_builder_load")
	}
	return nil
}
