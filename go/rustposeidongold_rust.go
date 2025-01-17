//go:build amd64

package poseidongold

/*
#cgo CFLAGS: -I.
#cgo LDFLAGS: -L. -lrustposeidongold
#include "lib.h"
*/
import "C"

import (
	"unsafe"
)

func HashWithResult(in *[8]uint64, capacity *[4]uint64, result *[4]uint64) {
	cInput := (*C.ulonglong)(unsafe.Pointer(&in[0]))
	cCapacity := (*C.ulonglong)(unsafe.Pointer(&capacity[0]))
	cResult := (*C.ulonglong)(unsafe.Pointer(&result[0]))

	C.hash12(cInput, cCapacity, cResult)
}
