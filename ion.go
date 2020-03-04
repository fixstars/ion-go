package ion

// #cgo CFLAGS: -I${SRCDIR}/ion-core
// #cgo LDFLAGS: -L${SRCDIR}/ion-core -lstdc++ -lion
// #include <stdio.h>
// #include <stdlib.h>
// #include "c_ion.h"
import "C"

func CFFI() int {

	var ret C.int
	var b C.ion_builder_t

	ret = C.ion_builder_create(&b)
	if ret != 0 {
		return int(ret)
	}
	ret = C.ion_builder_destroy(b)
	if ret != 0 {
		return int(ret)
	}

	return 0
}
