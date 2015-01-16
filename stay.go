package main

import (
	"fmt"
)

const (
	INIT = iota
	TRACE
	PUSH
	ADD
	PRINT
	HALT
	GOTO
)

type stayVM struct {
	code     []string
	codeArgs []int

	program []int
	pc      int

	stack []int
	sc    int

	stackSizeKB int
	shouldTrace bool
}

func (vm *stayVM) trace() {
	fmt.Print(vm.pc-1, " | ", vm.code[vm.program[vm.pc]])
	if vm.codeArgs[vm.program[vm.pc]] > 0 {
		numOfArgs := vm.codeArgs[vm.program[vm.pc]]
		fmt.Print(" ", vm.program[vm.pc:vm.pc+numOfArgs])
	}
	fmt.Println("\n-->", vm.stack, "\n")
}

func (vm *stayVM) run(program []int) {

	vm.code = []string{
		"INIT",
		"TRACE",
		"PUSH",
		"ADD",
		"PRINT",
		"HALT",
		"GOTO",
	}

	vm.codeArgs = []int{
		0,
		0,
		1,
		0,
		0,
		0,
		1,
	}

	vm.program = program
	vm.pc = 0

	vm.stackSizeKB = 128
	vm.stack = make([]int, vm.stackSizeKB/4)
	vm.sc = -1
	vm.shouldTrace = false

	for {
		vm.shouldTrace = false
		if vm.shouldTrace {
			vm.trace()
		}

		opp := vm.program[vm.pc]

		switch opp {
		case INIT:

		case TRACE:
			vm.shouldTrace = true

		case PUSH:
			vm.sc++
			vm.pc++

			vm.stack[vm.sc] = vm.program[vm.pc]

		case ADD:
			a := vm.stack[vm.sc]
			vm.stack[vm.sc] = 0

			vm.sc--
			b := vm.stack[vm.sc]

			vm.stack[vm.sc] = a + b

		case PRINT:
			fmt.Println(vm.stack[vm.sc])
		case GOTO:
			vm.pc = vm.program[vm.pc+1]
		case HALT:
			return
		}

		vm.pc++
	}
}

func main() {
	program := []int{
		PUSH, 1,
		PUSH, 12,
		ADD,
		PRINT,
		GOTO, 0,
		HALT,
	}

	stay := &stayVM{}
	stay.run(program)
}
