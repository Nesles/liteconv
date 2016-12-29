package liteconv

import "testing"

func Test(t *testing.T)  {

	if  InterfaceToString(32525) == "32525" {
		t.Log("INT32-PASS")
	}else {
		t.Fail()
	}

	if  InterfaceToString("32525") == "32525" {
		t.Log("string-PASS")
	}else {
		t.Fail()
	}

	if  InterfaceToString(int64(32525)) == "32525" {
		t.Log("INT64-PASS")
	}else {
		t.Fail()
	}
	if  InterfaceToString(uint32(32525)) == "32525" {
		t.Log("UINT32-PASS")
	}else {
		t.Fail()
	}
	if  InterfaceToString(uint64(32525)) == "32525" {
		t.Log("UINT64-PASS")
	}else {
		t.Fail()
	}
	/*
	if  InterfaceToString(32525) == "32525" {
		t.Log("INT32-PASS")
	}else {
		t.Fail()
	}
	*/

}