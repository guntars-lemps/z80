package z80

// MemoryAccessor is an interface to access memory addressed by the Z80.
// It defines four read/write method for accessing memory,
type MemoryAccessor interface {
	Read(address uint16) byte
	Write(address uint16, value byte)
	// Data returns the memory content.
	Data() []byte
}

// MemoryReader is a simple interface that defines only a ReadByte method.
// It's used mainly by the disassembler.
type MemoryReader interface {
	Read(address uint16) byte
}
