package io

import "testing"

func TestNewDataInputStream(t *testing.T) {
	dis := NewDataInputStream("./test/datainputstream")
	d, err := dis.ReadByte()
	if err != nil {
		t.Fatal(err)
	}
	d, err = dis.ReadByte()
	if err != nil {
		t.Fatal(err)
	}
	t.Fatalf("%c", d)
}
