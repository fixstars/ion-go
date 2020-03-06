package ion

// #cgo CFLAGS: -I${SRCDIR}/ion-core
// #cgo LDFLAGS: -L${SRCDIR}/ion-core -lstdc++ -lion
// #include <stdio.h>
// #include <stdlib.h>
// #include "c_ion.h"
import "C"

type TypeCode int

const (
	TypeInt    TypeCode = iota
	TypeUInt            = iota
	TypeFloat           = iota
	TypeHandle          = iota
)

type Type struct {
	code  TypeCode
	bits  uint
	lanes uint
}

func CFFI() int {

	ret := C.int(0)

	{
		t := C.ion_type_t{code: C.ion_type_int, bits: 32, lanes: 1}

		var min0, extent0, min1, extent1, v C.ion_port_t
		ret = C.ion_port_create(&min0, C.CString("min0"), t)
		if ret != 0 {
			return int(ret)
		}

		ret = C.ion_port_create(&extent0, C.CString("extent0"), t)
		if ret != 0 {
			return int(ret)
		}

		ret = C.ion_port_create(&min1, C.CString("min1"), t)
		if ret != 0 {
			return int(ret)
		}

		ret = C.ion_port_create(&extent1, C.CString("extent1"), t)
		if ret != 0 {
			return int(ret)
		}

		ret = C.ion_port_create(&v, C.CString("v"), t)
		if ret != 0 {
			return int(ret)
		}

		var v41 C.ion_param_t
		ret = C.ion_param_create(&v41, C.CString("v"), C.CString("41"))
		if ret != 0 {
			return int(ret)
		}

		var b C.ion_builder_t
		ret = C.ion_builder_create(&b)
		if ret != 0 {
			return int(ret)
		}

		ret = C.ion_builder_set_target(b, C.CString("host"))
		if ret != 0 {
			return int(ret)
		}

		ret = C.ion_builder_with_bb_module(b, C.CString("ion-bb-test"))
		if ret != 0 {
			return int(ret)
		}

		var n0 C.ion_node_t
		ret = C.ion_builder_add_node(b, C.CString("simple_graph_producer"), &n0)
		if ret != 0 {
			return int(ret)
		}

		ret = C.ion_node_set_param(n0, &v41, 1)
		if ret != 0 {
			return int(ret)
		}

		var n1 C.ion_node_t
		ret = C.ion_builder_add_node(b, C.CString("simple_graph_consumer"), &n1)
		if ret != 0 {
			return int(ret)
		}

		ports := make([]C.ion_port_t, 6)

		ret = C.ion_node_get_port(n0, C.CString("output"), &ports[0])
		if ret != 0 {
			return int(ret)
		}

		ports[1] = min0
		ports[2] = extent0
		ports[3] = min1
		ports[4] = extent1
		ports[5] = v

		ret = C.ion_node_set_port(n1, &ports[0], 6)
		if ret != 0 {
			return int(ret)
		}

		ret = C.ion_builder_save(b, C.CString("simple_graph.json"))
		if ret != 0 {
			return int(ret)
		}

		ret = C.ion_port_destroy(min0)
		if ret != 0 {
			return int(ret)
		}

		ret = C.ion_port_destroy(extent0)
		if ret != 0 {
			return int(ret)
		}

		ret = C.ion_port_destroy(min1)
		if ret != 0 {
			return int(ret)
		}

		ret = C.ion_port_destroy(extent1)
		if ret != 0 {
			return int(ret)
		}

		ret = C.ion_port_destroy(v)
		if ret != 0 {
			return int(ret)
		}

		ret = C.ion_port_destroy(ports[0])
		if ret != 0 {
			return int(ret)
		}

		ret = C.ion_node_destroy(n0)
		if ret != 0 {
			return int(ret)
		}

		ret = C.ion_node_destroy(n1)
		if ret != 0 {
			return int(ret)
		}
	}

	{
		var b C.ion_builder_t
		ret = C.ion_builder_create(&b)
		if ret != 0 {
			return int(ret)
		}

		ret = C.ion_builder_load(b, C.CString("simple_graph.json"))
		if ret != 0 {
			return int(ret)
		}

		ret = C.ion_builder_compile(b, C.CString("simple_graph"))
		if ret != 0 {
			return int(ret)
		}

		ret = C.ion_builder_destroy(b)
		if ret != 0 {
			return int(ret)
		}
	}

	return 0
}
