package main

import (
	"strings"
)

var (
	SG  string
	MS0 = "     < %s >    [2]:終了する"
	MS1 = "[1]:ゲームをはじめる  [2]:終了する"
	MS2 = "      コマを選択してください      "
	MS3 = "  コマの符号を入力してください    "
	//ms4  = "数値を入力してください\n"
	//ms5  = "数値を入力してください\n"
	//ms6  = "数値を入力してください\n"
)

var (
	error1 = "11~99の範囲の数値を入力してください"
	error2 = "0を含む数値は入力できません"
)

var (
	line   = "+---+---+---+---+---+---+---+---+---+"
	line36 = strings.Repeat("-", 36)
	w3     = "   "
)

var Count = 0
var FS bool

var (
	Input      int
	FirstInput int
)

var Data [111]string

func Initialize() {
	for i := range Data {
		Data[i] = w3
		i++
	}
	return
}

var (
	Y1Num = [9]int{91, 81, 71, 61, 51, 41, 31, 21, 11}
	Y2Num = [9]int{92, 82, 72, 62, 52, 42, 32, 22, 12}
	Y3Num = [9]int{93, 83, 73, 63, 53, 43, 33, 23, 13}
	Y4Num = [9]int{94, 84, 74, 64, 54, 44, 34, 24, 14}
	Y5Num = [9]int{95, 85, 75, 65, 55, 45, 35, 25, 15}
	Y6Num = [9]int{96, 86, 76, 66, 56, 46, 36, 26, 16}
	Y7Num = [9]int{97, 87, 77, 67, 57, 47, 37, 27, 17}
	Y8Num = [9]int{98, 88, 78, 68, 58, 48, 38, 28, 18}
	Y9Num = [9]int{99, 89, 79, 69, 59, 49, 39, 29, 19}
)
