module github.com/looplanguage/profiler

go 1.17

require (
	github.com/looplanguage/compiler v0.1.0
	github.com/looplanguage/lpvm v0.1.0
)

require github.com/looplanguage/loop v0.5.1 // indirect

replace github.com/looplanguage/compiler => ../compiler

replace github.com/looplanguage/lpvm => ../lpvm

replace github.com/looplanguage/loop => ../loop-test
