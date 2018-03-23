package main

import (
	"github.com/golang/freetype"
	"github.com/golang/freetype/truetype"
	"golang.org/x/image/font"
	"golang.org/x/image/math/fixed"
	"image"
	"image/color"
	"image/draw"
	"io/ioutil"
	"math"
	"math/rand"
)

var white = color.RGBA{255, 255, 255, 255}

func initFont() (*truetype.Font, error) {
	fontBytes, err := ioutil.ReadFile("./font.ttf")
	if err != nil {
		return nil, err
	}

	defaultFont, err := freetype.ParseFont(fontBytes)
	if err != nil {
		return nil, err
	}

	return defaultFont, nil
}

func randRGBValue() uint8 {
	return uint8(rand.Intn(256))
}

func randomColor() color.RGBA {
	return color.RGBA{randRGBValue(), randRGBValue(), randRGBValue(), 255}
}

func generateImage(text string, imageSize int) (image.Image, error) {
	defaultFont, err := initFont()
	if err != nil {
		return nil, err
	}

	foregroundColor := white
	fontSize := float64(imageSize) / 2
	img := image.NewRGBA(image.Rect(0, 0, imageSize, imageSize))

	draw.Draw(img, img.Bounds(), &image.Uniform{randomColor()}, image.ZP, draw.Src)

	d := &font.Drawer{
		Dst: img,
		Src: image.NewUniform(foregroundColor),
		Face: truetype.NewFace(defaultFont, &truetype.Options{
			Size: fontSize,
			DPI:  72,
		}),
	}
	d.Dot = fixed.Point26_6{
		X: (fixed.I(imageSize) - d.MeasureString(text)) / 2,
		Y: fixed.I(int(math.Ceil(fontSize * 1.35))),
	}
	d.DrawString(text)

	return img, nil
}
