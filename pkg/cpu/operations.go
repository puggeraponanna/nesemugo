package cpu

func (cpu *CPU) LDA(val uint8) {
	cpu.Register_A = val

	if cpu.Register_A == 0 {
		cpu.SetFlag(Z)
	} else {
		cpu.ResetFlag(Z)
	}

	if cpu.Register_A&0b1000_0000 != 0 {
		cpu.SetFlag(N)
	} else {
		cpu.ResetFlag(N)
	}

	cpu.PC += 1
}

func (cpu *CPU) TAX() {
	cpu.Register_X = cpu.Register_A

	if cpu.Register_X == 0 {
		cpu.SetFlag(Z)
	} else {
		cpu.ResetFlag(Z)
	}

	if cpu.Register_X&0b1000_0000 != 0 {
		cpu.SetFlag(N)
	} else {
		cpu.ResetFlag(N)
	}
}

func (cpu *CPU) INX() {
	if cpu.Register_X == 0xFF {
		cpu.SetFlag(V)
	} else {
		cpu.ResetFlag(V)
	}
	cpu.Register_X += 1

	if cpu.Register_X == 0 {
		cpu.SetFlag(Z)
	} else {
		cpu.ResetFlag(Z)
	}

	if cpu.Register_X&0b1000_0000 != 0 {
		cpu.SetFlag(N)
	} else {
		cpu.ResetFlag(N)
	}
}
