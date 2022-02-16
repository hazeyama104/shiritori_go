package main

import (
	"fmt"
	"os"
    "bufio"
)

func main() {
	fmt.Printf("しりとりスタート!\n")
	firstWord := "り"
	start(firstWord)
}

//ゲームスタートの関数
func start(word string){
	fmt.Printf("最初の文字：　%s\n", word)
	text := input()
	judge(text, word)
}

//入力するための関数
func input() string{
	fmt.Printf("文字を入力してください\n")
	stdin := bufio.NewScanner(os.Stdin)
	stdin.Scan()
	text := stdin.Text()
	return text
}

//判定の関数
func judge(text, word string){
	textRune := []rune(text)
	firstWord := string(textRune[0])
	size := len(textRune)

	if firstWord != word {
		fmt.Printf("最初の文字が違います。\n")
		fmt.Printf("入力した最初の文字：%s\n", firstWord)
		fmt.Printf("もう一度入力してください。\n")
		start(word)
	}

	//入力された最後の文字を取得
	lastWord := string(textRune[size-1])
	fmt.Printf("最後の文字： %s", lastWord)

	//最後が「ん」だったらゲーム終了
	if lastWord == "ん" {
		fmt.Printf("最後の文字が「ん」です。\n")
		end()
	}

	start(lastWord)
}
func end() {
	fmt.Printf("GAME OVER\n")
	os.Exit(0)
}
