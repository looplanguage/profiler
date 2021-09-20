package benchmark

import (
	"bytes"
	"encoding/gob"
	"fmt"
	"github.com/looplanguage/compiler/compiler"
	"github.com/looplanguage/lpvm/vm"
	"io/ioutil"
	"log"
	"os"
	"time"
)

func BenchmarkFile(path string) {
	log.Printf("Benchmarking: %s\n", path)

	log.Printf("Started")
	start := time.Now()
	StartVM("")
	log.Printf("Done!")
	elapsed := time.Since(start)
	log.Printf("It took: %s", elapsed)
}

func StartVM(path string) {
	compiler.RegisterGobTypes()
	bts, err := ioutil.ReadFile(os.Args[1])

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
	err = machine.Run()

	if err != nil {
		log.Fatal(err)
	}
}