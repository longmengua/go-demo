package common

import "testing"

func TestMathPlus(t *testing.T) {
	// 加法
	result1, _ := Math("add", "0.987654321", "0.123456789")
	assert1 := "1.11111111"
	if assert1 != result1 {
		t.Errorf("Math func failed, %s, %s", result1, assert1)
	}
	// 減法
	result2, _ := Math("sub", "0.987654321", "0.123456789")
	assert2 := "0.864197532"
	if assert2 != result2 {
		t.Errorf("Math func failed, %s, %s", result2, assert2)
	}
	// 乘法
	result3, _ := Math("mul", "0.987654321", "0.123456789")
	assert3 := "0.121932631112635269"
	if assert3 != result3 {
		t.Errorf("Math func failed, %s, %s", result3, assert3)
	}
	// 除法
	result4, _ := Math("div", "10", "3")
	assert4 := "3.3333333333333333"
	if assert4 != result4 {
		t.Errorf("Math func failed, %s, %s", result4, assert4)
	}
}
