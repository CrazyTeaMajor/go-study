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

var arraySize int = 10001

const totSize int = 100001
const adminName string = "Admin"
const userName = "czt"

const (
	//用来枚举
	Xiuji  = 0
	female = 1
	male   = 2
)

func testStudyTwo() {
	var number0 = 5.2
	fmt.Printf("%f\n", number0)
	newNumber := 4.8
	fmt.Println(newNumber + number0)
	var number1 int
	fmt.Println(number1)
	number1, number2 := 1, 2
	fmt.Println(number1)
	fmt.Println(number2)
	var (
		str1    string
		number3 int
		year    int
		day     int
		height  float32
	)
	//一般用作声明全局不同变量
	str1 = "Awasdsx"
	fmt.Println(str1 + " is XXXX")
	number3 = arraySize
	fmt.Printf("Const %d = %d\n", number3, arraySize)
	year = 2021
	day = 23
	height = 1.8
	fmt.Printf("%d.1.%d my height is %.2f\n", year, day, height)
}

func optTest() {
	var n1, n2, n3 = 2, 3, 4
	fmt.Printf("n1 = %d, n2 = %d, n3 = %d\n", n1, n2, n3)
	fmt.Print("n1 + n2 = ")
	fmt.Println(n1 + n2)
	fmt.Print("n1 - n2 = ")
	fmt.Println(n1 - n2)
	fmt.Print("n1 * n2 = ")
	fmt.Println(n1 * n2)
	fmt.Print("n1 / n2 = ")
	fmt.Println(n1 / n2)
	fmt.Print("n3 / n1 = ")
	fmt.Println(n3 / n1)
	fmt.Print("n2 % n1 = ")
	fmt.Println(n2 % n1)
	fmt.Print("n1++ =")
	n1++
	fmt.Println(n1)
	fmt.Print("n1-- =")
	n1--
	fmt.Println(n1)
	fmt.Print("n1 += n2 =")
	n1 += n2
	fmt.Println(n1)
	fmt.Print("n1 *= n2 =")
	n1 *= n2
	fmt.Println(n1)
	fmt.Print("n1 -= n2 =")
	n1 -= n2
	fmt.Println(n1)
	//fmt.Println(n1--)
	//	++n1
	//但是这两种不行，自增自减不能写在前面，也不能写在输出里面，保证了逻辑的简单性
	//Go的运算符基本上与C++没有差别，包括逻辑运算符等
	if n1 >= 0 && n2 >= 0 {
		fmt.Println("Yes")
	}
	// if(n1 >= 0 && n2 >= 0) {} 这种加小括号也行，但是没有必要
	if ((n1 + 5) * n2 - 100) >= 0 {
		fmt.Println("Oh, yeah")
	} else {
		fmt.Println("Oh, no")
	}
	var c1 = 1.0 + 2.0i
	fmt.Println(c1)
	//可直接使用复数符号
}

func ptrTest() {
	var a int = 4
	var ptr *int
	fmt.Println(a)
	ptr = &a
	*ptr++
	fmt.Println(a)
}
/*
import "base"
base.print("ASD")
要用自己写的包时，先在cmd里面go install 你的包的名称
*/
func main() {
	//fmt.Print(printtest("ASD"))
	// output()
	// testVarType()
	// testStudyTwo()
	// test()
	// optTest()
	ptrTest()
}