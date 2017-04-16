package main

import (
	"golang.org/x/tour/wc"
	"strings"
)

/*
	自力作成版
	引数で渡された文字列に含まれる各単語の出現数を計算
*/
func WordCount(s string) map[string]int {
	wordCounter := make(map[string]int)
	fields := strings.Fields(s) // 文字列sを単語(string)の配列に分解
	// 単語の数だけ繰り返す
	for _, v := range fields {
		_, ok := wordCounter[v] // 既出単語か確認
		if ok {
			wordCounter[v]++ // 既出単語なので出現数増加
		} else {
			wordCounter[v] = 1 // 初出単語なので出現数1で配列に追加
		}
	}
	return wordCounter
}

/*
	WordCountの改善版
*/
func WordCount_Web(s string) map[string]int {
	wordCounter := make(map[string]int)
	for _, v := range strings.Fields(s) {
		wordCounter[v] += 1
	}
	return wordCounter
}

func main() {
	wc.Test(WordCount)
	//wc.Test(WordCount_Web)
}
