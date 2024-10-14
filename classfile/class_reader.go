package classfile

import (
	"encoding/binary"
	"io"
)

type ClassReader struct {
	reader io.Reader
}

func newClassReader(reader io.Reader) *ClassReader {
	return &ClassReader{reader}
}

func (self *ClassReader) Read(size uint32) []byte {
	data := make([]byte, size)
	self.reader.Read(data)
	return data
}

func (self *ClassReader) ReadU1() uint8 {
	data := make([]byte, 1)
	self.reader.Read(data)
	return data[0]
}

func (self *ClassReader) ReadU2() uint16 {
	data := make([]byte, 2)
	self.reader.Read(data)
	return binary.BigEndian.Uint16(data)
}

func (self *ClassReader) ReadU4() uint32 {
	data := make([]byte, 4)
	self.reader.Read(data)
	return binary.BigEndian.Uint32(data)
}

func (self *ClassReader) Read4() int32 {
	data := make([]byte, 4)
	self.reader.Read(data)
	return int32(binary.BigEndian.Uint32(data))
}

