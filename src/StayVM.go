package main

import (
	"fmt"
)

type stayVM struct {
	program []int
	pl      int

	stack []int
	sl    int

	heap []int

	stackSizeBytes int
	shouldTrace    bool

	opps    []string
	oppsArg []int
}

func (vm *stayVM) trace() {
	if vm.shouldTrace {
		fmt.Println("\n[SL:", vm.sl, "] -> ", vm.stack, "\n")
		fmt.Print(vm.pl, " | ", vm.opps[vm.program[vm.pl]])

		if vm.oppsArg[vm.program[vm.pl]] > 0 {
			fmt.Print(" ", vm.program[vm.pl+1])
		}
	}
}

func (vm *stayVM) run(program []int) {
	vm.program = program
	for {
		vm.trace()
		switch vm.program[vm.pl] {

		// PUSHES VALUE TO STACK
		case PUSH:
			vm.sl++
			vm.pl++

			vm.stack[vm.sl] = vm.program[vm.pl]

		//////////////////////////////////////////

		// ADDS VM.sl AND VM.sl-1 TOGETHER AND PLACES AT VM.sl-1
		case ADD:
			a := vm.stack[vm.sl]
			vm.stack[vm.sl] = 0

			vm.sl--

			b := vm.stack[vm.sl]
			vm.stack[vm.sl] = a + b
		//////////////////////////////////////////

		// SUBS VM.sl FROM VM.sl-1 TOGETHER AND PLACES AT VM.sl-1
		case SUB:
			a := vm.stack[vm.sl]
			vm.stack[vm.sl] = 0

			vm.sl--

			b := vm.stack[vm.sl]
			vm.stack[vm.sl] = b - a
		//////////////////////////////////////////

		// JUMP TO VM.PROGRAM[VM.PC]
		case GOTO:
			vm.pl = vm.program[vm.pl+1] - 1
		//////////////////////////////////////////

		// SUBS VM.sl FROM VM.sl-1 TOGETHER AND PLACES AT VM.sl-1
		case PRINT:
			fmt.Println(vm.stack[vm.sl])
		//////////////////////////////////////////

		// FORCE EXITS PROGRAM WITH NO ERROR CODE
		case JIG:
			a := vm.stack[vm.sl]
			vm.stack[vm.sl] = 0

			vm.sl--

			b := vm.stack[vm.sl]

			if a > b {
				vm.pl = vm.program[vm.pl+1]
			} else {
				vm.pl++
			}

		//////////////////////////////////////////

		// FORCE EXITS PROGRAM WITH NO ERROR CODE
		case HALT:
			fmt.Println("\n> STAY_VM TERMINATED    ")
			return
		}
		//////////////////////////////////////////

		vm.pl++
	}
}

/////////////////////////////////////////

//--------------------------------->

const (
	PUSH = iota
	ADD
	SUB
	GOTO
	JIG
	PRINT
	GET
	PUT
	HALT
)

//--------------------------------->

func (vm *stayVM) init(trace bool, stackSize int, heapSize int) {
	vm.sl = -1 //Init to 0 so when the stack is accesed an increment must be made, bringing sl to 0 as a minimum
	vm.pl = 0

	vm.shouldTrace = trace
	vm.stackSizeBytes = stackSize

	vm.stackSizeBytes = 32
	vm.stack = make([]int, vm.stackSizeBytes/4)

	vm.opps = []string{
		"PUSH",
		"ADD",
		"SUB",
		"GOTO",
		"JIG",
		"PRINT",
		"GET",
		"PUT",
		"HALT",
	}
	vm.oppsArg = []int{
		1,
		0,
		0,
		1,
		1,
		0,
		1,
		1,
		0,
	}
}

//--------------------------------->

//--------------------------------->

func (vm *stayVM) check(program []int) {
	fmt.Println("-------------------------")
	fmt.Println(" > StayVM Code Overview |")
	fmt.Println("-------------------------")

	pl := 0
	for {
		//vm.trace()
		switch program[pl] {

		// PUSHES VALUE TO STACK
		case PUSH:
			fmt.Println(pl, " | ", "PUSH")
			pl++

		//////////////////////////////////////////

		// ADDS VM.sl AND VM.sl-1 TOGETHER AND PLACES AT VM.sl-1
		case ADD:
			fmt.Println(pl, " | ", "ADD")
		// SUBS VM.sl FROM VM.sl-1 TOGETHER AND PLACES AT VM.sl-1
		case SUB:
			fmt.Println(pl, " | ", "SUB")
		// JUMP TO program[VM.PC]
		case GOTO:
			fmt.Println(pl, " | ", "GOTO")
			pl++

		// SUBS VM.sl FROM VM.sl-1 TOGETHER AND PLACES AT VM.sl-1
		case PRINT:
			fmt.Println(pl, " | ", "PRINT")
		//////////////////////////////////////////

		// FORCE EXITS PROGRAM WITH NO ERROR CODE
		case JIG:
			fmt.Println(pl, " | ", "JIG")
		//////////////////////////////////////////

		// FORCE EXITS PROGRAM WITH NO ERROR CODE
		case HALT:
			fmt.Println(pl, " | ", "HALT")
		//////////////////////////////////////////
		default:
			fmt.Println("STAY_VM: SYNTAX ERROR AT ", pl, " - INVALID TYPE ", `"`, program[pl], `"`)
		}
		pl++

		if pl >= len(program) {
			return
		}
	}
	fmt.Println("-------------------------\n")
}

//--------------------------------->

/////////////////////////////////////////

func main() {
	program := []int{
		PUSH, 1,
		GOTO, 4,
		PRINT,
		HALT,
	}

	stay := &stayVM{}

	stay.init(true, 128, 0)
	stay.check(program)
	stay.run(program)
}
