package cpu

import (
	"log"
)

type CPU struct {
	Register_A uint8
	Register_X uint8
	Status     uint8
	PC         uint16
}

func NewCPU() *CPU {
	return &CPU{
		Register_A: 0,
		Status:     0,
		PC:         0,
	}
}

func (cpu *CPU) Interpret(pgm []uint8) {
	cpu.PC = 0

	for {
		opcode := pgm[cpu.PC]
		cpu.PC += 1
		switch opcode {
		case uint8(LDA):
			param := pgm[cpu.PC]
			cpu.LDA(param)

		case uint8(TAX):
			cpu.TAX()

		case uint8(INX):
			cpu.INX()

		case uint8(BREAK):
			return

		default:
			log.Fatal("Invalid Flag")
		}
	}
}

func (cpu *CPU) SetFlag(flag int) {
	switch flag {
	case C:
		cpu.Status = cpu.Status | 0b0000_0001
	case Z:
		cpu.Status = cpu.Status | 0b0000_0010
	case I:
		cpu.Status = cpu.Status | 0b0000_0100
	case D:
		cpu.Status = cpu.Status | 0b0000_1000
	case V:
		cpu.Status = cpu.Status | 0b0100_0000
	case N:
		cpu.Status = cpu.Status | 0b1000_0000
	default:
		log.Fatal("Invalid Flag")
	}
}

func (cpu *CPU) ResetFlag(flag int) {
	switch flag {
	case C:
		cpu.Status = cpu.Status & 0b1111_1110
	case Z:
		cpu.Status = cpu.Status & 0b1111_1101
	case I:
		cpu.Status = cpu.Status & 0b1111_1011
	case D:
		cpu.Status = cpu.Status & 0b1111_0111
	case V:
		cpu.Status = cpu.Status & 0b1011_1111
	case N:
		cpu.Status = cpu.Status & 0b0111_1111
	default:
		log.Fatal("Invalid Flag")
	}
}
