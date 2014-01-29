package main

import "os"
import "io"
import "fmt"
import "strconv"
import "image"
import _ "image/png"
import _ "image/color"
import "encoding/csv"


var colors = [][3]uint8{}

const TermWidth = 80


func init_colors(csv_file string) {
	file, _ := os.Open(csv_file)
	defer file.Close()
	reader := csv.NewReader(file)
	for {
		record, err := reader.Read()
		if err == io.EOF {
			break
		}
		r, _ := strconv.Atoi(record[0])
		g, _ := strconv.Atoi(record[1])
		b, _ := strconv.Atoi(record[2])
		colors = append(colors, [3]uint8{uint8(r), uint8(g), uint8(b)})
	}

}


func closest_color(r, g, b uint8) (int) {
	
	var min_sum int = 200000
	var closest_color = 0
	var yuv_coeff = [3]float32{0.299, 0.587, 0.114}
	for i := 0; i<256; i++ {
		var sum = 0
		var block = [3]uint8{r, g, b}
		for j := 0; j<3; j++ {
			value := int(float32((colors[i][j]-block[j]))*yuv_coeff[j])
			sum += value*value
		}
		if sum < min_sum {
			closest_color = i
			min_sum = sum
		}
	}
	
	return closest_color
}

//Print usage and exit program
func usage(prog_name string) {
	fmt.Println("Try to display a png image into your terminal.")
	fmt.Println("Notice: this program was intended for Golang practice only!\n")
	fmt.Println("Usage:", prog_name, "<path_to_png_file> [width]\n")
	os.Exit(2)
}


func main() {

	if len(os.Args) < 2 {
		usage(os.Args[0])
		os.Exit(2)
	}

	var termWidth = TermWidth
	if len(os.Args) == 3 {
		termWidth, _ = strconv.Atoi(os.Args[2])
	}

	init_colors("rgb.csv")
	
	fImg, _ := os.Open(os.Args[1])
	defer fImg.Close()
	img, _, _ := image.Decode(fImg)

	bounds := img.Bounds()
	w, h := bounds.Max.X, bounds.Max.Y

	var termHeight = int(termWidth * h / w)
	var blockWidth, blockHeight int
	blockWidth = (w / termWidth)
	blockHeight = (h / termHeight)

	//Consider each pixel in each block
	for j := 0; j < termHeight; j++ {
		for i := 0; i < termWidth; i++ {	
			var r, g, b int = 0, 0, 0
			//Compute the rgb mean value in this block
			for y := (j*blockHeight); y < (j+1)*blockHeight; y++ {
				for x := (i*blockWidth); x < (i+1)*blockWidth; x++ {
					tmpr, tmpg, tmpb, _ := img.At(x, y).RGBA()
					r += int(uint8(tmpr>>8))
					g += int(uint8(tmpg>>8))
					b += int(uint8(tmpb>>8))
				}
			}
			r = r/(blockWidth*blockHeight)
			g = g/(blockWidth*blockHeight)
			b = b/(blockWidth*blockHeight)

			//Get the 'closest' shell available color
			color := closest_color(uint8(r), uint8(g), uint8(b))
			//And print something with it
			fmt.Print("\033[48;5;", strconv.Itoa(color), "m \033[m")
		}
		fmt.Print("\n")
	}

}
