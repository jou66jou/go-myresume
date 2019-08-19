package main

import (
	"flag"
	"log"
	"os"
)

var (
	// 宣告一個可以被放置git commit number的變數
	gitcommitnum string
	// 增加一個判斷是否顯示git commit number的旗幟
	checkcommit = flag.Bool("version", false, "burry code for check version")
)

func main() {
	flag.Parse()

	// 如果旗幟顯示要檢查git commit num就顯示後並跳出程式
	if *checkcommit {
		checkComimit()
		os.Exit(1)
	}
}

// 顯示git commit num的函數
func checkComimit() {
	log.Println(gitcommitnum)
}
