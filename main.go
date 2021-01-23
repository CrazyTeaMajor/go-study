package main

// import "fmt"
import (
	"fmt"
	"math"
)

//两种import库的方式

func output() {

	fmt.Println("Hello, World!")
	fmt.Println("Google" + "Baidu")
	//字符串可直接相加

	var character1 = 'a'
	fmt.Println(character1) //这样子输出的是ASCII码值
	var character2 = "a"
	fmt.Println(character2)
	var str string = "a"
	fmt.Println(str)

	var number1 float64 = 9.0
	fmt.Println(math.Sqrt(number1))
	//数学库

	var n1, n2, n3 = "String", 50, 4.13
	fmt.Printf("%s;%d;%f\n", n1, n2, n3)
	fmt.Printf("float = %.3f\n", n3)
	fmt.Printf("Long float = %lf\n", n3) //这种不对，其余和C差不多
	fmt.Printf("% \n", n3)               //错误
	fmt.Printf("%v\n", n3)               //不知道类型可以直接使用%v

	fmt.Print("\n")

	var n4 complex64 = complex(5, 2)
	fmt.Println(n4)
}

/*
int8~64 int -> long long
uint8~64 unsigned int -> unsigned long long
float32~64 float -> double
complex64~128 -> 一个complex结构体
*/

func maxInt(x, y int) int {
	if x > y {
		return x
	} else {
		return y
	}
}

//int 默认和 uint 一样位数，只不过有符号

func minInt(x, y int) int {
	// return x > y ? y : x 没有三目运算符
	if x > y {
		return y
	} else {
		return x
	}
}

func testVarType() {
	var number1, number2 int
	fmt.Print("Please input two numbers, in a line, and then enter\n")
	fmt.Scanf("%d%d\n", &number1, &number2) //输入后会自动使输入消失
	fmt.Println(maxInt(number1, number2))
	fmt.Printf("Small one is %d\n", minInt(number1, number2))

	var flag bool
	fmt.Println(flag)
	//fmt.Printf("%d\n", flag) 错误的，Go的bool不能像C++一样和int型互通
	flag = true
	fmt.Println(flag)
	//flag = 0 不能这样写

	var testStr string
	fmt.Print("String test\n")
	fmt.Println(testStr)
	fmt.Printf("%s\n%q\n", testStr, testStr)
	//%q 是让字符串带双引号输出
	fmt.Scanf("%s", &testStr) // 字符串也要带&符号才能改变值
	fmt.Printf("%q\n", testStr)

	fmt.Println()
}

/*
默认变量赋值
数值为0，布尔为false，字符串为空字符串""，其余为空nil
其余为
var a *int
var a []int
var a map[string] int
var a chan int
var a func(string) int
var a error // error 是接口
*/

func main() {
	//	output()
	testVarType()
}
