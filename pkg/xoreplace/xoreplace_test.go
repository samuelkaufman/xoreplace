package xoreplace

import (
	"bytes"
	"testing"
)

func TestXoReplace(t *testing.T) {
	var b bytes.Buffer
	xo := New("X0", &b)
	xo.Run()
	if b.String() != "00\n10\n" {
		t.Errorf("Got incorrect sequence for X0 [%s] [%s]", b.String(), "00\n10\n")
	}
	var b2 bytes.Buffer
	xox := New("X0X", &b2)
	xox.Run()
	if b2.String() != "000\n001\n100\n101\n" {
		t.Errorf("Got incorrect sequence for X0X: %s", b2.String())
	}
	var b3 bytes.Buffer
	xo = New("10X10X0", &b3)
	xo.Run()
	if b3.String() != "1001000\n1001010\n1011000\n1011010\n" {
		t.Errorf("Got incorrect sequence for 10X10X0\n[%s]\n", b3.String())
	}
}
