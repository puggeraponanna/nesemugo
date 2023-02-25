package cpu_test

import (
	"nesemugo/pkg/cpu"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_LDA_IMM(t *testing.T) {
	cpu := cpu.NewCPU()
	cpu.Interpret([]uint8{0xa9, 0x05, 0x00})

	assert.EqualValues(t, 0x05, cpu.Register_A)
	assert.EqualValues(t, 0b0000_0000, cpu.Status&0b0000_0010)
	assert.EqualValues(t, 0b0000_0000, cpu.Status&0b1000_0000)
}

func Test_LDA_ZF(t *testing.T) {
	cpu := cpu.NewCPU()
	cpu.Interpret([]uint8{0xa9, 0x00, 0x00})

	assert.EqualValues(t, 0b0000_0010, cpu.Status&0b0000_0010)
}

func Test_LDA_NF(t *testing.T) {
	cpu := cpu.NewCPU()
	cpu.Interpret([]uint8{0xa9, 0x82, 0x00})

	assert.EqualValues(t, 0b1000_0000, cpu.Status&0b1000_0000)
}

func Test_TAX(t *testing.T) {
	cpu := cpu.NewCPU()
	cpu.Register_A = 0x0A
	cpu.Interpret([]uint8{0xAA, 0x00})

	assert.EqualValues(t, 0x0A, cpu.Register_X)
	assert.EqualValues(t, 0b0000_0000, cpu.Status&0b0000_0010)
	assert.EqualValues(t, 0b0000_0000, cpu.Status&0b1000_0000)
}

func Test_TAX_ZF(t *testing.T) {
	cpu := cpu.NewCPU()
	cpu.Register_A = 0x00
	cpu.Interpret([]uint8{0xAA, 0x00})

	assert.EqualValues(t, 0b0000_0010, cpu.Status&0b0000_0010)
}

func Test_TAX_NF(t *testing.T) {
	cpu := cpu.NewCPU()
	cpu.Register_A = 0x80
	cpu.Interpret([]uint8{0xAA, 0x00})

	assert.EqualValues(t, 0b1000_0000, cpu.Status&0b1000_0000)
}

func Test_INX(t *testing.T) {
	cpu := cpu.NewCPU()
	cpu.Register_X = 0x01
	cpu.Interpret([]uint8{0xE8, 0x00})

	assert.EqualValues(t, 0x02, cpu.Register_X)
}

func Test_INX_ZFVF(t *testing.T) {
	cpu := cpu.NewCPU()
	cpu.Register_X = 0xFF
	cpu.Interpret([]uint8{0xE8, 0x00})

	assert.EqualValues(t, 0b0100_0000, cpu.Status&0b0100_0000)
	assert.EqualValues(t, 0b0000_0010, cpu.Status&0b0000_0010)
}
