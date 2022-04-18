# Go-LLVM-Bindings
This module is a copy from the original folder of the [llvm-project](https://github.com/llvm/llvm-project/tree/main/llvm/bindings/go).
I only added some additional functionality (namely irreader.go) which was missing for me.

I may extend this module in the future.

To build a project which uses this module you should use the provided build.sh file.
It sets up some environment variables for the cgo bindings.
You need to have llvm-config.exe in your PATH (or modify build.sh to point to your local llvm-config.exe)