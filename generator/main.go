package main

import (
	"fmt"
	"image/color"
	"image/png"
	"log"
	"os"
	"strings"
	"text/template"
)

func main() {
	GenerateLogoRGBAFile("generator/images/1.png", 1)
	GenerateLogoRGBAFile("generator/images/2.png", 2)
}

func GenerateLogoRGBAFile(filepath string, index int) {
	colors := generateLogoRGBA(filepath)
	colorsStr := convertToString(colors)
	generateFile(colorsStr, index)
}

func generateLogoRGBA(filepath string) []color.RGBA {
	file, _ := os.Open(filepath)
	img, err := png.Decode(file)
	if err != nil {
		log.Fatal("failed to decode image file", err)
	}

	colors := make([]color.RGBA, 0)

	for y := 0; y < img.Bounds().Max.Y; y++ {
		for x := 0; x < img.Bounds().Max.X; x++ {
			r, g, b, _ := img.At(x, y).RGBA()
			colors = append(colors, color.RGBA{R: uint8(r >> 8), G: uint8(g >> 8), B: uint8(b >> 8), A: uint8(255)})
		}
	}

	return colors
}

func convertToString(colors []color.RGBA) string {
	var content strings.Builder
	var err error

	for i, c := range colors {
		str := fmt.Sprintf("{%d, %d, %d, %d}", c.R, c.G, c.B, c.A)
		_, err = content.WriteString(str)
		if err != nil {
			log.Fatal("failed to write string")
		}

		if i < len(colors)-1 {
			_, err = content.WriteString(", ")
			if err != nil {
				log.Fatal("failed to write string")
			}
		}
	}

	return content.String()
}

func generateFile(colorsStr string, index int) {
	tmp, err := template.ParseFiles("./generator/logo-template" + fmt.Sprintf("%v", index) + ".txt")
	if err != nil {
		log.Fatal("failed to parse template", err)
	}

	f, err := os.Create("./logo" + fmt.Sprintf("%v", index) + ".go")
	if err != nil {
		log.Fatal("failed to create output file", err)
	}

	type Colors struct {
		LogoRGBA string
	}

	err = tmp.Execute(f, Colors{LogoRGBA: colorsStr})
	if err != nil {
		log.Fatal("failed to execute template", err)
	}
}
