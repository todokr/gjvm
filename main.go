package main

import (
	"fmt"
	"os"

	"gjvm/classfile"
)

func main() {
	file, err := os.Open("./java/Hello.class")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	//	classFile := &ClassFile{}

	//	binary.Read(file, binary.BigEndian, &classFile)
	parser := classfile.NewClassFileParser(file)
	class, err := parser.Parse()
	if err != nil {
		panic(err)
	}

	fmt.Printf("minorVersion: %#04d\n", class.MinorVersion)
	fmt.Printf("majorVersion: %#04d\n", class.MajorVersion)
	fmt.Printf("constantPoolCount: %d\n", class.ConstantPoolCount)
	for i := 1; i < int(class.ConstantPoolCount); i++ {
		fmt.Printf("cp[%d]: %s\n", i, class.ConstantPool[i])
	}
	fmt.Printf("accessFlags: %s\n", class.AccessFlags)
	
	fmt.Printf("thisClass: #%d\n", class.ThisClass)
	fmt.Printf("superClass: #%d\n", class.SuperClass)
	fmt.Printf("interfacesCount: %d\n", class.InterfacesCount)
	for i := 0; i < int(class.InterfacesCount); i++ {
		fmt.Printf("interface[%d]: #%d\n", i, class.Interfaces[i])
	}
	fmt.Printf("fieldCount: %d\n", class.FieldCount)
	for i := 0; i < int(class.FieldCount); i++ {
		fmt.Printf("field[%d]: %s\n", i, class.Fields[i])
	}
	fmt.Printf("methodCount: %d\n", class.MethodCount)
	for i := 0; i < int(class.MethodCount); i++ {
		fmt.Printf("method[%d]: %s\n", i, class.Methods[i])
	}
}
