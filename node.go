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

type Node struct {
	n C.ion_node_t
}

func NewNode() (*Node, error) {
	n := Node{}
	ret := C.ion_node_create(&n.n)
	if ret != 0 {
		return nil, errors.New("ion_node_create")
	}
	return &n, nil
}

func DeleteNode(n *Node) error {
	ret := C.ion_node_destroy(n.n)
	if ret != 0 {
		return errors.New("ion_node_destroy")
	}
	return nil
}

func (n Node) GetPort(key string) (*Port, error) {
	p := Port{}
	ret := C.ion_node_get_port(n.n, C.CString(key), &p.p)
	if ret != 0 {
		return nil, errors.New("ion_node_get_port")
	}
	return &p, nil
}

func (n Node) SetPort(ports []*Port) error {
	ps := make([]C.ion_port_t, len(ports))

	for i := 0; i < len(ports); i++ {
		ps[i] = ports[i].p
	}

	ret := C.ion_node_set_port(n.n, &ps[0], C.int(len(ps)))
	if ret != 0 {
		return errors.New("ion_node_set_port")
	}
	return nil
}

func (n Node) SetParam(params []*Param) error {
	ps := make([]C.ion_param_t, len(params))

	for i := 0; i < len(params); i++ {
		ps[i] = params[i].p
	}

	ret := C.ion_node_set_param(n.n, &ps[0], C.int(len(ps)))
	if ret != 0 {
		return errors.New("ion_node_set_param")
	}
	return nil
}
