package ion

// #cgo CFLAGS: -I${SRCDIR}/ion-core
// #cgo LDFLAGS: -L${SRCDIR}/ion-core -lstdc++ -lion
// #include <stdio.h>
// #include <stdlib.h>
// #include "c_ion.h"
import "C"

import (
	"errors"
	"runtime"
)

type Node struct {
	n C.ion_node_t
}

func NewNode() (*Node, error) {
	n := Node{}
	ret := C.ion_node_create(&n.n)
	if ret != 0 {
		return nil, errors.New("ion_node_create")
	}

	runtime.SetFinalizer(&n, func(n *Node) {
		DeleteNode(n)
	})

	return &n, nil
}

func DeleteNode(n *Node) error {
	ret := C.ion_node_destroy(n.n)
	if ret != 0 {
		return errors.New("ion_node_destroy")
	}
	return nil
}

func (n Node) Get(key string) (*Port, error) {
	p := Port{}
	ret := C.ion_node_get_port(n.n, C.CString(key), &p.p)
	if ret != 0 {
		return nil, errors.New("ion_node_get_port")
	}
	return &p, nil
}

func (n Node) Set(ports ...*Port) error {
	ps := []C.ion_port_t{}

	for _, p := range ports {
		ps = append(ps, p.p)
	}

	ret := C.ion_node_set_port(n.n, &ps[0], C.int(len(ps)))
	if ret != 0 {
		return errors.New("ion_node_set_port")
	}
	return nil
}

func (n Node) SetParam(params ...*Param) error {
	ps := []C.ion_param_t{}

	for _, p := range params {
		ps = append(ps, p.p)
	}

	ret := C.ion_node_set_param(n.n, &ps[0], C.int(len(ps)))
	if ret != 0 {
		return errors.New("ion_node_set_param")
	}
	return nil
}
