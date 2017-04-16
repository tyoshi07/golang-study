package main

import "golang.org/x/tour/pic"

/*
Pic関数の実装演習
*/

// 自力実装Pic
func Pic(dx, dy int) [][]uint8 {
	pic := make([][]uint8, dy) // 生成する画像のスライス
	for y := range pic {
		pic_x := make([]uint8, dx) // 生成する画像の行スライス
		for x := range pic_x {
			pic_x[x] = uint8(x * y)
		}
		pic[y] = pic_x
	}
	return pic
}

// 改善版Pic
func Pic_Web(dx, dy int) [][]uint8 {
	pic := make([][]uint8, dy)
	for y := range pic {
		pic[y] = make([]uint8, dx)
		for x := range pic[y] {
			pic[y][x] = uint8(x * y)
		}
	}
	return pic
}

/*
func Pic_Array(dx, dy int) [][]uint8 {
	var pic [256][256]uint8
	for y := range pic {
		for x := 0; x < dx; x++ {
			pic[y][x] = uint8(y*x)
		}
	}
	pic_slice := make([][]uint8, dy)
	for y := range pic_slice {
		pic_slice[y] = pic[y][:]
	}
	return pic_slice[:][:]
}
*/

func main() {
	pic.Show(Pic)
	//pic.Show(Pic_Web)
}
