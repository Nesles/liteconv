package liteconv

import "testing"

func Test_ToString(t *testing.T) {

	if Float64ToString(5)=="5"{
		t.Log("Float64ToString OK")
	}else {
		t.Fail()
	}

	if BoolToString(true) == "true" && BoolToString(false) == "false"{
		t.Fail()
	}else {
		t.Log("BoolToString OK")
	}

	if Uint64ToString(256)=="256"{
		t.Log("Uint64ToString OK")
	}

}
