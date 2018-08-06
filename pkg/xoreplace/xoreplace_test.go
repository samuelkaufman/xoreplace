package xoreplace

import "testing"

func TestXoReplace(t *testing.T) {
	xo := New("X0")
	xo.Run()
	xo = New("X0X")
	xo.Run()
	xo = New("XX")
	xo.Run()
}
