package classfile

// ClassFile {
//   u4             magic;
//   u2             minor_version;
//   u2             major_version;
//   u2             constant_pool_count;
//   cp_info        constant_pool[constant_pool_count-1];
//   u2             access_flags;
//   u2             this_class;
//   u2             super_class;
//   u2             interfaces_count;
//   u2             interfaces[interfaces_count];
//   u2             fields_count;
//   field_info     fields[fields_count];
//   u2             methods_count;
//   method_info    methods[methods_count];
//   u2             attributes_count;
//   attribute_info attributes[attributes_count];
// }

type ClassFile struct {
	//magic      uint32
	MinorVersion      uint16
	MajorVersion      uint16
	ConstantPoolCount uint16
	ConstantPool      ConstantPool
	AccessFlags       AccessFlags
	ThisClass         uint16
	SuperClass        uint16
	InterfacesCount   uint16
	Interfaces        []uint16
	FieldCount        uint16
	Fields            []FieldInfo
	MethodCount       uint16
	Methods           []MethodInfo
	//	AttributeTable
}
