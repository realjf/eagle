package io

import "testing"

func TestNewDataOutputStream(t *testing.T) {
	var dos *DataOutputStream
	dos = NewDataOutputStream("./test/dataoutputstream")
	if dos == nil {
		t.Fatal("file open error")
	}
	err := dos.WriteByte('c')
	if err != nil {
		t.Fatal(err)
	}
	n, err := dos.WriteInt(1)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(n)
	n, err = dos.WriteString("world")
	if err != nil {
		t.Fatal(err)
	}
	t.Fatal(n)
}
