/*
	This file is not part of the official llvm Go bindings
	I added it myself because functionality was I needed was missing
*/
package llvm

/*
#include "llvm-c/IRReader.h"
#include "llvm-c/Core.h"
#include <stdlib.h>
*/
import "C"

import (
	"errors"
	"unsafe"
)

// ParseIRFile parses the LLVM IR (textual) in the file with the
// specified name, and returns a new LLVM module.
func ParseIRFile(name string, context Context) (Module, error) {
	var buf C.LLVMMemoryBufferRef
	var errmsg *C.char
	var cfilename *C.char = C.CString(name)
	defer C.free(unsafe.Pointer(cfilename))
	result := C.LLVMCreateMemoryBufferWithContentsOfFile(cfilename, &buf, &errmsg)
	if result != 0 {
		err := errors.New(C.GoString(errmsg))
		C.free(unsafe.Pointer(errmsg))
		return Module{}, err
	}
	//defer C.LLVMDisposeMemoryBuffer(buf)

	var m Module
	result = C.LLVMParseIRInContext(context.C, buf, &m.C, &errmsg)
	if result != 0 {
		err := errors.New(C.GoString(errmsg))
		C.free(unsafe.Pointer(errmsg))
		return Module{}, err
	}

	return m, nil
}
