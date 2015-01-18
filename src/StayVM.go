package main

import (
	"fmt"
	"os"
	"strconv"
)

type stayVM struct {
	program []int
	pl      int

	heap []int

	stack []int
	sl    int

	heapSizeBytes int

	stackSizeBytes int
	shouldTrace    bool

	opps    []string
	oppsArg []int
}

func (vm *stayVM) trace() {
	if vm.shouldTrace {
		// Current Location in Stack/Total Size and the Stack itself
		fmt.Println("SL:", vm.sl, " PL:", vm.pl)
		fmt.Println("[STACK] -> ", vm.stack, "")

		// Current Location in Stack/Total Size and the Stack itself
		fmt.Println("[HEAP ] -> ", vm.heap, "")

		fmt.Print("[OPP ", vm.pl, "] ->  ", vm.opps[vm.program[vm.pl]])
		if vm.oppsArg[vm.program[vm.pl]] > 0 {
			fmt.Print(" ", vm.program[vm.pl+1])
		}
		fmt.Print("\n\n\n")
	}
}

/*
	Function: checkStack()

*/
func (vm *stayVM) checkStack() {
	if vm.sl >= len(vm.stack)-1 {
		var s string = "\n[ERROR] STAY_VM: STACK OVERFLOW - SL " + strconv.Itoa(vm.sl) + "/" + strconv.Itoa(len(vm.stack))
		fmt.Println(s)
		os.Exit(601)
	}
}

/////////////////////////////////////////

func (vm *stayVM) run(program []int) {
	fmt.Println("\n> STAY_VM STARTED\n")
	vm.program = program
	for {

		vm.trace()
		vm.checkStack()

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

		// JUMP TO GIVEN LINE IF GREATER
		case JIG:
			a := vm.stack[vm.sl]
			vm.sl--

			b := vm.stack[vm.sl]

			if a < b {
				vm.pl = vm.program[vm.pl+1] - 1
			} else {
				vm.pl++
			}
		//////////////////////////////////////////

		// JUMP TO GIVEN LINE IF LESS
		case JIL:
			a := vm.stack[vm.sl]
			vm.sl--

			b := vm.stack[vm.sl]

			if a > b {
				vm.pl = vm.program[vm.pl+1] - 1
			} else {
				vm.pl++
			}
		//////////////////////////////////////////

		// JUMP TO GIVEN LINE IF EQUAL
		case JIE:
			a := vm.stack[vm.sl]
			vm.sl--

			b := vm.stack[vm.sl]

			if a == b {
				vm.pl = vm.program[vm.pl+1] - 1
			} else {
				vm.pl++
			}
		//////////////////////////////////////////

		// CLEAR STACK
		case CLRS:
			for i := range vm.stack {
				vm.stack[i] = 0
			}
			vm.sl = -1

		//////////////////////////////////////////

		// CLEAR HEAP
		case CLRH:
			for i := range vm.heap {
				vm.heap[i] = 0
			}
		// MOVE FROM STACK TO HEAP
		case MOV:
			vm.heap[vm.program[vm.pl+1]] = vm.stack[vm.sl]
			vm.pl++

		//////////////////////////////////////////

		// MOVE FROM STACK TO HEAP
		case GET:
			vm.sl++
			vm.stack[vm.sl] = vm.heap[vm.program[vm.pl+1]]
			vm.pl++
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

func (vm *stayVM) init(trace bool, stackSize int, heapSize int) {
	vm.sl = -1 //Init to 0 so when the stack is accesed an increment must be made, bringing sl to 0 as a minimum
	vm.pl = 0

	vm.shouldTrace = trace

	vm.stackSizeBytes = stackSize
	vm.stack = make([]int, vm.stackSizeBytes/4+1)

	vm.heapSizeBytes = heapSize
	vm.heap = make([]int, vm.heapSizeBytes/4+1)

	vm.opps = []string{
		"PUSH",
		"ADD",
		"SUB",
		"GOTO",
		"JIG",
		"JIL",
		"JIE",
		"CLRS",
		"CLRH",
		"PRINT",
		"GET",
		"PUT",
		"HALT",
		"MOV",
		"GET",
	}
	vm.oppsArg = []int{
		1,
		0,
		0,
		1,
		1,
		1,
		1,
		0,
		0,
		0,
		1,
		1,
		0,
		1,
		1,
	}
}

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

		// JUMP IF GREATER
		case JIG:
			fmt.Println(pl, " | ", "JIG")
			pl++
		//////////////////////////////////////////

		// JUMP IF LESS
		case JIL:
			fmt.Println(pl, " | ", "JIL")
			pl++
		//////////////////////////////////////////

		// JUMP IF LESS
		case JIE:
			fmt.Println(pl, " | ", "JIE")
			pl++
		//////////////////////////////////////////

		// CLEAR STACK
		case CLRS:
			fmt.Println(pl, " | ", "CLRS")
			pl++
		//////////////////////////////////////////

		// CLEAR HEAP
		case CLRH:
			fmt.Println(pl, " | ", "CLRS")
			pl++
		//////////////////////////////////////////

		// MOVE FROM STACK TO HEAP
		case MOV:
			fmt.Println(pl, " | ", "MOV")

		//////////////////////////////////////////

		// MOVE FROM STACK TO HEAP
		case GET:
			fmt.Println(pl, " | ", "GETw")

		//////////////////////////////////////////

		// FORCE EXITS PROGRAM WITH NO ERROR CODE
		case HALT:
			fmt.Println(pl, " | ", "HALT")
		//////////////////////////////////////////
		default:
			fmt.Println("[ERROR] STAY_VM: SYNTAX ERROR AT ", pl, " - INVALID TYPE ", `"`, program[pl], `"`)
			os.Exit(602)
		}
		pl++

		if pl >= len(program) {
			fmt.Println("-------------------------\n")
			return
		}
	}
}

//--------------------------------->

const (
	PUSH = iota
	ADD
	SUB
	GOTO
	JIG
	JIL
	JIE
	CLRS
	CLRH
	PRINT
	GET
	MOV
	HALT
)

//--------------------------------->

func main() {
	/*
		program := []int{
			PUSH, 0,
			PUSH, 2,
			ADD,
			PUSH, 128,
			JIE, 11,
			GOTO, 2,
			PRINT,
			CLRS,
			HALT,
		}
	*/

	program := []int{
		PUSH, 1,
		GET, 0,
		ADD,
		MOV, 0,
		CLRS,
		GOTO, 0,
		HALT,
	}

	stay := &stayVM{}

	stay.init(true, 128, 128)
	stay.check(program)
	stay.run(program)
}
