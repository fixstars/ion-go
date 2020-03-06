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
	ret := C.ion_node_destroy(n.n)
	if ret != 0 {
		return nil, errors.New("ion_node_destroy")
	}
	return &p, nil
}
