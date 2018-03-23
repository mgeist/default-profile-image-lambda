package main

import (
	"image"
	"testing"
)

func TestRandRGBValue(t *testing.T) {
	randRGB := randRGBValue()
	randRGB2 := randRGBValue()

	if randRGB == randRGB2 {
		t.Error("Expected", randRGB, "to be different from", randRGB2)
	}
}

func TestRandomColor(t *testing.T) {
	randColor := randomColor()
	randColor2 := randomColor()

	if randColor == randColor2 {
		t.Error("Expected", randColor, "to be different from", randColor2)
	}
}

func TestGenerateImage(t *testing.T) {
	img, _ := generateImage("AB", 50)
	img2, _ := generateImage("AB", 50)
	rect50 := image.Rect(0, 0, 50, 50)

	if img.Bounds() != rect50 {
		t.Error("Expected image to be of size", rect50, ", got", img.Bounds())
	}

	if img == img2 {
		t.Error("Expected different images, got identical")
	}
}
