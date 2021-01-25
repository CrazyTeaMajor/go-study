package main

import "fmt"

var codeList [26][26]int
var codeStr string
var codeLength int

func transform(letter byte) int {
	if 'a' <= letter && letter <= 'z' {
		return (int)(letter - 'a')
	} else if 'A' <= letter && letter <= 'Z' {
		return (int)(letter - 'A')
	} else {
		return -1
	}
}

func letterType(letter byte) int {
	if 'a' <= letter && letter <= 'z' {
		return 'a'
	} else if 'A' <= letter && letter <= 'Z' {
		return 'A'
	} else {
		return -1
	}
}

func initCodeList() {
	for i := 0; i < 26; i++ {
		for j := 0; j < 26; j++ {
			codeList[i][j] = (i + j) % 26
		}
	}
}

func encode(input string) string {
	var output string
	var tx, ty int
	var inputLength = len(input)
	for i := 0; i < inputLength; i++ {
		tx = transform(input[i])
		ty = transform(codeStr[i%codeLength])
		output += string(letterType(input[i]) + codeList[tx][ty])
	}
	return output
}

func decode(input string) string {
	var output string
	var tx, ty int
	var inputLength = len(input)
	for i := 0; i < inputLength; i++ {
		tx = transform(input[i])
		ty = transform(codeStr[i%codeLength])
		for j := 0; j < 26; j++ {
			if codeList[j][ty] == tx {
				output += string(letterType(input[i]) + j)
				break
			}
		}
	}
	return output
}

func work() {
	initCodeList()
	fmt.Println("请输入加密密钥:(仅包含大小写英文字母，不包含空格等其它任何字符)\n")
	fmt.Scanf("%s\n", &codeStr)
	codeLength = len(codeStr)
	fmt.Println("请输入要进行的操作：1加密，2解密\n")
	var optType int
	fmt.Scanf("%d\n", &optType)
	if optType == 1 {
		fmt.Println("请输入要加密的密文:(仅包含大小写英文字母，不包含空格等其它任何字符)\n")
		var inputStr string
		fmt.Scanf("%s\n", &inputStr)
		var outputStr string
		outputStr = encode(inputStr)
		fmt.Println(outputStr)
	} else if optType == 2 {
		fmt.Println("请输入要解密的密:(仅包含大小写英文字母，不包含空格等其它任何字符)\n")
		var inputStr string
		fmt.Scanf("%s\n", &inputStr)
		var outputStr string
		outputStr = decode(inputStr)
		fmt.Println(outputStr)
	}

}
