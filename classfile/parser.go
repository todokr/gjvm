package classfile

import (
	"fmt"
	"io"
)

type ClassFileParser struct {
	reader *ClassReader
	cp     []ConstantInfo
}

func NewClassFileParser(reader io.Reader) *ClassFileParser {
	r := &ClassReader{reader}
	return &ClassFileParser{r, nil}
}

func (self *ClassFileParser) Parse() (*ClassFile, error) {
	magic := self.reader.ReadU4()
	if !validateMagic(magic) {
		return nil, fmt.Errorf("java.lang.ClassFormatError: %d", magic)
	}
	cf := &ClassFile{}

	cf.MinorVersion = self.reader.ReadU2()
	cf.MajorVersion = self.reader.ReadU2()
	cf.ConstantPoolCount = self.reader.ReadU2()
	cp, err := self.parseConstantPool(cf.ConstantPoolCount)
	if err != nil {
		return nil, err
	}
	cf.ConstantPool = cp
	self.cp = cp
	cf.AccessFlags = AccessFlags(self.reader.ReadU2())
	cf.ThisClass = self.reader.ReadU2()
	cf.SuperClass = self.reader.ReadU2()
	cf.InterfacesCount = self.reader.ReadU2()
	var ifs []uint16
	for i := 0; i < int(cf.InterfacesCount); i++ {
		ifs = append(ifs, self.reader.ReadU2())
	}
	cf.Interfaces = ifs
	cf.FieldCount = self.reader.ReadU2()
	fields, err := self.parseFiledInfo(cf.FieldCount)
	if err != nil {
		return nil, err
	}
	cf.Fields = fields
	cf.MethodCount = self.reader.ReadU2()
	methods, err := self.parseMethodInfo(cf.MethodCount)
	if err != nil {
		return nil, err
	}
	cf.Methods = methods

	return cf, nil
}

func validateMagic(magic uint32) bool {
	return magic == 0xCAFEBABE
}

func (self ClassFileParser) parseConstantPool(count uint16) ([]ConstantInfo, error) {
	constantPool := make([]ConstantInfo, count)
	for i := 1; i < int(count); i++ {
		tag := self.reader.ReadU1()
		switch tag {
		case ConstantClassTag:
			nameIndex := self.reader.ReadU2()
			constantPool[i] = &ConstantClassInfo{nameIndex}
		case ConstantFieldrefTag:
			classIndex := self.reader.ReadU2()
			nameAndTypeIndex := self.reader.ReadU2()
			constantPool[i] = &ConstantFieldrefInfo{classIndex, nameAndTypeIndex}
		case ConstantMethodrefTag:
			classIndex := self.reader.ReadU2()
			nameAndTypeIndex := self.reader.ReadU2()
			constantPool[i] = &ConstantMethodrefInfo{classIndex, nameAndTypeIndex}
		case ConstantInterfaceMethodrefTag:
			classIndex := self.reader.ReadU2()
			nameAndTypeIndex := self.reader.ReadU2()
			constantPool[i] = &ConstantInterfaceMethodrefInfo{classIndex, nameAndTypeIndex}
		case ConstantStringTag:
			stringIndex := self.reader.ReadU2()
			constantPool[i] = &ConstantStringInfo{stringIndex}
		case ConstantIntegerTag:
			b := self.reader.Read(4)
			constantPool[i] = &ConstantIntegerInfo{b}
		case ConstantFloatTag:
			b := self.reader.Read(4)
			constantPool[i] = &ConstantFloatInfo{b}
		case ConstantLongTag:
			highBytes := self.reader.ReadU4()
			lowBytes := self.reader.ReadU4()
			constantPool[i] = &ConstantLongInfo{highBytes, lowBytes}
		case ConstantDoubleTag:
			highBytes := self.reader.ReadU4()
			lowBytes := self.reader.ReadU4()
			constantPool[i] = &ConstantDoubleInfo{highBytes, lowBytes}
		case ConstantNameAndTypeTag:
			nameIndex := self.reader.ReadU2()
			descriptorIndex := self.reader.ReadU2()
			constantPool[i] = &ConstantNameAndTypeInfo{nameIndex, descriptorIndex}
		case ConstantUtf8Tag:
			length := self.reader.ReadU2()
			bytes := self.reader.Read(uint32(length))
			constantPool[i] = &ConstantUtf8Info{length, bytes}
		case ConstantMethodHandleTag:
			referenceKind := self.reader.ReadU1()
			referenceIndex := self.reader.ReadU2()
			constantPool[i] = &ConstantMethodHandleInfo{referenceKind, referenceIndex}
		case ConstantMethodTypeTag:
			descriptorIndex := self.reader.ReadU2()
			constantPool[i] = &ConstantMethodTypeInfo{descriptorIndex}
		case ConstantDynamicTag:
			bootstrapMethodAttrIndex := self.reader.ReadU2()
			nameAndTypeIndex := self.reader.ReadU2()
			constantPool[i] = &ConstantDynamicInfo{bootstrapMethodAttrIndex, nameAndTypeIndex}
		case ConstantInvokeDynamicTag:
			bootstrapMethodAttrIndex := self.reader.ReadU2()
			nameAndTypeIndex := self.reader.ReadU2()
			constantPool[i] = &ConstantInvokeDynamicInfo{bootstrapMethodAttrIndex, nameAndTypeIndex}
		case ConstantModuleTag:
			nameIndex := self.reader.ReadU2()
			constantPool[i] = &ConstantModuleInfo{nameIndex}
		case ConstantPackageTag:
			nameIndex := self.reader.ReadU2()
			constantPool[i] = &ConstantPackageInfo{nameIndex}
		default:
			return nil, fmt.Errorf("java.lang.ClassFormatError: constant pool tag %d", tag)
		}
	}
	return constantPool, nil
}

func (self ClassFileParser) parseFiledInfo(size uint16) ([]FieldInfo, error) {
	fields := make([]FieldInfo, size)
	for i := 0; i < int(size); i++ {
		accessFlags := AccessFlags(self.reader.ReadU2())
		nameIndex := self.reader.ReadU2()
		descriptorIndex := self.reader.ReadU2()
		attributesCount := self.reader.ReadU2()
		attributes, err := self.parseAttributeInfo(attributesCount)
		if err != nil {
			return nil, err
		}
		fields[i] = FieldInfo{accessFlags, nameIndex, descriptorIndex, attributes}
	}
	return fields, nil
}

func (self ClassFileParser) parseMethodInfo(size uint16) ([]MethodInfo, error) {
	methods := make([]MethodInfo, size)
	for i := 0; i < int(size); i++ {
		accessFlags := AccessFlags(self.reader.ReadU2())
		nameIndex := self.reader.ReadU2()
		name := self.cp[nameIndex].(*ConstantUtf8Info).Value()
		descriptorIndex := self.reader.ReadU2()
		descriptor := self.cp[descriptorIndex].(*ConstantUtf8Info).Value()
		attributesCount := self.reader.ReadU2()
		attributes, err := self.parseAttributeInfo(attributesCount)
		
		code := []byte{}
		for _, attr := range attributes {
			switch attr.(type) {
			case CodeAttribute:
				code = attr.(CodeAttribute).Code
			}
		}
		if err != nil {
			return nil, err
		}
		methods[i] = MethodInfo{accessFlags, name, nameIndex, descriptor, descriptorIndex, attributes, code}
	}
	return methods, nil
}
