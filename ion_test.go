package ion

import (
	"testing"
)

func Test_CFFI(t *testing.T) {
	if CFFI() != 0 {
		t.Errorf("X")
	}
}

func Test_Go(t *testing.T) {

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

		var v41 *Param
		if v41, err = NewParam("v", "41"); err != nil {
			t.Error(err)
		}

		var b *Builder
		if b, err = NewBuilder(); err != nil {
			t.Error(err)
		}

		if err = b.SetTarget("host"); err != nil {
			t.Error(err)
		}

		if err = b.WithBBModule("ion-bb-test"); err != nil {
			t.Error(err)
		}

		var n0 *Node
		if n0, err = b.Add("simple_graph_producer"); err != nil {
			t.Error(err)
		}

		if err = n0.SetParam(v41); err != nil {
			t.Error(err)
		}

		var n1 *Node
		if n1, err = b.Add("simple_graph_consumer"); err != nil {
			t.Error(err)
		}

		var p *Port
		if p, err = n0.Get("output"); err != nil {
			t.Error(err)
		}

		if err = n1.Set(p, min0, extent0, min1, extent1, v); err != nil {
			t.Error(err)
		}

		if err = b.Save("simple_graph.json"); err != nil {
			t.Error(err)
		}
	}

	{
		var b *Builder
		if b, err = NewBuilder(); err != nil {
			t.Error(err)
		}

		if err = b.Load("simple_graph.json"); err != nil {
			t.Error(err)
		}

		if err = b.Compile("simple_graph"); err != nil {
			t.Error(err)
		}
	}
}
