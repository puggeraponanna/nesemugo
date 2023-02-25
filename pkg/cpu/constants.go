package cpu

type OpCode uint8

const (
	LDA   OpCode = 0xA9
	TAX   OpCode = 0xAA
	INX   OpCode = 0xE8
	BREAK OpCode = 0x00
)

const (
	C int = iota
	Z
	I
	D
	V
	N
)
