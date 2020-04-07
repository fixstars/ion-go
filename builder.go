package ion

// #cgo pkg-config: libion
// #include <ion/c_ion.h>
import "C"

import (
	"errors"
	"io"
	"io/ioutil"
	"os"
	"runtime"
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

	runtime.SetFinalizer(&b, func(b *Builder) {
		DeleteBuilder(b)
	})

	return &b, nil
}

func DeleteBuilder(b *Builder) error {
	if b.b == nil {
		return nil
	}
	ret := C.ion_builder_destroy(b.b)
	if ret != 0 {
		return errors.New("ion_builder_destroy")
	}
	b.b = nil
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

func (b Builder) LoadFromReader(r io.Reader) error {
	var err error
	var f *os.File

	if f, err = ioutil.TempFile("", ""); err != nil {
		return err
	}

	if _, err = io.Copy(f, r); err != nil {
		return err
	}
	return b.Load(f.Name())
}

func (b Builder) BBMetadata() (string, error) {
	var n C.int
	ret := C.ion_builder_bb_metadata(b.b, nil, 0, &n)
	if ret != 0 {
		return "", errors.New("ion_builder_bb_metadata")
	}

	buf := make([]C.char, n)
	ret = C.ion_builder_bb_metadata(b.b, &buf[0], n, nil)
	if ret != 0 {
		return "", errors.New("ion_builder_bb_metadata")
	}

	return C.GoString(&buf[0]), nil
}
