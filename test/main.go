package main

import (
	"fmt"
	"github.com/GeneralRelativityTheory/clientProject/util"
)

func main() {
	var str string = "3434343434343434343434FFFFFFFFFFFFFFFFFFFF"
	//util.GetStringLength(str)
	//util.String2Ascii(str)
	//util.TrimRight(str)
	result := util.StringAscii2String(str)
	fmt.Println(result)
}
