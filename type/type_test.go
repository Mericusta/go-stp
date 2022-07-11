package stptype

import (
	"fmt"
	"testing"
)

type customType struct {
	i Int
	s String
}

func (v customType) EQ(rv customType) bool {
	return v.i == rv.i && v.s == rv.s
}

func (v customType) LT(rv customType) bool {
	return v.i < rv.i && v.s < rv.s
}

func (v customType) GT(rv customType) bool {
	return v.i > rv.i && v.s > rv.s
}

func Test_Type(t *testing.T) {
	int1, int2 := Int(1), Int(2)
	if int1.LT(int2) {
		fmt.Printf("Int %v LT %v\n", int1, int2)
	} else {
		panic(fmt.Sprintf("Int %v not LT %v\n", int1, int2))
	}
	if int1.GT(int2) {
		panic(fmt.Sprintf("Int %v GT %v\n", int1, int2))
	} else {
		fmt.Printf("Int %v not GT %v\n", int1, int2)
	}
	if int1.EQ(int2) {
		panic(fmt.Sprintf("Int %v EQ %v\n", int1, int2))
	} else {
		fmt.Printf("Int %v not EQ %v\n", int1, int2)
	}
	if int1.V() == 1 {
		fmt.Printf("Int %v == %v\n", int1.V(), 1)
	} else {
		panic(fmt.Sprintf("Int %v != %v\n", int1.V(), 1))
	}
	if int2.V() == 2 {
		fmt.Printf("Int %v == %v\n", int2.V(), 2)
	} else {
		panic(fmt.Sprintf("Int %v != %v\n", int2.V(), 2))
	}

	ct1, ct2 := customType{i: 1, s: "1"}, customType{i: 2, s: "2"}
	if ct1.LT(ct2) {
		fmt.Printf("Int %v LT %v\n", ct1, ct2)
	} else {
		panic(fmt.Sprintf("Int %v not LT %v\n", ct1, ct2))
	}
	if ct1.GT(ct2) {
		panic(fmt.Sprintf("Int %v GT %v\n", ct1, ct2))
	} else {
		fmt.Printf("Int %v not GT %v\n", ct1, ct2)
	}
	if ct1.EQ(ct2) {
		panic(fmt.Sprintf("Int %v EQ %v\n", ct1, ct2))
	} else {
		fmt.Printf("Int %v not EQ %v\n", ct1, ct2)
	}
}
