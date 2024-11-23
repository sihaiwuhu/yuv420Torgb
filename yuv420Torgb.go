package main

import (
	"fmt"
	"image"
	"image/color"
	"sync"
)

func yuv420ToRGB(yuvData []byte, width, height int) (*image.RGBA, error) {
	img := image.NewRGBA(image.Rect(0, 0, width, height))

	ySize := width * height
	uSize := ySize / 4
	vSize := ySize / 4

	if len(yuvData) < ySize+uSize+vSize {
		return nil, fmt.Errorf("YUV 数据长度不足")
	}

	yData := yuvData[:ySize]
	uData := yuvData[ySize : ySize+uSize]
	vData := yuvData[ySize+uSize : ySize+uSize+vSize]

	var wg sync.WaitGroup

	for row := 0; row < height; row++ {

		wg.Add(1)
		go func(row int) {

			for col := 0; col < width; col++ {
				yIndex := row*width + col
				// 对于 U、V，每个值对应 2x2 像素块
				uIndex := (row/2)*(width/2) + (col / 2)
				vIndex := uIndex

				Y := yData[yIndex]
				U := uData[uIndex]
				V := vData[vIndex]

				R, G, B := yuvToRgb(Y, U, V)

				img.SetRGBA(col, row, color.RGBA{R, G, B, 255})
			}

			wg.Done()

		}(row)

	}

	wg.Wait()

	return img, nil
}

func yuvToRgb(Y, U, V byte) (R, G, B uint8) {
	y1 := float64(Y)
	u1 := float64(U) - 128
	v1 := float64(V) - 128

	r := y1 + 1.402*v1
	g := y1 - 0.344136*u1 - 0.714136*v1
	b := y1 + 1.772*u1

	R = clamp(r)
	G = clamp(g)
	B = clamp(b)
	return
}

func clamp(val float64) uint8 {
	if val < 0 {
		return 0
	}
	if val > 255 {
		return 255
	}
	return uint8(val)
}
