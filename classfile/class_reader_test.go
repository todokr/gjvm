package classfile

import (
	"bytes"
	"testing"
)

var data = []byte{0x01, 0x02, 0x03, 0x04}

func TestReadU1(t *testing.T) {
	r := bytes.NewReader(data)
	reader := newClassReader(r)
	if reader.ReadU1() != 1 {
		t.Errorf("readU1() failed")
	}
}

func TestReadU2(t *testing.T) {
	r := bytes.NewReader(data)
	reader := newClassReader(r)
	if reader.ReadU2() != 0x0102 {
		t.Errorf("readU2() failed")
	}
}

func TestReadU4(t *testing.T) {
	r := bytes.NewReader(data)
	reader := newClassReader(r)
	if reader.ReadU4() != 0x01020304 {
		t.Errorf("readU4() failed")
	}
}

func TestRead4(t *testing.T) {
	r := bytes.NewReader(data)
	reader := newClassReader(r)
	if reader.Read4() != 0x01020304 {
		t.Errorf("read4() failed")
	}
}
