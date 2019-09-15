## Google slide
https://docs.google.com/presentation/d/1H0AOn_BquDLDvRer-ObzCoY0bSWrR18zI-bVTVTXUH8/edit?usp=sharing

## 宣告變數

Golang變數基本有以下型態：

```go
	// 正整數
	var postivInt uint = 100
	fmt.Println(postivInt)

	// 帶符號整數
	var interger int = -100
	fmt.Println(interger)

	// 小數
	var dec float32 = 3.14
	fmt.Println(dec)

	// 布林
	var tf bool = true
	fmt.Println(tf)

	// 字串，不可用''
	var str string = "hello world"
	var newYear string = "新年快樂"
	// Go中的字串類似陣列
	fmt.Println(string(str[0]))
	fmt.Println(string(newYear[0]))
	// 一個中文字(Unicode)佔3個長度
	fmt.Println(len(newYear))

	// 字元
	var char rune = 'a'
	fmt.Println(char)
	fmt.Println(string(char))

	// 處理中文字元
	var new rune = []rune(newYear)[1]
	fmt.Print(new, string(new))
```

接著在Go中有不同宣告的方式，效果都是一樣的：

```go
    var a int = 1
    var b = 1
    var c int; c = 1
    var d = int(1)
    e := 1
```