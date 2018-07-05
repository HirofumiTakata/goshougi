package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	if Count == 0 {
		Initialize()
		Hirate()
	}
	Stage1()
}

func Stage1() {
	UserInput(MS1) //"[1]:ゲームをはじめる  [2]:終了する"
	switch Input {
	case 1, 0:
		Stage2()
	case 2:
		os.Exit(0)
	}
}

func Stage2() {
	FirstOrSecond()
	Display()
	SenteGote()
	fmt.Printf(MS0, SG) //"＜" + 先手or後手 + "＞  [2]:終了する"
	fmt.Println()
	UserInput(MS2) //"コマを選択してください"
	FirstInput = Input
	Stage3()
}

func Stage3() {
	if Input < 11 || 99 < Input {
		fmt.Println(error1)
		InputAgain()
	} else if strings.Contains(string(Input), "0") {
		fmt.Println(error2)
		InputAgain()
	} else {
		Stage4()
	}
}

func Stage4() {
	//validation
	UserInput(MS3)
	Data[Input] = Data[FirstInput]
	Data[FirstInput] = "   "
	Count++
	Stage2()
}
