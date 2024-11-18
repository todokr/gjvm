package classfile

import (
	"fmt"
	"strings"
)

//	field_info {
//	    u2             access_flags;
//	    u2             name_index;
//	    u2             descriptor_index;
//	    u2             attributes_count;
//	    attribute_info attributes[attributes_count];
//	}
type FieldInfo struct {
	AccessFlags
	NameIndex       uint16
	DescriptorIndex uint16
	Attributes      []AttributeInfo
}

func (f FieldInfo) String() string {
	s := fmt.Sprintf("FieldInfo: accessFlags %v, nameIndex #%d, descriptorIndex #%d,", f.AccessFlags, f.NameIndex, f.DescriptorIndex)
	attrs := []string{}
	for _, attr := range f.Attributes {
		attrs = append(attrs, attr.String())
	}
	if len(attrs) > 0 {
		s = s + "\n  " + strings.Join(attrs, " ")
	}
	return s
}

//	method_info {
//	    u2             access_flags;
//	    u2             name_index;
//	    u2             descriptor_index;
//	    u2             attributes_count;
//	    attribute_info attributes[attributes_count];
//	}
type MethodInfo struct {
	AccessFlags
	Name            string
	NameIndex       uint16
	Descriptor      string
	DescriptorIndex uint16
	Attributes      []AttributeInfo
	Code            []byte
}

func (m MethodInfo) String() string {
	s := fmt.Sprintf("MethodInfo: accessFlags %v, name %s, descriptor %s\n", m.AccessFlags, m.Name, m.Descriptor)
	s = s + fmt.Sprintf("code %v", m.Code)
	attrs := []string{}
	for _, attr := range m.Attributes {
		attrs = append(attrs, attr.String())
	}
	if len(attrs) > 0 {
		s = s + "\n  " + strings.Join(attrs, " ")
	}
	return s
}
