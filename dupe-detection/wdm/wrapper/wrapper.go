/* ----------------------------------------------------------------------------
 * This file was automatically generated by SWIG (http://www.swig.org).
 * Version 3.0.12
 *
 * This file is not intended to be easily readable and contains a number of
 * coding conventions designed to improve portability and efficiency. Do not make
 * changes to this file unless you know what you are doing--modify the SWIG
 * interface file instead.
 * ----------------------------------------------------------------------------- */

// source: wrapper.i

package wrapper

/*
#define intgo swig_intgo
typedef void *swig_voidp;

#include <stdint.h>


typedef long long intgo;
typedef unsigned long long uintgo;



typedef struct { char *p; intgo n; } _gostring_;
typedef struct { void* array; intgo len; intgo cap; } _goslice_;


typedef long long swig_type_1;
typedef long long swig_type_2;
typedef long long swig_type_3;
typedef long long swig_type_4;
typedef _gostring_ swig_type_5;
typedef _gostring_ swig_type_6;
typedef _gostring_ swig_type_7;
typedef _gostring_ swig_type_8;
extern void _wrap_Swig_free_wrapper_d8938e3a6795767e(uintptr_t arg1);
extern uintptr_t _wrap_Swig_malloc_wrapper_d8938e3a6795767e(swig_intgo arg1);
extern uintptr_t _wrap_new_DoubleVector__SWIG_0_wrapper_d8938e3a6795767e(void);
extern uintptr_t _wrap_new_DoubleVector__SWIG_1_wrapper_d8938e3a6795767e(swig_type_1 arg1);
extern swig_type_2 _wrap_DoubleVector_size_wrapper_d8938e3a6795767e(uintptr_t arg1);
extern swig_type_3 _wrap_DoubleVector_capacity_wrapper_d8938e3a6795767e(uintptr_t arg1);
extern void _wrap_DoubleVector_reserve_wrapper_d8938e3a6795767e(uintptr_t arg1, swig_type_4 arg2);
extern _Bool _wrap_DoubleVector_isEmpty_wrapper_d8938e3a6795767e(uintptr_t arg1);
extern void _wrap_DoubleVector_clear_wrapper_d8938e3a6795767e(uintptr_t arg1);
extern void _wrap_DoubleVector_add_wrapper_d8938e3a6795767e(uintptr_t arg1, double arg2);
extern double _wrap_DoubleVector_get_wrapper_d8938e3a6795767e(uintptr_t arg1, swig_intgo arg2);
extern void _wrap_DoubleVector_set_wrapper_d8938e3a6795767e(uintptr_t arg1, swig_intgo arg2, double arg3);
extern void _wrap_delete_DoubleVector_wrapper_d8938e3a6795767e(uintptr_t arg1);
extern double _wrap_wdm__SWIG_0_wrapper_d8938e3a6795767e(uintptr_t arg1, uintptr_t arg2, swig_type_5 arg3, uintptr_t arg4, _Bool arg5);
extern double _wrap_wdm__SWIG_1_wrapper_d8938e3a6795767e(uintptr_t arg1, uintptr_t arg2, swig_type_6 arg3, uintptr_t arg4);
extern double _wrap_wdm__SWIG_2_wrapper_d8938e3a6795767e(uintptr_t arg1, uintptr_t arg2, swig_type_7 arg3);
extern double _wrap_wdm__SWIG_3_wrapper_d8938e3a6795767e(swig_voidp arg1, swig_intgo arg2, swig_voidp arg3, swig_intgo arg4, swig_type_8 arg5);
#undef intgo
*/
import "C"

import (
	_ "runtime/cgo"
	"sync"
	"unsafe"
)

type _ unsafe.Pointer

var Swig_escape_always_false bool
var Swig_escape_val interface{}

type _swig_fnptr *byte
type _swig_memberptr *byte

type _ sync.Mutex

func Swig_free(arg1 uintptr) {
	_swig_i_0 := arg1
	C._wrap_Swig_free_wrapper_d8938e3a6795767e(C.uintptr_t(_swig_i_0))
}

func Swig_malloc(arg1 int) (_swig_ret uintptr) {
	var swig_r uintptr
	_swig_i_0 := arg1
	swig_r = (uintptr)(C._wrap_Swig_malloc_wrapper_d8938e3a6795767e(C.swig_intgo(_swig_i_0)))
	return swig_r
}

type SwigcptrDoubleVector uintptr

func (p SwigcptrDoubleVector) Swigcptr() uintptr {
	return (uintptr)(p)
}

func (p SwigcptrDoubleVector) SwigIsDoubleVector() {
}

func NewDoubleVector__SWIG_0() (_swig_ret DoubleVector) {
	var swig_r DoubleVector
	swig_r = (DoubleVector)(SwigcptrDoubleVector(C._wrap_new_DoubleVector__SWIG_0_wrapper_d8938e3a6795767e()))
	return swig_r
}

func NewDoubleVector__SWIG_1(arg1 int64) (_swig_ret DoubleVector) {
	var swig_r DoubleVector
	_swig_i_0 := arg1
	swig_r = (DoubleVector)(SwigcptrDoubleVector(C._wrap_new_DoubleVector__SWIG_1_wrapper_d8938e3a6795767e(C.swig_type_1(_swig_i_0))))
	return swig_r
}

func NewDoubleVector(a ...interface{}) DoubleVector {
	argc := len(a)
	if argc == 0 {
		return NewDoubleVector__SWIG_0()
	}
	if argc == 1 {
		return NewDoubleVector__SWIG_1(a[0].(int64))
	}
	panic("No match for overloaded function call")
}

func (arg1 SwigcptrDoubleVector) Size() (_swig_ret int64) {
	var swig_r int64
	_swig_i_0 := arg1
	swig_r = (int64)(C._wrap_DoubleVector_size_wrapper_d8938e3a6795767e(C.uintptr_t(_swig_i_0)))
	return swig_r
}

func (arg1 SwigcptrDoubleVector) Capacity() (_swig_ret int64) {
	var swig_r int64
	_swig_i_0 := arg1
	swig_r = (int64)(C._wrap_DoubleVector_capacity_wrapper_d8938e3a6795767e(C.uintptr_t(_swig_i_0)))
	return swig_r
}

func (arg1 SwigcptrDoubleVector) Reserve(arg2 int64) {
	_swig_i_0 := arg1
	_swig_i_1 := arg2
	C._wrap_DoubleVector_reserve_wrapper_d8938e3a6795767e(C.uintptr_t(_swig_i_0), C.swig_type_4(_swig_i_1))
}

func (arg1 SwigcptrDoubleVector) IsEmpty() (_swig_ret bool) {
	var swig_r bool
	_swig_i_0 := arg1
	swig_r = (bool)(C._wrap_DoubleVector_isEmpty_wrapper_d8938e3a6795767e(C.uintptr_t(_swig_i_0)))
	return swig_r
}

func (arg1 SwigcptrDoubleVector) Clear() {
	_swig_i_0 := arg1
	C._wrap_DoubleVector_clear_wrapper_d8938e3a6795767e(C.uintptr_t(_swig_i_0))
}

func (arg1 SwigcptrDoubleVector) Add(arg2 float64) {
	_swig_i_0 := arg1
	_swig_i_1 := arg2
	C._wrap_DoubleVector_add_wrapper_d8938e3a6795767e(C.uintptr_t(_swig_i_0), C.double(_swig_i_1))
}

func (arg1 SwigcptrDoubleVector) Get(arg2 int) (_swig_ret float64) {
	var swig_r float64
	_swig_i_0 := arg1
	_swig_i_1 := arg2
	swig_r = (float64)(C._wrap_DoubleVector_get_wrapper_d8938e3a6795767e(C.uintptr_t(_swig_i_0), C.swig_intgo(_swig_i_1)))
	return swig_r
}

func (arg1 SwigcptrDoubleVector) Set(arg2 int, arg3 float64) {
	_swig_i_0 := arg1
	_swig_i_1 := arg2
	_swig_i_2 := arg3
	C._wrap_DoubleVector_set_wrapper_d8938e3a6795767e(C.uintptr_t(_swig_i_0), C.swig_intgo(_swig_i_1), C.double(_swig_i_2))
}

func DeleteDoubleVector(arg1 DoubleVector) {
	_swig_i_0 := arg1.Swigcptr()
	C._wrap_delete_DoubleVector_wrapper_d8938e3a6795767e(C.uintptr_t(_swig_i_0))
}

type DoubleVector interface {
	Swigcptr() uintptr
	SwigIsDoubleVector()
	Size() (_swig_ret int64)
	Capacity() (_swig_ret int64)
	Reserve(arg2 int64)
	IsEmpty() (_swig_ret bool)
	Clear()
	Add(arg2 float64)
	Get(arg2 int) (_swig_ret float64)
	Set(arg2 int, arg3 float64)
}

func Wdm__SWIG_0(arg1 DoubleVector, arg2 DoubleVector, arg3 string, arg4 DoubleVector, arg5 bool) (_swig_ret float64) {
	var swig_r float64
	_swig_i_0 := arg1.Swigcptr()
	_swig_i_1 := arg2.Swigcptr()
	_swig_i_2 := arg3
	_swig_i_3 := arg4.Swigcptr()
	_swig_i_4 := arg5
	swig_r = (float64)(C._wrap_wdm__SWIG_0_wrapper_d8938e3a6795767e(C.uintptr_t(_swig_i_0), C.uintptr_t(_swig_i_1), *(*C.swig_type_5)(unsafe.Pointer(&_swig_i_2)), C.uintptr_t(_swig_i_3), C._Bool(_swig_i_4)))
	if Swig_escape_always_false {
		Swig_escape_val = arg3
	}
	return swig_r
}

func Wdm__SWIG_1(arg1 DoubleVector, arg2 DoubleVector, arg3 string, arg4 DoubleVector) (_swig_ret float64) {
	var swig_r float64
	_swig_i_0 := arg1.Swigcptr()
	_swig_i_1 := arg2.Swigcptr()
	_swig_i_2 := arg3
	_swig_i_3 := arg4.Swigcptr()
	swig_r = (float64)(C._wrap_wdm__SWIG_1_wrapper_d8938e3a6795767e(C.uintptr_t(_swig_i_0), C.uintptr_t(_swig_i_1), *(*C.swig_type_6)(unsafe.Pointer(&_swig_i_2)), C.uintptr_t(_swig_i_3)))
	if Swig_escape_always_false {
		Swig_escape_val = arg3
	}
	return swig_r
}

func Wdm__SWIG_2(arg1 DoubleVector, arg2 DoubleVector, arg3 string) (_swig_ret float64) {
	var swig_r float64
	_swig_i_0 := arg1.Swigcptr()
	_swig_i_1 := arg2.Swigcptr()
	_swig_i_2 := arg3
	swig_r = (float64)(C._wrap_wdm__SWIG_2_wrapper_d8938e3a6795767e(C.uintptr_t(_swig_i_0), C.uintptr_t(_swig_i_1), *(*C.swig_type_7)(unsafe.Pointer(&_swig_i_2))))
	if Swig_escape_always_false {
		Swig_escape_val = arg3
	}
	return swig_r
}

func Wdm__SWIG_3(arg1 *float64, arg2 int, arg3 *float64, arg4 int, arg5 string) (_swig_ret float64) {
	var swig_r float64
	_swig_i_0 := arg1
	_swig_i_1 := arg2
	_swig_i_2 := arg3
	_swig_i_3 := arg4
	_swig_i_4 := arg5
	swig_r = (float64)(C._wrap_wdm__SWIG_3_wrapper_d8938e3a6795767e(C.swig_voidp(_swig_i_0), C.swig_intgo(_swig_i_1), C.swig_voidp(_swig_i_2), C.swig_intgo(_swig_i_3), *(*C.swig_type_8)(unsafe.Pointer(&_swig_i_4))))
	if Swig_escape_always_false {
		Swig_escape_val = arg5
	}
	return swig_r
}

func Wdm(a ...interface{}) float64 {
	argc := len(a)
	if argc == 3 {
		return Wdm__SWIG_2(a[0].(DoubleVector), a[1].(DoubleVector), a[2].(string))
	}
	if argc == 4 {
		return Wdm__SWIG_1(a[0].(DoubleVector), a[1].(DoubleVector), a[2].(string), a[3].(DoubleVector))
	}
	if argc == 5 {
		if _, ok := a[0].(SwigcptrDoubleVector); !ok {
			goto check_3
		}
		if _, ok := a[1].(SwigcptrDoubleVector); !ok {
			goto check_3
		}
		if _, ok := a[2].(string); !ok {
			goto check_3
		}
		if _, ok := a[3].(SwigcptrDoubleVector); !ok {
			goto check_3
		}
		if _, ok := a[4].(bool); !ok {
			goto check_3
		}
		return Wdm__SWIG_0(a[0].(DoubleVector), a[1].(DoubleVector), a[2].(string), a[3].(DoubleVector), a[4].(bool))
	}
check_3:
	if argc == 5 {
		return Wdm__SWIG_3(a[0].(*float64), a[1].(int), a[2].(*float64), a[3].(int), a[4].(string))
	}
	panic("No match for overloaded function call")
}
