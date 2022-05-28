// +build tensorflow

package wasmedge

/*
#cgo linux LDFLAGS: -lwasmedge-tensorflowlite_c -ltensorflow -ltensorflow_framework -ltensorflowlite_c

#include <wasmedge/wasmedge-tensorflowlite.h>
*/
import "C"

func NewTensorflowLiteImportObject() *ImportObject {
	obj := C.WasmEdge_TensorflowLite_ImportObjectCreate()
	if obj == nil {
		return nil
	}
	return &ImportObject{_inner: obj, _own: true}
}
