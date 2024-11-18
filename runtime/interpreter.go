package runtime

import (
	"bytes"
	"encoding/binary"
	"fmt"

	"gjvm/classfile"
)

func Interpret(code []byte, stack *OperandStack, cp classfile.ConstantPool) {
	r := newReader(code)
	sys := NewSystem()

	frames := make([]interface{}, 0)

	for r.hasNext() {
		frame, err := r.readFrame(cp)
		if err != nil {
			fmt.Println(err)
			break
		}
		fmt.Printf("- %v\n", frame)
		frames = append(frames, frame)
	}

	fmt.Println("=================================================================")
	fmt.Println("Exec result")
	fmt.Println("=================================================================")
	for _, frame := range frames {
		switch frame.(type) {
		case *GetStatic:
			stack.Push(frame)
		case *Ldc:
			stack.Push(frame)
		case *InvokeVirtual:
			frame := frame.(*InvokeVirtual)
			args := make([]interface{}, len(frame.ArgTypes))
			for i := range frame.ArgTypes {
				frame := stack.Pop()
				switch frame.(type) {
				case *Ldc:
					args[i] = frame.(*Ldc).value
				}
			}

			callable := stack.Pop().(*GetStatic)
			method := fmt.Sprintf("%s.%s.%s", callable.Class, callable.Name, frame.Name)
			sys.Call(method, args...)
		}
	}
}

type reader struct {
	r *bytes.Reader
}

func newReader(code []byte) *reader {
	return &reader{bytes.NewReader(code)}
}

func (r *reader) hasNext() bool {
	return r.r.Len() > 0
}

func (r *reader) readFrame(cp classfile.ConstantPool) (Frame, error) {
	opcode, err := r.r.ReadByte()
	if err != nil {
		return nil, err
	}

	switch opcode {
	case 0xb2: // getstatic
		index := make([]byte, 2)
		r.r.Read(index)
		i := binary.BigEndian.Uint16(index)
		constant := cp[i].(*classfile.ConstantFieldrefInfo)
		filedRef := constant.Resolve(cp)
		return &GetStatic{filedRef.Class, filedRef.Name, filedRef.Descriptor}, nil
	case 0x12: // ldc
		index := make([]byte, 1)
		r.r.Read(index)
		i := uint16(index[0])
		constantInfo := cp[i]
		switch constantInfo.(type) {
		case *classfile.ConstantStringInfo:
			str := constantInfo.(*classfile.ConstantStringInfo).Resolve(cp)
			return &Ldc{str}, nil
		default:
			return nil, fmt.Errorf("unsupported: %T", constantInfo)
		}
	case 0xb6: // invokevirtual
		index := make([]byte, 2)
		r.r.Read(index)
		i := binary.BigEndian.Uint16(index)
		method := cp[i].(*classfile.ConstantMethodrefInfo).Resolve(cp)
		return &InvokeVirtual{method.Class, method.Name, method.ArgTypes}, nil
	case 0xb1: // return
		return &Return{}, nil
	}

	return nil, fmt.Errorf("unknown opcode: %d", opcode)
}

type Frame interface {
	String() string
}

type GetStatic struct {
	Class      string
	Name       string
	Descriptor string
}

func (gs *GetStatic) String() string {
	return fmt.Sprintf("0xb2 getstatic: %s.%s %s", gs.Class, gs.Name, gs.Descriptor)
}

type Ldc struct {
	value any
}

func (ldc *Ldc) String() string {
	return fmt.Sprintf("0x12 ldc: %v", ldc.value)
}

type InvokeVirtual struct {
	Class    string
	Name     string
	ArgTypes []string
}

func (iv *InvokeVirtual) String() string {
	return fmt.Sprintf("0xb6 invokevirtual: %s.%s %v", iv.Class, iv.Name, iv.ArgTypes)
}

type Return struct{}

func (r *Return) String() string {
	return fmt.Sprintf("0xb1 return")
}

type Unsupported struct {
	opcode byte
}

func (u *Unsupported) String() string {
	return fmt.Sprintf("unsupported opcode: %d", u.opcode)
}

type OperandStack []interface{}

func (s *OperandStack) Push(value interface{}) {
	*s = append(*s, value)
}
func (s *OperandStack) Pop() interface{} {
	v := (*s)[len(*s)-1]
	*s = (*s)[:len(*s)-1]
	return v
}
func NewOperandStack() *OperandStack {
	return &OperandStack{}
}
