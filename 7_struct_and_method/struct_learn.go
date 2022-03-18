package main

import (
	"fmt"
	"github.com/Dlimingliang/go-learn/7_struct_and_method/structPack"
	"reflect"
)

type struct1 struct {
	i1  int
	f1  float32
	str string
}

type TagType struct { // tags
	field1 bool   "An important answer"
	field2 string "The name of the thing"
	field3 int    "How much there are"
}

type TwoInts struct {
	a int
	b int
}

func main() {
	s1 := new(struct1)
	s1.i1 = 1
	s1.f1 = 3.14
	s1.str = "aaa"
	fmt.Println(s1)

	struct1 := new(structPack.ExpStruct)
	struct1.Mi1 = 10
	struct1.Mf1 = 16.

	fmt.Printf("Mi1 = %d\n", struct1.Mi1)
	fmt.Printf("Mf1 = %f\n", struct1.Mf1)

	tt := TagType{true, "Barak Obama", 1}
	for i := 0; i < 3; i++ {
		regTag(tt, i)
	}

	th := new(TwoInts)
	th.a = 1
	th.b = 2
	fmt.Println(th.addThem(7))

}

func regTag(tt TagType, ix int) {
	ttType := reflect.TypeOf(tt)
	ixField := ttType.Field(ix)
	fmt.Printf("%v\n", ixField.Tag)
	fmt.Printf("%v\n", ixField.Name)
}

func (th *TwoInts) addThem(param int) int {
	return th.a + th.b + param
}
