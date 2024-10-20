package classfile

import (
	"fmt"
	"strings"
	"regexp"
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
	NameIndex uint16
}
func (self ConstantClassInfo) String() string {
	return fmt.Sprintf("ConstantClassInfo: nameIndex #%d", self.NameIndex)
}

// https://docs.oracle.com/javase/specs/jvms/se21/html/jvms-4.html#jvms-4.4.2
type ConstantFieldrefInfo struct {
	ClassIndex       uint16
	NameAndTypeIndex uint16
}
func (self ConstantFieldrefInfo) String() string {
	return fmt.Sprintf("ConstantFieldInfo: classIndex #%d, nameAndTypeIndex #%d", self.ClassIndex, self.NameAndTypeIndex)
}
func (self ConstantFieldrefInfo) Resolve(cp ConstantPool) FieldRef {
	class := cp[self.ClassIndex].(*ConstantClassInfo)
	className := cp[class.NameIndex].(*ConstantUtf8Info).Value()
	className = strings.ReplaceAll(className, "/", ".")

	nameAndType := cp[self.NameAndTypeIndex].(*ConstantNameAndTypeInfo)
	name := cp[nameAndType.NameIndex].(*ConstantUtf8Info).Value()
	descriptor := cp[nameAndType.DescriptorIndex].(*ConstantUtf8Info).Value()
	return FieldRef{className, name, descriptor}
}

type FieldRef struct {
	Class string
	Name string
	Descriptor string
}

type ConstantMethodrefInfo struct {
	ClassIndex       uint16
	NameAndTypeIndex uint16
}
func (self ConstantMethodrefInfo) String() string {
	return fmt.Sprintf("ConstantMethodrefInfo: classIndex #%d, nameAndTypeIndex #%d", self.ClassIndex, self.NameAndTypeIndex)
}

// https://docs.oracle.com/javase/specs/jvms/se21/html/jvms-4.html#jvms-4.3.2
const ArgPattern = `L([^;]+);`
var argRegexp = regexp.MustCompile(ArgPattern)
func (self ConstantMethodrefInfo) Resolve(cp ConstantPool) MethodRef {
	class := cp[self.ClassIndex].(*ConstantClassInfo)
	className := cp[class.NameIndex].(*ConstantUtf8Info).Value()
	className = strings.ReplaceAll(className, "/", ".")
	nameAndType := cp[self.NameAndTypeIndex].(*ConstantNameAndTypeInfo)
	name := cp[nameAndType.NameIndex].(*ConstantUtf8Info).Value()
	descriptor := cp[nameAndType.DescriptorIndex].(*ConstantUtf8Info).Value()
	args := []string{}
	for _, match := range argRegexp.FindAllStringSubmatch(descriptor, -1) {
		arg := strings.ReplaceAll(match[1], "/", ".")
		args = append(args, arg)
	}

	return MethodRef{className, name, args}
}
type MethodRef struct {
	Class string
	Name string
	ArgTypes []string
}

type ConstantInterfaceMethodrefInfo struct {
	ClassIndex       uint16
	NameAndTypeIndex uint16
}

func (self ConstantInterfaceMethodrefInfo) String() string {
	return fmt.Sprintf("ConstantInterfaceMethodrefInfo: classIndex #%d, nameAndTypeIndex #%d", self.ClassIndex, self.NameAndTypeIndex)
}

// https://docs.oracle.com/javase/specs/jvms/se21/html/jvms-4.html#jvms-4.4.3
type ConstantStringInfo struct {
	StringIndex uint16
}
func (self ConstantStringInfo) String() string {
	return fmt.Sprintf("ConstantStringInfo: stringIndex #%d", self.StringIndex)
}

func (self ConstantStringInfo) Resolve(cp ConstantPool) string {
	return cp[self.StringIndex].(*ConstantUtf8Info).Value()
}

// https://docs.oracle.com/javase/specs/jvms/se21/html/jvms-4.html#jvms-4.4.4
type ConstantIntegerInfo struct {
	Bytes []byte
}
func (self ConstantIntegerInfo) String() string {
	return fmt.Sprintf("ConstantIntegerInfo: %d", self.Bytes)
}

type ConstantFloatInfo struct {
	Bytes []byte
}
func (self ConstantFloatInfo) String() string {
	return fmt.Sprintf("ConstantFloatInfo: %d", self.Bytes)
}

// https://docs.oracle.com/javase/specs/jvms/se21/html/jvms-4.html#jvms-4.4.5
type ConstantLongInfo struct {
	HighBytes uint32
	LowBytes  uint32
}
func (self ConstantLongInfo) String() string {
	return fmt.Sprintf("ConstantLongInfo: highBytes %d, lowBytes %d", self.HighBytes, self.LowBytes)
}

type ConstantDoubleInfo struct {
	HighBytes uint32
	LowBytes  uint32
}
func (self ConstantDoubleInfo) String() string {
	return fmt.Sprintf("ConstantDoubleInfo: highBytes %d, lowBytes %d", self.HighBytes, self.LowBytes)
}

// https://docs.oracle.com/javase/specs/jvms/se21/html/jvms-4.html#jvms-4.4.6
type ConstantNameAndTypeInfo struct {
	NameIndex       uint16
	DescriptorIndex uint16
}
func (self ConstantNameAndTypeInfo) String() string {
	return fmt.Sprintf("ConstantNameAndTypeInfo: nameIndex #%d, descriptorIndex #%d", self.NameIndex, self.DescriptorIndex)
}

// https://docs.oracle.com/javase/specs/jvms/se21/html/jvms-4.html#jvms-4.4.7
type ConstantUtf8Info struct {
	Length uint16
	Bytes  []byte
}
func (self ConstantUtf8Info) String() string {
	return fmt.Sprintf("ConstantUtf8Info: length %d, %s", self.Length, string(self.Bytes))
}
func (self ConstantUtf8Info) Value() string {
	return string(self.Bytes)
}

// https://docs.oracle.com/javase/specs/jvms/se21/html/jvms-4.html#jvms-4.4.8
type ConstantMethodHandleInfo struct {
	ReferenceKind  uint8
	ReferenceIndex uint16
}
func (self ConstantMethodHandleInfo) String() string {
	return fmt.Sprintf("ConstantMethodHandleInfo: referenceKind %d, referenceIndex #%d", self.ReferenceKind, self.ReferenceIndex)
}

// https://docs.oracle.com/javase/specs/jvms/se21/html/jvms-4.html#jvms-4.4.9
type ConstantMethodTypeInfo struct {
	DescriptorIndex uint16
}
func (self ConstantMethodTypeInfo) String() string {
	return fmt.Sprintf("ConstantMethodTypeInfo: descriptorIndex #%d", self.DescriptorIndex)
}

// https://docs.oracle.com/javase/specs/jvms/se21/html/jvms-4.html#jvms-4.4.10
type ConstantDynamicInfo struct {
	BootstrapMethodAttrIndex uint16
	NameAndTypeIndex         uint16
}
func (self ConstantDynamicInfo) String() string {
	return fmt.Sprintf("ConstantDynamicInfo: bootstrapMethodAttrIndex #%d, nameAndTypeIndex #%d", self.BootstrapMethodAttrIndex, self.NameAndTypeIndex)
}

type ConstantInvokeDynamicInfo struct {
	BootstrapMethodAttrIndex uint16
	NameAndTypeIndex         uint16
}
func (self ConstantInvokeDynamicInfo) String() string {
	return fmt.Sprintf("ConstantInvokeDynamicInfo: bootstrapMethodAttrIndex #%d, nameAndTypeIndex #%d", self.BootstrapMethodAttrIndex, self.NameAndTypeIndex)
}

// https://docs.oracle.com/javase/specs/jvms/se21/html/jvms-4.html#jvms-4.4.11
type ConstantModuleInfo struct {
	NameIndex uint16
}
func (self ConstantModuleInfo) String() string {
	return fmt.Sprintf("ConstantModuleInfo: nameIndex #%d", self.NameIndex)
}

// https://docs.oracle.com/javase/specs/jvms/se21/html/jvms-4.html#jvms-4.4.12
type ConstantPackageInfo struct {
	NameIndex uint16
}
func (self ConstantPackageInfo) String() string {
	return fmt.Sprintf("ConstantPackageInfo: nameIndex #%d", self.NameIndex)
}
