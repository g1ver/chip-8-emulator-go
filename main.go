package main

import (
	"fmt"
	"os"
)

const fontSetSize = 80

type (
	currentOpCode  uint8
	register       uint8
	vRegisters     [16]register
	indexRegister  uint16
	programCounter uint16
	chipGraphics   [64 * 32]uint8
	stackData      [16]uint16
	stackPointer   uint16
	chipKeyPad     [16]uint8
	chipMemory     [4096]uint8
	fontSet        [fontSetSize]uint8
	// delayTimer uint8
	// soundTimer uint8
)

// MEMORY LAYOUT (4096 BYTES)
// 0x000 to 0x1FF interpreter
// 0x050 to 0x0A0 font set
// 0x200 to 0xFFF program ROM and ram

var chipFontSet = fontSet{
	0xF0, 0x90, 0x90, 0x90, 0xF0, // 0
	0x20, 0x60, 0x20, 0x20, 0x70, // 1
	0xF0, 0x10, 0xF0, 0x80, 0xF0, // 2
	0xF0, 0x10, 0xF0, 0x10, 0xF0, // 3
	0x90, 0x90, 0xF0, 0x10, 0x10, // 4
	0xF0, 0x80, 0xF0, 0x10, 0xF0, // 5
	0xF0, 0x80, 0xF0, 0x90, 0xF0, // 6
	0xF0, 0x10, 0x20, 0x40, 0x40, // 7
	0xF0, 0x90, 0xF0, 0x90, 0xF0, // 8
	0xF0, 0x90, 0xF0, 0x10, 0xF0, // 9
	0xF0, 0x90, 0xF0, 0x90, 0x90, // A
	0xE0, 0x90, 0xE0, 0x90, 0xE0, // B
	0xF0, 0x80, 0x80, 0x80, 0xF0, // C
	0xE0, 0x90, 0x90, 0x90, 0xE0, // D
	0xF0, 0x80, 0xF0, 0x80, 0xF0, // E
	0xF0, 0x80, 0xF0, 0x80, 0x80, // F
}

type chip8 struct {
	co  currentOpCode
	m   chipMemory
	v   vRegisters
	i   indexRegister
	pc  programCounter
	cg  chipGraphics
	stk stackData
	sp  stackPointer
	kp  chipKeyPad
}

func (c *chip8) initialize() {
	// all values are already intialized to zero
	// only need to set program counter to default location (0x200)
	c.pc = 0x200

	// load fonts
	for i := 0; i < fontSetSize; i++ {
		c.m[i] = chipFontSet[i]
	}
}

func (c *chip8) loadROM(fileName string) {
	bin, err := os.ReadFile(fileName)
	if err != nil {
		fmt.Println("File error: ", err)
		os.Exit(1)
	}

	// Load ROM into memory starting at 0x200
	for i := 0; i < len(bin); i++ {
		c.m[i+int(c.pc)] = bin[0]
	}
}

func main() {
	var c chip8
	c.initialize()
	c.loadROM("binaries/test.ch8")
	// fmt.Printf("c mem: %v\n", c.m)
}
