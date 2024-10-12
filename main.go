package main

import (
	"os"
	"encoding/binary"
	"fmt"
)

type ClassFile struct {
	magic        uint32
	minorVersion uint16
	majorVersion uint16
	constantPoolCount uint16
	// cpInfo       ConstantPool
	// accessFlags  uint16
	// thisClass    uint16
	// superClass   uint16
	// interfaceCount uint16
	// interfaces   []*InterfaceInfo
	// fieldCount   uint16
	// fields       []*FieldInfo
	// methods      []*MethodInfo
	// attributes   []AttributeInfo
}

type ConstantPool []ConstantInfo
type ConstantInfo interface{}
type InterfaceInfo struct {
	index uint16
}

type FieldInfo struct {
	accessFlags     uint16
	nameIndex       uint16
	descriptorIndex uint16
	attributes      []AttributeInfo
}

type MethodInfo struct {
	accessFlags     uint16
	nameIndex       uint16
	descriptorIndex uint16
	attributes      []AttributeInfo
}

type AttributeInfo interface{}

func parseClassFile(data []byte) (cf *ClassFile, err error) {
	return
}

func main() {
	file, err := os.Open("./java/Hello.class")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	classFile := &ClassFile{}

	binary.Read(file, binary.BigEndian, &classFile)

}
