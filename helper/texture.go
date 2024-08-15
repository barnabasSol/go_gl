package helper

import (
	"image/jpeg"
	"image/png"
	"os"

	"github.com/go-gl/gl/v3.3-core/gl"
)

func LoadTextureAlphaJpeg(file_path string) TextureId {
	infile, os_err := os.Open(file_path)
	if os_err != nil {
		panic(os_err)
	}
	defer infile.Close()
	img, jpg_err := jpeg.Decode(infile)
	if jpg_err != nil {
		panic(jpg_err)
	}
	w := img.Bounds().Bounds().Max.X
	h := img.Bounds().Bounds().Max.Y

	pixels := make([]byte, w*h*4)
	bIndex := 0

	for y := 0; y < h; y++ {
		for x := 0; x < w; x++ {
			r, g, b, a := img.At(x, y).RGBA()
			pixels[bIndex] = byte(r / 256)
			bIndex++
			pixels[bIndex] = byte(g / 256)
			bIndex++
			pixels[bIndex] = byte(b / 256)
			bIndex++
			pixels[bIndex] = byte(a / 256)
			bIndex++
		}
	}
	texture := GenBindTexture()
	gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_WRAP_S, gl.REPEAT)
	gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_WRAP_S, gl.REPEAT)
	gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_MIN_FILTER, gl.LINEAR)
	gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_MAG_FILTER, gl.LINEAR)
	gl.TexImage2D(gl.TEXTURE_2D, 0, gl.RGBA, int32(w), int32(h), 0, gl.RGBA, gl.UNSIGNED_BYTE, gl.Ptr(pixels))
	gl.GenerateMipmap(gl.TEXTURE_2D)

	return texture
}

func LoadTextureAlphaPng(file_path string) TextureId {
	infile, os_err := os.Open(file_path)
	if os_err != nil {
		panic(os_err)
	}
	defer infile.Close()
	img, jpg_err := png.Decode(infile)
	if jpg_err != nil {
		panic(jpg_err)
	}
	w := img.Bounds().Bounds().Max.X
	h := img.Bounds().Bounds().Max.Y

	pixels := make([]byte, w*h*4)
	bIndex := 0

	for y := 0; y < h; y++ {
		for x := 0; x < w; x++ {
			r, g, b, a := img.At(x, y).RGBA()
			pixels[bIndex] = byte(r / 256)
			bIndex++
			pixels[bIndex] = byte(g / 256)
			bIndex++
			pixels[bIndex] = byte(b / 256)
			bIndex++
			pixels[bIndex] = byte(a / 256)
			bIndex++
		}
	}
	texture := GenBindTexture()
	gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_WRAP_S, gl.REPEAT)
	gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_WRAP_S, gl.REPEAT)
	gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_MIN_FILTER, gl.LINEAR)
	gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_MAG_FILTER, gl.LINEAR)
	gl.TexImage2D(gl.TEXTURE_2D, 0, gl.RGBA, int32(w), int32(h), 0, gl.RGBA, gl.UNSIGNED_BYTE, gl.Ptr(pixels))
	gl.GenerateMipmap(gl.TEXTURE_2D)

	return texture
}

func GenBindTexture() TextureId {
	var textId uint32
	gl.GenTextures(1, &textId)
	gl.BindTexture(gl.TEXTURE_2D, textId)
	return TextureId(textId)
}

func BindTexture(id TextureId) {
	gl.BindTexture(gl.TEXTURE_2D, uint32(id))
}
