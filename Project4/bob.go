package main

import (
	"crypto/aes"
	"fmt"
)

// Bob the evaluator

func evaluateGarbledCircuit(inputs [][]byte, gates []Gate) []byte {
	n := len(inputs) // number of inputs
	m := len(gates)  // number of gates

	// array of signals have a size of n+m
	signals := make([][]byte, m+n)

	// setup inputs signals
	for i := 0; i < n; i++ {
		signals[i] = inputs[i]
		fmt.Printf("input %d=%x\n", i, signals[i][:2])
	}

	// add code below to evaluate the gates

	for _, gate := range gates {

		a := signals[gate.in0]
		b := signals[gate.in1]
		test1 := (a[0] & 0x80) >> 7
		test2 := (b[0] & 0x80) >> 7
		test3 := test1*2 + test2
		sel := gate.table[test3]
		c, _ := aes.NewCipher(append(a, b...))
		test4 := make([]byte, 16)
		c.Decrypt(test4, sel)
		signals[gate.out] = test4
	}
	// the last signal is the output
	return signals[n+m-1]

}
