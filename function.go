package main

import (
	"fmt"
)

func FirstOrSecond() {
	if Count%2 == 0 {
		FS = true //先手
	} else {
		FS = false //後手
	}
}

func SenteGote() {
	if FS {
		SG = "先手"
	} else {
		SG = "後手"
	}
}

//--------------------------------------------------------------------------------
// ▼ 将棋盤の表示

func Display() {
	XNum := [9]int{9, 8, 7, 6, 5, 4, 3, 2, 1}
	fmt.Printf("  ")
	for _, i := range XNum {
		fmt.Printf("%d%s", i, w3)
	}
	fmt.Println("")
	fmt.Println(line)
	RepeatYNum()
}

func RepeatYNum() {
	Position(Y1Num, "一")
	Position(Y2Num, "二")
	Position(Y3Num, "三")
	Position(Y4Num, "四")
	Position(Y5Num, "五")
	Position(Y6Num, "六")
	Position(Y7Num, "七")
	Position(Y8Num, "八")
	Position(Y9Num, "九")
}
func Position(yline [9]int, y string) {
	fmt.Printf("|")
	for _, x := range yline {
		fmt.Printf("%s%s", Data[x], "|")
	}
	fmt.Println(" ", y)
	fmt.Println(line)
	return
}

//      将棋盤の表示（初期値：スペース３つ分）
//    9   8   7   6   5   4   3   2   1
//   +---+---+---+---+---+---+---+---+---+
//   |91 |81 |71 |61 |51 |41 |31 |21 |11 | 一
//   +---+---+---+---+---+---+---+---+---+
//   |92 |82 |72 |62 |52 |42 |32 |22 |12 | 二
//   +---+---+---+---+---+---+---+---+---+
//   |93 |83 |73 |63 |53 |43 |33 |23 |13 | 三
//   +---+---+---+---+---+---+---+---+---+
//   |94 |84 |74 |64 |54 |44 |34 |24 |14 | 四
//   +---+---+---+---+---+---+---+---+---+
//   |95 |85 |75 |65 |55 |45 |35 |25 |15 | 五
//   +---+---+---+---+---+---+---+---+---+
//   |96 |86 |76 |66 |56 |46 |36 |26 |16 | 六
//   +---+---+---+---+---+---+---+---+---+
//   |97 |87 |77 |67 |57 |47 |37 |27 |17 | 七
//   +---+---+---+---+---+---+---+---+---+
//   |98 |88 |78 |68 |58 |48 |38 |28 |18 | 八
//   +---+---+---+---+---+---+---+---+---+
//   |99 |89 |79 |69 |59 |49 |39 |29 |19 | 九
//   +---+---+---+---+---+---+---+---+---+
//--------------------------------------------------------------------------------

func InputAgain() {
	fmt.Scan(&Input)
}

func Hirate() {
	for _, i := range Y7Num {
		Data[i] = "歩 "
	}
	for _, i := range Y3Num {
		Data[i] = "#歩"
	}
	Data[82], Data[28] = "#飛", "飛 "
	Data[22], Data[88] = "#角", "角 "
	Data[51], Data[59] = "#王", "玉 "
	Data[11], Data[91], Data[19], Data[99] = "#香", "#香", "香 ", "香 "
	Data[21], Data[81], Data[29], Data[89] = "#桂", "#桂", "桂 ", "桂 "
	Data[31], Data[71], Data[39], Data[79] = "#銀", "#銀", "銀 ", "銀 "
	Data[41], Data[61], Data[49], Data[69] = "#金", "#金", "金 ", "金 "
}

//ユーザーの入力
func UserInput(words string) int {
	fmt.Printf("%s%s%s", "+", line36, "+")
	fmt.Println()
	fmt.Println("|", words, "|")
	fmt.Printf("%s%s%s", "+", line36, "+")
	fmt.Println()
	fmt.Scan(&Input)
	return Input
}
