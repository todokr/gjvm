package classfile

import (
	"fmt"
	"strings"
)

// field_info {
//     u2             access_flags;
//     u2             name_index;
//     u2             descriptor_index;
//     u2             attributes_count;
//     attribute_info attributes[attributes_count];
// }
type FieldInfo struct {
	AccessFlags
	NameIndex       uint16
	DescriptorIndex uint16
	Attributes      []AttributeInfo
}
func (self FieldInfo) String() string {
	s := fmt.Sprintf("FieldInfo: accessFlags %v, nameIndex #%d, descriptorIndex #%d,", self.AccessFlags, self.NameIndex, self.DescriptorIndex)
	attrs := []string{}
	for _, attr := range self.Attributes {
		attrs = append(attrs, attr.String())
	}
	if len(attrs) > 0 {
		s = s + "\n  " +  strings.Join(attrs, " ")
	}
	return s
}

// method_info {
//     u2             access_flags;
//     u2             name_index;
//     u2             descriptor_index;
//     u2             attributes_count;
//     attribute_info attributes[attributes_count];
// }
type MethodInfo struct {
	AccessFlags
	NameIndex       uint16
	DescriptorIndex uint16
	Attributes      []AttributeInfo
}
func (self MethodInfo) String() string {
	s := fmt.Sprintf("MethodInfo: accessFlags %v, nameIndex #%d, descriptorIndex #%d,", self.AccessFlags, self.NameIndex, self.DescriptorIndex)
	attrs := []string{}
	for _, attr := range self.Attributes {
		attrs = append(attrs, attr.String())
	}
	if len(attrs) > 0 {
		s = s + "\n  " +  strings.Join(attrs, " ")
	}
	return s
}
