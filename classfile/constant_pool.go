package classfile

import (
	"fmt"
)

type ConstantPool []ConstantInfo
type ConstantInfo interface{}

// https://docs.oracle.com/javase/specs/jvms/se21/html/jvms-4.html#jvms-4.4-140
const (
	ConstantClassTag              = 7
	ConstantFieldrefTag           = 9
	ConstantMethodrefTag          = 10
	ConstantInterfaceMethodrefTag = 11
	ConstantStringTag             = 8
	ConstantIntegerTag            = 3
	ConstantFloatTag              = 4
	ConstantLongTag               = 5
	ConstantDoubleTag             = 6
	ConstantNameAndTypeTag        = 12
	ConstantUtf8Tag               = 1
	ConstantMethodHandleTag       = 15
	ConstantMethodTypeTag         = 16
	ConstantDynamicTag            = 17
	ConstantInvokeDynamicTag      = 18
	ConstantModuleTag             = 19
	ConstantPackageTag            = 20
)

// https://docs.oracle.com/javase/specs/jvms/se21/html/jvms-4.html#jvms-4.4.1
type ConstantClassInfo struct {
	nameIndex uint16
}
func (self ConstantClassInfo) String() string {
	return fmt.Sprintf("ConstantClassInfo: nameIndex #%d", self.nameIndex)
}

// https://docs.oracle.com/javase/specs/jvms/se21/html/jvms-4.html#jvms-4.4.2
type ConstantFieldrefInfo struct {
	classIndex       uint16
	nameAndTypeIndex uint16
}
func (self ConstantFieldrefInfo) String() string {
	return fmt.Sprintf("ConstantFieldInfo: classIndex #%d, nameAndTypeIndex #%d", self.classIndex, self.nameAndTypeIndex)
}

type ConstantMethodrefInfo struct {
	classIndex       uint16
	nameAndTypeIndex uint16
}
func (self ConstantMethodrefInfo) String() string {
	return fmt.Sprintf("ConstantMethodrefInfo: classIndex #%d, nameAndTypeIndex #%d", self.classIndex, self.nameAndTypeIndex)
}

type ConstantInterfaceMethodrefInfo struct {
	classIndex       uint16
	nameAndTypeIndex uint16
}
func (self ConstantInterfaceMethodrefInfo) String() string {
	return fmt.Sprintf("ConstantInterfaceMethodrefInfo: classIndex #%d, nameAndTypeIndex #%d", self.classIndex, self.nameAndTypeIndex)
}

// https://docs.oracle.com/javase/specs/jvms/se21/html/jvms-4.html#jvms-4.4.3
type ConstantStringInfo struct {
	stringIndex uint16
}
func (self ConstantStringInfo) String() string {
	return fmt.Sprintf("ConstantStringInfo: stringIndex #%d", self.stringIndex)
}

// https://docs.oracle.com/javase/specs/jvms/se21/html/jvms-4.html#jvms-4.4.4
type ConstantIntegerInfo struct {
	bytes []byte
}
func (self ConstantIntegerInfo) String() string {
	return fmt.Sprintf("ConstantIntegerInfo: %d", self.bytes)
}

type ConstantFloatInfo struct {
	bytes []byte
}
func (self ConstantFloatInfo) String() string {
	return fmt.Sprintf("ConstantFloatInfo: %d", self.bytes)
}

// https://docs.oracle.com/javase/specs/jvms/se21/html/jvms-4.html#jvms-4.4.5
type ConstantLongInfo struct {
	highBytes uint32
	lowBytes  uint32
}
func (self ConstantLongInfo) String() string {
	return fmt.Sprintf("ConstantLongInfo: highBytes %d, lowBytes %d", self.highBytes, self.lowBytes)
}

type ConstantDoubleInfo struct {
	highBytes uint32
	lowBytes  uint32
}
func (self ConstantDoubleInfo) String() string {
	return fmt.Sprintf("ConstantDoubleInfo: highBytes %d, lowBytes %d", self.highBytes, self.lowBytes)
}

// https://docs.oracle.com/javase/specs/jvms/se21/html/jvms-4.html#jvms-4.4.6
type ConstantNameAndTypeInfo struct {
	nameIndex       uint16
	descriptorIndex uint16
}
func (self ConstantNameAndTypeInfo) String() string {
	return fmt.Sprintf("ConstantNameAndTypeInfo: nameIndex #%d, descriptorIndex #%d", self.nameIndex, self.descriptorIndex)
}

// https://docs.oracle.com/javase/specs/jvms/se21/html/jvms-4.html#jvms-4.4.7
type ConstantUtf8Info struct {
	length uint16
	bytes  []byte
}
func (self ConstantUtf8Info) String() string {
	return fmt.Sprintf("ConstantUtf8Info: length %d, %s", self.length, string(self.bytes))
}
func (self ConstantUtf8Info) Value() string {
	return string(self.bytes)
}

// https://docs.oracle.com/javase/specs/jvms/se21/html/jvms-4.html#jvms-4.4.8
type ConstantMethodHandleInfo struct {
	referenceKind  uint8
	referenceIndex uint16
}
func (self ConstantMethodHandleInfo) String() string {
	return fmt.Sprintf("ConstantMethodHandleInfo: referenceKind %d, referenceIndex #%d", self.referenceKind, self.referenceIndex)
}

// https://docs.oracle.com/javase/specs/jvms/se21/html/jvms-4.html#jvms-4.4.9
type ConstantMethodTypeInfo struct {
	descriptorIndex uint16
}
func (self ConstantMethodTypeInfo) String() string {
	return fmt.Sprintf("ConstantMethodTypeInfo: descriptorIndex #%d", self.descriptorIndex)
}

// https://docs.oracle.com/javase/specs/jvms/se21/html/jvms-4.html#jvms-4.4.10
type ConstantDynamicInfo struct {
	bootstrapMethodAttrIndex uint16
	nameAndTypeIndex         uint16
}
func (self ConstantDynamicInfo) String() string {
	return fmt.Sprintf("ConstantDynamicInfo: bootstrapMethodAttrIndex #%d, nameAndTypeIndex #%d", self.bootstrapMethodAttrIndex, self.nameAndTypeIndex)
}

type ConstantInvokeDynamicInfo struct {
	bootstrapMethodAttrIndex uint16
	nameAndTypeIndex         uint16
}
func (self ConstantInvokeDynamicInfo) String() string {
	return fmt.Sprintf("ConstantInvokeDynamicInfo: bootstrapMethodAttrIndex #%d, nameAndTypeIndex #%d", self.bootstrapMethodAttrIndex, self.nameAndTypeIndex)
}

// https://docs.oracle.com/javase/specs/jvms/se21/html/jvms-4.html#jvms-4.4.11
type ConstantModuleInfo struct {
	nameIndex uint16
}
func (self ConstantModuleInfo) String() string {
	return fmt.Sprintf("ConstantModuleInfo: nameIndex #%d", self.nameIndex)
}

// https://docs.oracle.com/javase/specs/jvms/se21/html/jvms-4.html#jvms-4.4.12
type ConstantPackageInfo struct {
	nameIndex uint16
}
func (self ConstantPackageInfo) String() string {
	return fmt.Sprintf("ConstantPackageInfo: nameIndex #%d", self.nameIndex)
}
