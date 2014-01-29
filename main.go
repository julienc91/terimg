package main

import "os"
import "io"
import "fmt"
import "strconv"
import "image"
import _ "image/png"
import _ "image/jpeg"
import _ "image/color"
import "encoding/csv"


var colors = [][3]uint8{}

const RgbConf = "rgb.csv"
const NbColors = 256
const TermWidth = 80


//Retrieve RGB values from the csv file
func init_colors(csv_file string) {
	
	file, err := os.Open(csv_file)
	if err != nil {
		fmt.Println("Could not open file", csv_file)
		os.Exit(2)
	}
	defer file.Close()
	reader := csv.NewReader(file)
	
	var count_lines int = 0
	for {
		record, err := reader.Read()
		if err == io.EOF {
			break
		} else if err != nil {
			fmt.Println("Error while parsing", csv_file)
			os.Exit(2)
		}
		var rgb [3]uint8
		for i := 0; i < 3; i++ {
			tmp, err := strconv.Atoi(record[i])
			if err != nil {
				fmt.Println("Error while parsing", csv_file, ":", record[i], "is not a correct value")
				os.Exit(2)
			}
			rgb[i] = uint8(tmp)
		}
		colors = append(colors, rgb)
		count_lines ++
	}
	if count_lines != 256 {
		fmt.Println("Error while parsing", csv_file, ": there should be", NbColors, "lines, but", count_lines, "were found")
		os.Exit(2)
	}

}


//Find the closest 8bit color
func closest_color(r, g, b uint8) (int) {

	var yuv_coeff = [3]float32{0.299, 0.587, 0.114}
	var min_sum, closest_color int = 200000, 0
	
	for i := 0; i < NbColors; i++ {
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


//Main function: read the image and print the 8bit version of it
func main() {

	if len(os.Args) < 2 {
		usage(os.Args[0])
		os.Exit(2)
	}

	var termWidth = TermWidth
	if len(os.Args) == 3 {
		tmp, err := strconv.Atoi(os.Args[2])
		if err != nil || tmp <= 0{
			fmt.Println("Optionnal parameter [TermWidth] must be a positive integer")
			os.Exit(2)
		}
		termWidth = tmp
	}

	init_colors(RgbConf)
	
	fImg, err := os.Open(os.Args[1])
	if err != nil {
		fmt.Println("Could not open the file", os.Args[1])
		os.Exit(2)
	}
	defer fImg.Close()
	
	img, _, err := image.Decode(fImg)
	if err != nil {
		fmt.Println("Could not decode image", os.Args[1])
		os.Exit(2)
	}

	bounds := img.Bounds()
	w, h := bounds.Max.X, bounds.Max.Y
	
	var termHeight = int(termWidth * h / w)
	if termHeight > h || termHeight <= 0 || termWidth > w {
		fmt.Println("The image is too small. Specify a smaller [TermWidth] value (maximum", w, ")")
		os.Exit(2)
	}
	
	var blockWidth, blockHeight int = w/termWidth, h/termHeight

	//Consider each pixel in each block
	var block_surface = blockWidth*blockHeight
	for j := 0; j < termHeight; j++ {
		for i := 0; i < termWidth; i++ {	
			var rgb = [3]int{0, 0, 0}
			//Compute the rgb mean value in this block
			for y := (j*blockHeight); y < (j+1)*blockHeight; y++ {
				for x := (i*blockWidth); x < (i+1)*blockWidth; x++ {
					var tmp = [3]uint32{}
					tmp[0], tmp[1], tmp[2], _ = img.At(x, y).RGBA()
					for z := 0; z < 3; z++ {
						rgb[z] += int(uint8(tmp[z]>>8))
					}
				}
			}
			for z := 0; z < 3; z++ {
				rgb[z] = rgb[z]/block_surface
			}

			//Get the 'closest' shell available color
			color := closest_color(uint8(rgb[0]), uint8(rgb[1]), uint8(rgb[2]))
			//And print something with it
			fmt.Print("\033[48;5;", strconv.Itoa(color), "m \033[m")
		}
		fmt.Print("\n")
	}

}
