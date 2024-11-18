package classfile

import (
	"encoding/binary"
	"io"
)

/// A Reader which has some convenience methods for reading Java class files
type ClassReader struct {
	reader io.Reader
}

func newClassReader(r io.Reader) *ClassReader {
	return &ClassReader{r}
}

func (cr *ClassReader) Read(size uint32) []byte {
	data := make([]byte, size)
	cr.reader.Read(data)
	return data
}

func (cr *ClassReader) ReadU1() uint8 {
	data := make([]byte, 1)
	cr.reader.Read(data)
	return data[0]
}

func (cr *ClassReader) ReadU2() uint16 {
	data := make([]byte, 2)
	cr.reader.Read(data)
	return binary.BigEndian.Uint16(data)
}

func (cr *ClassReader) ReadU4() uint32 {
	data := make([]byte, 4)
	cr.reader.Read(data)
	return binary.BigEndian.Uint32(data)
}

func (cr *ClassReader) Read4() int32 {
	data := make([]byte, 4)
	cr.reader.Read(data)
	return int32(binary.BigEndian.Uint32(data))
}

