package ion

import (
	"fmt"
	"testing"
)

func Test_CFFI(t *testing.T) {
	if CFFI() != 0 {
		t.Errorf("X")
	}
}

func Test_Go(t *testing.T) {

	fmt.Printf("Go\n")
	var err error

	{
		ty := Type{code: TypeInt, bits: 32, lanes: 1}

		var min0, extent0, min1, extent1, v *Port
		if min0, err = NewPort("min0", ty); err != nil {
			t.Error(err)
		}

		if extent0, err = NewPort("extent0", ty); err != nil {
			t.Error(err)
		}

		if min1, err = NewPort("min1", ty); err != nil {
			t.Error(err)
		}

		if extent1, err = NewPort("extent1", ty); err != nil {
			t.Error(err)
		}

		if v, err = NewPort("v", ty); err != nil {
			t.Error(err)
		}

		_ = min0
		_ = extent0
		_ = min1
		_ = extent1
		_ = v

		var v41 *Param
		if v41, err = NewParam("v", "41"); err != nil {
			t.Error(err)
		}

		_ = v41

		var b *Builder
		if b, err = NewBuilder(); err != nil {
			t.Error(err)
		}

		_ = b

		/*
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
		*/
	}

	{
		/*
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
		*/
	}
}
