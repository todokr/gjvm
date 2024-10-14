package classfile

import (
	"strings"
	"fmt"
)

type AccessFlags uint16
func (self AccessFlags) String() string {
	flgs := []string{}
	if self.IsPublic() {
		flgs = append(flgs, "public")
	}
	if self.IsFinal() {
		flgs = append(flgs, "final")
	}
	if self.IsSuper() {
		flgs = append(flgs, "super")
	}
	if self.IsInterface() {
		flgs = append(flgs, "interface")
	}
	if self.IsAbstract() {
		flgs = append(flgs, "abstract")
	}
	if self.IsSynthetic() {
		flgs = append(flgs, "synthetic")
	}
	if self.IsAnnotation() {
		flgs = append(flgs, "annotation")
	}
	if self.IsEnum() {
		flgs = append(flgs, "enum")
	}
	if self.IsModule() {
		flgs = append(flgs, "module")
	}
	return fmt.Sprintf("[%s]", strings.Join(flgs, " "))
}

func (self AccessFlags) IsPublic() bool {
	return self&ACC_PUBLIC == ACC_PUBLIC
}
func (self AccessFlags) IsFinal() bool {
	return self&ACC_FINAL == ACC_FINAL
}
func (self AccessFlags) IsSuper() bool {
	return self&ACC_SUPER == ACC_SUPER
}
func (self AccessFlags) IsInterface() bool {
	return self&ACC_INTERFACE == ACC_INTERFACE
}
func (self AccessFlags) IsAbstract() bool {
	return self&ACC_ABSTRACT == ACC_ABSTRACT
}
func (self AccessFlags) IsSynthetic() bool {
	return self&ACC_SYNTHETIC == ACC_SYNTHETIC
}
func (self AccessFlags) IsAnnotation() bool {
	return self&ACC_ANNOTATION == ACC_ANNOTATION
}
func (self AccessFlags) IsEnum() bool {
	return self&ACC_ENUM == ACC_ENUM
}
func (self AccessFlags) IsModule() bool {
	return self&ACC_MODULE == ACC_MODULE
}

const (
	ACC_PUBLIC     = 0x0001
	ACC_FINAL      = 0x0010
	ACC_SUPER      = 0x0020
	ACC_INTERFACE  = 0x0200
	ACC_ABSTRACT   = 0x0400
	ACC_SYNTHETIC  = 0x1000
	ACC_ANNOTATION = 0x2000
	ACC_ENUM       = 0x4000
	ACC_MODULE     = 0x8000
)
