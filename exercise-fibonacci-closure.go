package main

import "fmt"

func fibonacci() func() int {
	f_0 := 0
	f_1 := 1
	count := 0
	return func() int {
		var f int
		if count == 0 {
			f = f_0
		} else if count == 1 {
			f = f_1
		} else {
			f = f_0 + f_1
			f_0 = f_1
			f_1 = f
		}
		count++
		return f
	}
}

/*
改善版
countのインクリメントを始めの2回のみに変更
フィボナッチ数の計算用変数fを廃止
*/
func fibonacci_Web() func() int {
	f_0 := 0
	f_1 := 1
	count := 0
	return func() int {
		switch count {
		case 0:
			count++
			return f_0
		case 1:
			count++
			return f_1
		default:
			f_0, f_1 = f_1, f_1+f_0
			return f_1
		}
	}
}

func main() {
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}
