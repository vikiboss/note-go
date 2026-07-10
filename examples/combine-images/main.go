package main

import (
	"image"
	"image/draw"
	"image/jpeg"
	"image/png"
	"net/http"
	"net/url"
)

func main() {
	http.HandleFunc("/", combineImages)
	http.ListenAndServe(":8080", nil)
}

func combineImages(w http.ResponseWriter, r *http.Request) {
	// 获取 URL 参数中指定的两个图片链接
	bgURL := r.URL.Query().Get("background")
	fgURL := r.URL.Query().Get("foreground")

	// 验证链接是否为有效 URL
	if !isValidURL(bgURL) || !isValidURL(fgURL) {
		http.Error(w, "Invalid URL provided", http.StatusBadRequest)
		return
	}

	// 下载背景图像并解码为 image.Image 类型
	bgImg := downloadImage(bgURL)

	// 下载前景图像并解码为 image.Image 类型
	fgImg := downloadImage(fgURL)

	// 合成图像
	resultImg := combine(bgImg, fgImg)

	// 编码为 PNG 格式并写入 HTTP 响应中
	w.Header().Set("Content-Type", "image/png")
	err := png.Encode(w, resultImg)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func downloadImage(url string) image.Image {
	// 发送 HTTP 请求并解码为 image.Image 类型
	resp, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	// 根据图像类型解码为 image.Image 类型
	if resp.Header.Get("Content-Type") == "image/jpeg" {
		img, err := jpeg.Decode(resp.Body)
		if err != nil {
			panic(err)
		}
		return img
	} else if resp.Header.Get("Content-Type") == "image/png" {
		img, err := png.Decode(resp.Body)
		if err != nil {
			panic(err)
		}
		return img
	} else {
		panic("unsupported image format")
	}
}

func combine(bgImg image.Image, fgImg image.Image) image.Image {
	// 调整前景图像大小为背景图像大小
	fgImg = resizeImage(bgImg.Bounds().Dx(), bgImg.Bounds().Dy(), fgImg)

	// 背景图像上绘制前景图像
	resultImg := image.NewRGBA(bgImg.Bounds())
	draw.Draw(resultImg, bgImg.Bounds(), bgImg, image.Point{0, 0}, draw.Src)
	draw.Draw(resultImg, fgImg.Bounds(), fgImg, image.Point{0, 0}, draw.Over)

	return resultImg
}

func resizeImage(width, height int, img image.Image) image.Image {
	// 调整图像大小为给定宽度和高度
	newImg := image.NewRGBA(image.Rect(0, 0, width, height))
	draw.Draw(newImg, newImg.Bounds(), img, img.Bounds().Min, draw.Src)
	return newImg
}

// 辅助函数：验证 URL 是否有效
func isValidURL(link string) bool {
	_, err := url.ParseRequestURI(link)
	return err == nil
}
