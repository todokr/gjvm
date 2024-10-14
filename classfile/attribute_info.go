package classfile

import (
	"fmt"
	"strings"
)

type AttributeInfo interface {
	String() string
}

type NotImplementedAttributeInfo struct {
	AttributeName   string
	AttributeLength uint32
	Info            []byte
}

func (self NotImplementedAttributeInfo) String() string {
	return fmt.Sprintf("UnknownAttributeInfo{AttributeName=%s, AttributeLength=%d}", self.AttributeName, self.AttributeLength)
}

// https://docs.oracle.com/javase/specs/jvms/se21/html/jvms-4.html#jvms-4.7
//
// Seven attributes are critical to correct interpretation of the class file by the Java Virtual Machine:
// - ConstantValue
// - Code
// - StackMapTable
// - BootstrapMethods
// - NestHost
// - NestMembers
// - PermittedSubclasses

const (
	ConstantValue string = "ConstantValue"
	Code                 = "Code"
	StackMapTable        = "StackMapTable"
	// Exceptions                                 = "Exceptions"
	// InnerClasses                               = "InnerClasses"
	// EnclosingMethod                            = "EnclosingMethod"
	// Synthetic                                  = "Synthetic"
	// Signature                                  = "Signature"
	// SourceFile                                 = "SourceFile"
	// SourceDebugExtension                       = "SourceDebugExtension"
	// LineNumberTable                            = "LineNumberTable"
	// LocalVariableTable                         = "LocalVariableTable"
	// LocalVariableTypeTable                     = "LocalVariableTypeTable"
	// Deprecated                                 = "Deprecated"
	// RuntimeVisibleAnnotation                   = "RuntimeVisibleAnnotation"
	// RuntimeInvisibleAnnotation                 = "RuntimeInvisibleAnnotation"
	// RuntimeVisibleParameterAnnotation          = "RuntimeVisibleParameterAnnotation"
	// RuntimeInvisibleParameterAnnotation        = "RuntimeInvisibleParameterAnnotation"
	// AnnotationDefault                          = "AnnotationDefault"
	BootstrapMethods = "BootstrapMethods"
	// MethodParameters                           = "MethodParameters"
	// Module                                     = "Module"
	// ModulePackages                             = "ModulePackages"
	// ModuleMainClass                            = "ModuleMainClass"
	NestHost    = "NestHost"
	NestMembers = "NestMembers"
	//	Record                                     = "Record"
	PermittedSubclasses = "PermittedSubclasses"
)

func (self ClassFileParser) parseAttributeInfo(size uint16) ([]AttributeInfo, error) {
	attributes := make([]AttributeInfo, size)
	for i := 0; i < int(size); i++ {
		nameIndex := self.reader.ReadU2()
		attrName := self.cp[nameIndex].(*ConstantUtf8Info).Value()
		attrLen := self.reader.ReadU4()

		switch attrName {
		case ConstantValue:
			attributes[i] = self.parseConstantValueAttribute()
		case Code:
			attributes[i] = self.parseCodeAttribute()
		// case StackMapTable:
		// 	attributes[i] = self.parseStackMapTableAttribute(nameIndex, attrLen)
		// case Exceptions:
		// 	attributes[i] = self.parseExceptionsAttribute(nameIndex, attrLen)
		// case InnerClasses:
		// 	attributes[i] = self.parseInnerClassesAttribute(nameIndex, attrLen)
		// case EnclosingMethod:
		// 	attributes[i] = self.parseEnclosingMethodAttribute(nameIndex, attrLen)
		// case Synthetic:
		// 	attributes[i] = self.parseSyntheticAttribute(nameIndex, attrLen)
		// case Signature:
		// 	attributes[i] = self.parseSignatureAttribute(nameIndex, attrLen)
		// case SourceFile:
		// 	attributes[i] = self.parseSourceFileAttribute(nameIndex, attrLen)
		// case SourceDebugExtension:
		// 	attributes[i] = self.parseSourceDebugExtensionAttribute(nameIndex, attrLen)
		// case LineNumberTable:
		// 	attributes[i] = self.parseLineNumberTableAttribute(nameIndex, attrLen)
		// case LocalVariableTable:
		// 	attributes[i] = self.parseLocalVariableTableAttribute(nameIndex, attrLen)
		// case LocalVariableTypeTable:
		// 	attributes[i] = self.parseLocalVariableTypeTableAttribute(nameIndex, attrLen)
		// case Deprecated:
		// 	attributes[i] = self.parseDeprecatedAttribute(nameIndex, attrLen)
		// case RuntimeVisibleAnnotation:
		// 	attributes[i] = self.parseRuntimeVisibleAnnotationAttribute(nameIndex, attrLen)
		// case RuntimeInvisibleAnnotation:
		// 	attributes[i] = self.parseRuntimeInvisibleAnnotationAttribute(nameIndex, attrLen)
		// case RuntimeVisibleParameterAnnotation:
		// 	attributes[i] = self.parseRuntimeVisibleParameterAnnotationAttribute(nameIndex, attrLen)
		// case RuntimeInvisibleParameterAnnotation:
		// 	attributes[i] = self.parseRuntimeInvisibleParameterAnnotationAttribute(nameIndex, attrLen)
		// case AnnotationDefault:
		// 	attributes[i] = self.parseAnnotationDefaultAttribute(nameIndex, attrLen)
		case BootstrapMethods:
			attributes[i] = self.parseBootstrapMethodsAttribute(nameIndex, attrLen)
		// case MethodParameters:
		// 	attributes[i] = self.parseMethodParametersAttribute(nameIndex, attrLen)
		// case Module:
		// 	attributes[i] = self.parseModuleAttribute(nameIndex, attrLen)
		// case ModulePackages:
		// 	attributes[i] = self.parseModulePackagesAttribute(nameIndex, attrLen)
		// case ModuleMainClass:
		// 	attributes[i] = self.parseModuleMainClassAttribute(nameIndex, attrLen)
		// case NestHost:
		// 	attributes[i] = self.parseNestHostAttribute(nameIndex, attrLen)
		// case NestMembers:
		// 	attributes[i] = self.parseNestMembersAttribute(nameIndex, attrLen)
		// case Record:
		// 	attributes[i] = self.parseRecordAttribute(nameIndex, attrLen)
		// case PermittedSubclasses:
		// 	attributes[i] = self.parsePermittedSubclassesAttribute(nameIndex, attrLen)
		default:
			attributes[i] = NotImplementedAttributeInfo{attrName, attrLen, self.reader.Read(attrLen)}
		}
	}
	return attributes, nil
}

// https://docs.oracle.com/javase/specs/jvms/se21/html/jvms-4.html#jvms-4.7.2
//
//	ConstantValue_attribute {
//	    u2 attribute_name_index;
//	    u4 attribute_length;
//	    u2 constantvalue_index;
//	}
type ConstantValueAttribute struct {
	ConstantValueIndex uint16
}

func (self ConstantValueAttribute) String() string {
	return fmt.Sprintf("ConstantValue: #%d", self.ConstantValueIndex)
}

func (self ClassFileParser) parseConstantValueAttribute() ConstantValueAttribute {
	return ConstantValueAttribute{self.reader.ReadU2()}
}

// https://docs.oracle.com/javase/specs/jvms/se21/html/jvms-4.html#jvms-4.7.3
//
//	Code_attribute {
//	    u2 attribute_name_index;
//	    u4 attribute_length;
//	    u2 max_stack;
//	    u2 max_locals;
//	    u4 code_length;
//	    u1 code[code_length];
//	    u2 exception_table_length;
//	    {   u2 start_pc;
//	        u2 end_pc;
//	        u2 handler_pc;
//	        u2 catch_type;
//	    } exception_table[exception_table_length];
//	    u2 attributes_count;
//	    attribute_info attributes[attributes_count];
//	}
type CodeAttribute struct {
	MaxStack             uint16
	MaxLocals            uint16
	CodeLength           uint32
	Code                 []byte
	ExceptionTableLength uint16
	ExceptionTable       []ExceptionTableEntry
	AttributeCount       uint16
	Attributes           []AttributeInfo
}

func (self CodeAttribute) String() string {
	codes := []string{}
	for _, code := range self.Code {
		codes = append(codes, fmt.Sprintf("%#x", code))
	}
	s := fmt.Sprintf("Code: max_stack=%d, max_locals=%d, code_length=%d, code=%s, exception_table_length=%d, attributes_count=%d", self.MaxStack, self.MaxLocals, self.CodeLength, codes, self.ExceptionTableLength, self.AttributeCount)
	attr := []string{}
	for _, a := range self.Attributes {
		attr = append(attr, a.String())
	}
	if len(attr) > 0 {
		s = s + "\n  " + strings.Join(attr, "\n  ")
	}
	return s
}

type ExceptionTableEntry struct {
	StartPc   uint16
	EndPc     uint16
	HandlerPc uint16
	CatchType uint16
}

func (self ClassFileParser) parseCodeAttribute() CodeAttribute {
	maxStack := self.reader.ReadU2()
	maxLocals := self.reader.ReadU2()
	codeLength := self.reader.ReadU4()
	code := self.reader.Read(codeLength)
	exceptionTableLength := self.reader.ReadU2()
	exceptionTable := make([]ExceptionTableEntry, exceptionTableLength)
	for i := range exceptionTable {
		exceptionTable[i] = ExceptionTableEntry{
			StartPc:   self.reader.ReadU2(),
			EndPc:     self.reader.ReadU2(),
			HandlerPc: self.reader.ReadU2(),
			CatchType: self.reader.ReadU2(),
		}
	}
	attributesCount := self.reader.ReadU2()
	attributes, _ := self.parseAttributeInfo(attributesCount)
	return CodeAttribute{
		MaxStack:             maxStack,
		MaxLocals:            maxLocals,
		CodeLength:           codeLength,
		Code:                 code,
		ExceptionTableLength: exceptionTableLength,
		ExceptionTable:       exceptionTable,
		AttributeCount:       attributesCount,
		Attributes:           attributes,
	}
}

// https://docs.oracle.com/javase/specs/jvms/se21/html/jvms-4.html#jvms-4.7.4
// StackMapTable_attribute {
//     u2              attribute_name_index;
//     u4              attribute_length;
//     u2              number_of_entries;
//     stack_map_frame entries[number_of_entries];
// }
// type StackMapTableAttribute struct {
// 	NumberOfEntries uint16
// 	Entries         []StackMapFrame
// }
// type StackMapFrame interface {
// 	FrameType() uint8
// }

// https://docs.oracle.com/javase/specs/jvms/se21/html/jvms-4.html#jvms-4.7.23
//
//	BootstrapMethods_attribute {
//	  u2 attribute_name_index;
//	  u4 attribute_length;
//	  u2 num_bootstrap_methods;
//	  {   u2 bootstrap_method_ref;
//	      u2 num_bootstrap_arguments;
//	      u2 bootstrap_arguments[num_bootstrap_arguments];
//	  } bootstrap_methods[num_bootstrap_methods];
//	}
type BootstrapMethodsAttribute struct {
	BootstrapMethods []BootstrapMethod
}

func (self BootstrapMethodsAttribute) String() string {
	methods := []string{}
	for _, method := range self.BootstrapMethods {
		methods = append(methods, method.String())
	}
	return fmt.Sprintf("BootstrapMethods:\n %s", strings.Join(methods, "\n"))
}

type BootstrapMethod struct {
	BootstrapMethodRef uint16
	BootstrapArguments []uint16
}

func (self BootstrapMethod) String() string {
	args := []string{}
	for _, arg := range self.BootstrapArguments {
		args = append(args, fmt.Sprintf("#%d", arg))
	}
	return fmt.Sprintf("BootStrapMethod: #%d, args=%s", self.BootstrapMethodRef, strings.Join(args, ", "))
}

func (self ClassFileParser) parseBootstrapMethodsAttribute(nameIndex uint16, attrLen uint32) BootstrapMethodsAttribute {
	methodNum := self.reader.ReadU2()
	methods := make([]BootstrapMethod, methodNum)
	for i := range methods {
		methodRef := self.reader.ReadU2()
		argNum := self.reader.ReadU2()
		args := make([]uint16, argNum)
		for j := range args {
			args[j] = self.reader.ReadU2()
		}
		methods[i] = BootstrapMethod{
			BootstrapMethodRef: methodRef,
			BootstrapArguments: args,
		}
	}
	return BootstrapMethodsAttribute{methods}
}
