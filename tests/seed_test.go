package tests_test

import "github.com/SnareChops/6502-simulator/sim"

func raw(b ...byte) *sim.Model {
	m := sim.NewModel()
	for i := uint16(0); i < uint16(len(b)); i++ {
		m.Set(b[i], sim.AsBytes(i)...)
	}
	return m
}

type seeder = func(*sim.Model) []byte

func seed(op byte, s seeder) *sim.Model {
	m := sim.NewModel()
	a := s(m)
	m.Set(op, 0x0)
	for i := uint16(0); i < uint16(len(a)); i++ {
		m.Set(a[i], sim.AsBytes(i+1)...)
	}
	return m
}

func seedA(op, val byte) (*sim.Model, []byte) {
	a := []byte{0x12, 0xaa}
	m := raw(op, a[0], a[1])
	m.Set(val, a...)
	return m, a
}

func seedAX(op, val byte) (*sim.Model, []byte) {
	a := []byte{0x12, 0xaa}
	m := raw(op, a[0]-0x02, a[1])
	m.Set(val, a...)
	m.X = 0x02
	return m, a
}

func seedAY(op, val byte) (*sim.Model, []byte) {
	a := []byte{0x12, 0xaa}
	m := raw(op, a[0]-0x02, a[1])
	m.Set(val, a...)
	m.Y = 0x02
	return m, a
}

func seedZP(op, val byte) (*sim.Model, byte) {
	a := byte(0xa0)
	m := raw(op, a)
	m.Set(val, a)
	return m, a
}

func seedZPX(op, val byte) (*sim.Model, byte) {
	a := byte(0xa0)
	m := raw(op, a-0x10)
	m.Set(val, a)
	m.X = 0x10
	return m, a
}

func seedZPY(op, val byte) (*sim.Model, byte) {
	a := byte(0xa0)
	m := raw(op, a-0x10)
	m.Set(val, a)
	m.Y = 0x10
	return m, a
}

func seedZPIX(op, val byte) (*sim.Model, []byte) {
	a := []byte{0xb2, 0xb3}
	m := raw(op, a[0]-0x10)
	m.Set(a[0], a[0])
	m.Set(a[1], a[1])
	m.Set(val, a...)
	m.X = 0x10
	return m, a
}

func seedZPIY(op, val byte) (*sim.Model, []byte) {
	a := []byte{0xa2, 0xa1}
	y := byte(0x02)
	m := raw(op, a[0]-y)
	m.Set(a[0]-y, a[0]-y)
	m.Set(a[0]-y+0x01, a[0]-y+0x01)
	m.Set(val, a...)
	m.Y = y
	return m, a
	// return func(m *sim.Model) []byte {
	// 	m.Set(0xa0, 0xb0)
	// 	m.Set(0xaa, 0xb1)
	// 	m.Set(b, 0xa2, 0xaa)
	// 	m.Y = 0x10
	// 	return []byte{0xb0}
	// }
}
