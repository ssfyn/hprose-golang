/*--------------------------------------------------------*\
|                                                          |
|                          hprose                          |
|                                                          |
| Official WebSite: https://hprose.com                     |
|                                                          |
| encoding/relect.go                                       |
|                                                          |
| LastModified: Jun 27, 2020                               |
| Author: Ma Bingyao <andot@hprose.com>                    |
|                                                          |
\*________________________________________________________*/

package encoding

import (
	"container/list"
	"math/big"
	"reflect"
	"time"
	"unsafe"

	"github.com/google/uuid"
)

type eface struct {
	typ uintptr
	ptr unsafe.Pointer
}

func unpackEFace(ptr *interface{}) *eface {
	return (*eface)(unsafe.Pointer(ptr))
}

func unsafeString(bytes []byte) string {
	return *(*string)(unsafe.Pointer(&bytes))
}

// sliceHeader is a safe version of SliceHeader used within this package.
type sliceHeader struct {
	Data unsafe.Pointer
	Len  int
	Cap  int
}

func setSliceHeader(slicePtr unsafe.Pointer, arrayPtr unsafe.Pointer, length int) {
	sliceHeader := (*sliceHeader)(slicePtr)
	sliceHeader.Data = arrayPtr
	sliceHeader.Len = length
	sliceHeader.Cap = length
}

var boolType = reflect.TypeOf(false)
var intType = reflect.TypeOf((int)(0))
var int8Type = reflect.TypeOf((int8)(0))
var int16Type = reflect.TypeOf((int16)(0))
var int32Type = reflect.TypeOf((int32)(0))
var int64Type = reflect.TypeOf((int64)(0))
var uintType = reflect.TypeOf((uint)(0))
var uint8Type = reflect.TypeOf((uint8)(0))
var uint16Type = reflect.TypeOf((uint16)(0))
var uint32Type = reflect.TypeOf((uint32)(0))
var uint64Type = reflect.TypeOf((uint64)(0))
var uintptrType = reflect.TypeOf((uintptr)(0))
var float32Type = reflect.TypeOf((float32)(0))
var float64Type = reflect.TypeOf((float64)(0))
var complex64Type = reflect.TypeOf((complex64)(0))
var complex128Type = reflect.TypeOf((complex128)(0))
var interfaceType = reflect.TypeOf((interface{})(nil))
var bytesType = reflect.TypeOf(([]byte)(nil))
var stringType = reflect.TypeOf("")
var timeType = reflect.TypeOf((*time.Time)(nil)).Elem()
var uuidType = reflect.TypeOf((*uuid.UUID)(nil)).Elem()
var bigIntValueType = reflect.TypeOf((*big.Int)(nil)).Elem()
var bigFloatValueType = reflect.TypeOf((*big.Float)(nil)).Elem()
var bigRatValueType = reflect.TypeOf((*big.Rat)(nil)).Elem()

var boolPtrType = reflect.TypeOf((*bool)(nil))
var intPtrType = reflect.TypeOf((*int)(nil))
var int8PtrType = reflect.TypeOf((*int8)(nil))
var int16PtrType = reflect.TypeOf((*int16)(nil))
var int32PtrType = reflect.TypeOf((*int32)(nil))
var int64PtrType = reflect.TypeOf((*int64)(nil))
var uintPtrType = reflect.TypeOf((*uint)(nil))
var uint8PtrType = reflect.TypeOf((*uint8)(nil))
var uint16PtrType = reflect.TypeOf((*uint16)(nil))
var uint32PtrType = reflect.TypeOf((*uint32)(nil))
var uint64PtrType = reflect.TypeOf((*uint64)(nil))
var uintptrPtrType = reflect.TypeOf((*uintptr)(nil))
var float32PtrType = reflect.TypeOf((*float32)(nil))
var float64PtrType = reflect.TypeOf((*float64)(nil))
var complex64PtrType = reflect.TypeOf((*complex64)(nil))
var complex128PtrType = reflect.TypeOf((*complex128)(nil))
var interfacePtrType = reflect.TypeOf((*interface{})(nil))
var bytesPtrType = reflect.TypeOf((*[]byte)(nil))
var stringPtrType = reflect.TypeOf((*string)(nil))
var timePtrType = reflect.TypeOf((*time.Time)(nil))
var uuidPtrType = reflect.TypeOf((*uuid.UUID)(nil))
var bigIntType = reflect.TypeOf((*big.Int)(nil))
var bigFloatType = reflect.TypeOf((*big.Float)(nil))
var bigRatType = reflect.TypeOf((*big.Rat)(nil))

var listType = reflect.TypeOf((*list.List)(nil))
