package main

import "fmt"

func main() {
	var a int = 8
	var b int = 18

	string1 := "Hello" //:= 左侧如果没有声明新的变量，就产生编译错误，只能被用在函数体内，而不可以用于全局变量的声明与赋值
	var string2 string = "World"
	fmt.Println(string1 + "  " + string2)

	fmt.Println("a:", a)
	fmt.Println("b:", b)
	fmt.Println(a + b)

	//交换两个变量的值
	a, b = b, a
	fmt.Println("a:", a)
	fmt.Println("b:", b)

	//空白标识符_:  _ 实际上是一个只写变量，你不能得到它的值。
	//这样做是因为 Go 语言中你必须使用所有被声明的变量
	//但有时你并不需要使用从一个函数得到的所有返回值。
	u, n, _ := members()
	fmt.Println(u, n)
}

func members() (int, int, string) {
	x, y, z := 6, 6, "mmm"
	return x, y, z
}
