package benchmark

import (
	"bytes"
	"encoding/gob"
	"fmt"
	"github.com/looplanguage/compiler/code"
	"github.com/looplanguage/compiler/compiler"
	"github.com/looplanguage/lpvm/vm"
	"io/ioutil"
	"log"
	"os"
	"sort"
	"time"
)

func StartBenchmark(path string) {
	log.Printf("Benchmarking: %s\n", os.Args[1])

	log.Printf("Started")
	start := time.Now()
	StartVM(path)
	log.Printf("Done!")
	elapsed := time.Since(start)
	log.Printf("It took: %s", elapsed)
}

type OpcodeRan struct {
	Op     code.OpCode
	Called int
	Time   time.Duration
}

type Profiler struct {
	OpCodes []OpcodeRan
}

type LastExecutor struct {
	code code.OpCode
	time time.Time
	ran  bool
}

func FindInProfiler(opCode code.OpCode) (*OpcodeRan, int) {
	for k, v := range profiler.OpCodes {
		if v.Op == opCode {
			return &v, k
		}
	}

	return nil, 0
}

var profiler Profiler = Profiler{OpCodes: []OpcodeRan{}}
var lastExecution LastExecutor

func RanOpcode(code code.OpCode) {
	if lastExecution.ran {
		elapsed := time.Since(lastExecution.time)

		current, key := FindInProfiler(lastExecution.code)

		if current != nil {
			current.Called += 1
			current.Time += elapsed
		}

		profiler.OpCodes[key] = *current
	}

	if val, _ := FindInProfiler(code); val == nil {
		profiler.OpCodes = append(profiler.OpCodes, OpcodeRan{Op: code, Called: 0, Time: 0})
	}

	lastExecution = LastExecutor{
		code: code,
		time: time.Now(),
		ran:  true,
	}
}

func StartVM(path string) {
	compiler.RegisterGobTypes()
	bts, err := ioutil.ReadFile(path)

	if err != nil {
		log.Fatalln(err)
	}

	var constantBytes bytes.Buffer
	constantBytes.Write(bts)

	dec := gob.NewDecoder(&constantBytes)
	var Bytecode compiler.Bytecode
	err = dec.Decode(&Bytecode)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(Bytecode.Instructions.String())

	machine := vm.Create(&Bytecode)
	err = machine.Run(RanOpcode)

	fmt.Println("Profiler data: ")

	sort.SliceStable(profiler.OpCodes, func(i, j int) bool {
		return profiler.OpCodes[i].Time > profiler.OpCodes[j].Time
	})

	if len(profiler.OpCodes) == 0 {
		fmt.Println("No data!")
	} else {
		for op, val := range profiler.OpCodes {
			disassembled, ok := code.Lookup(byte(op))

			if ok != nil {
				fmt.Printf("Was unable to lookup byte: %d (%s)\n", op, err)
			} else {
				fmt.Printf("[%s] Execution time: %s. Called: %d\n", disassembled.Name, val.Time, val.Called)
			}
		}
	}

	if err != nil {
		log.Fatal(err)
	}

}
