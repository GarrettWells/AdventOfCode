package main

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"

	"github.com/GarrettWells/AdventOfCode/util"
)

var gates = map[string]*Gate{}

type identifier = string

type Circuit interface {
	evaluate() uint16
}

type Not struct {
	input Input
}

func (self *Not) evaluate() uint16 {
	return ^self.input.evaluate()
}

type Lshift struct {
	input Input
	shift uint8
}

func (self *Lshift) evaluate() uint16 {
	return self.input.evaluate() << self.shift
}

type Rshift struct {
	input Input
	shift uint8
}

func (self *Rshift) evaluate() uint16 {
	return self.input.evaluate() >> self.shift
}

type And struct {
	input1 Input
	input2 Input
}

func (self *And) evaluate() uint16 {
	return self.input1.evaluate() & self.input2.evaluate()
}

type Or struct {
	input1 Input
	input2 Input
}

func (self *Or) evaluate() uint16 {
	return self.input1.evaluate() | self.input2.evaluate()
}

type NoOp struct {
	input Input
}

func (self *NoOp) evaluate() uint16 {
	return self.input.evaluate()
}

type Input interface {
	evaluate() uint16
}

type FromGate struct {
	id identifier
}

func (self *FromGate) evaluate() uint16 {
	return gates[self.id].getOutput()
}

type BaseSignal struct {
	signal uint16
}

func (self *BaseSignal) evaluate() uint16 {
	return self.signal
}

type Gate struct {
	id        string
	output    uint16
	circuit   Circuit
	evaluated bool
}

func (self *Gate) getOutput() uint16 {
	if self.evaluated {
		return self.output
	}

	self.output = self.circuit.evaluate()
	self.evaluated = true
	return self.output
}

func createInput(input string) Input {
	_input, err := strconv.Atoi(input)
	if err == nil {
		return &BaseSignal{uint16(_input)}
	}
	return &FromGate{input}
}

func createGate(input string) *Gate {
	notRegex := regexp.MustCompile(`NOT (?P<input>[[:alpha:]]+) -> (?P<id>[[:alpha:]]*)`)
	andRegex := regexp.MustCompile(`(?P<input1>\w+) AND (?P<input2>[[:alpha:]]+) -> (?P<id>[[:alpha:]]*)`)
	orRegex := regexp.MustCompile(`(?P<input1>\w+) OR (?P<input2>[[:alpha:]]+) -> (?P<id>[[:alpha:]]*)`)
	lshiftRegex := regexp.MustCompile(`(?P<input>[[:alpha:]]+) LSHIFT (?P<shift>[[:digit:]]+) -> (?P<id>[[:alpha:]]*)`)
	rshiftRegex := regexp.MustCompile(`(?P<input>[[:alpha:]]+) RSHIFT (?P<shift>[[:digit:]]+) -> (?P<id>[[:alpha:]]*)`)
	directRegex := regexp.MustCompile(`(?P<input>\w*) -> (?P<id>[[:alpha:]]*)`)

	output := util.CreateMap(input, notRegex)
	if output != nil {
		return &Gate{output["id"], 0, &Not{createInput(output["input"])}, false}
	}

	output = util.CreateMap(input, andRegex)
	if output != nil {
		return &Gate{output["id"], 0, &And{createInput(output["input1"]), createInput(output["input2"])}, false}
	}

	output = util.CreateMap(input, orRegex)
	if output != nil {
		return &Gate{output["id"], 0, &Or{createInput(output["input1"]), createInput(output["input2"])}, false}
	}

	output = util.CreateMap(input, lshiftRegex)
	if output != nil {
		shift, _ := strconv.Atoi(output["shift"])
		return &Gate{output["id"], 0, &Lshift{createInput(output["input"]), uint8(shift)}, false}
	}

	output = util.CreateMap(input, rshiftRegex)
	if output != nil {
		shift, _ := strconv.Atoi(output["shift"])
		return &Gate{output["id"], 0, &Rshift{createInput(output["input"]), uint8(shift)}, false}
	}

	output = util.CreateMap(input, directRegex)
	if output != nil {
		return &Gate{output["id"], 0, &NoOp{createInput(output["input"])}, false}
	}

	panic("Something is wrong")
}

func part1(input string) uint16 {
	for _, line := range strings.Split(input, "\n") {
		gate := createGate(line)
		gates[gate.id] = gate
	}

	return gates["a"].getOutput()
}

func part2(input string) uint16 {
	bSignal := part1(input)

	for k := range gates {
		delete(gates, k)
	}

	for _, line := range strings.Split(input, "\n") {
		gate := createGate(line)
		gates[gate.id] = gate
	}

	gates["b"] = createGate(fmt.Sprint(bSignal) + " -> b")

	return gates["a"].getOutput()
}

func main() {
	input := util.ReadFile("input.txt")
	fmt.Println(part1(input))
	fmt.Println(part2(input))
}
