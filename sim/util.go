package sim

import "encoding/binary"

// AsUint16 converts a little endian
// []byte to uint16
func AsUint16(b ...byte) uint16 {
	if len(b) == 1 {
		return uint16(b[0])
	}
	return binary.LittleEndian.Uint16(b)
}

// AsBytes converts a uint16 to little endian
// []byte
func AsBytes(n uint16) []byte {
	result := make([]byte, 2)
	binary.LittleEndian.PutUint16(result, n)
	return result
}

// Int8Overflows returns true if the
// result of the inputs results in
// signed overflow
func Int8Overflows(start, end int8) bool {
	if start > 0 && end < 0 {
		return true
	}
	if start < 0 && end >= 0 {
		return true
	}
	return false
}
