package main

import (
	"image"
)


func Img2rgb(img image.Image) []uint8 {

	bounds := img.Bounds()

	width, height := bounds.Dx(), bounds.Dy()

	size := width * height

	// 初始化每个通道的二维切片
	R := make([]uint8, size)
	G := make([]uint8, size)
	B := make([]uint8, size)

	i := 0
	for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
		for x := bounds.Min.X; x < bounds.Max.X; x++ {
			r, g, b, _ := img.At(x, y).RGBA()
			R[i] = uint8(r >> 8)
			G[i] = uint8(g >> 8)
			B[i] = uint8(b >> 8)
			i++
		}
	}

	R = append(R, append(G, B...)...)

	return R
}
