package main

import "fmt"

//包内部变量
var (
	 aa = 3
	 ss = "kkk"
	 bb = true
)

func variableZeroValue()  {
	var i int
	var s string
	fmt.Println(i,s)
	fmt.Printf("%d %q\n",i,s)
}

func variableInitValue()  {
	var i,j int = 3,4
	var s string = "abc"
	fmt.Printf("%d %d %q",i,j,s)
	fmt.Println()
}

func variableTypeDeduction()  {
	var i,j,d,f = 3,4,true,"def"
	fmt.Println(i,j,d,f)
}

func variableShorter()  {
	i,j,d,f := 3,4,true,"def"
	fmt.Println(i,j,d,f)
}

func main() {
	fmt.Println("hello")
	variableZeroValue()
	variableInitValue()
	variableTypeDeduction()
	variableShorter()
	fmt.Println(aa,bb,ss)
}
