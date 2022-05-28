// +build tensorflow

package wasmedge

/*
#cgo linux LDFLAGS: -lwasmedge-tensorflowlite_c -ltensorflow -ltensorflow_framework -ltensorflowlite_c
#cgo darwin LDFLAGS: -lwasmedge-tensorflow_c -lwasmedge-tensorflowlite_c -ltensorflow -ltensorflow_framework -ltensorflowlite_c

#include <wasmedge/wasmedge-tensorflowlite.h>
*/
import "C"

func NewTensorflowLiteModule() *Module {
	obj := C.WasmEdge_TensorflowLite_ModuleInstanceCreate()
	if obj == nil {
		return nil
	}
	return &Module{_inner: obj, _own: true}
}
