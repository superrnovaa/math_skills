package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"path/filepath"
	"sort"
	"strconv"
)

func main() {
	lengthofArg := 0
	for i := range os.Args {
		lengthofArg = i
	}
	if lengthofArg < 1 {
		fmt.Println("No file given to read from") //to ensure the data.txt is in the given arg
	} else if lengthofArg > 1 {
		fmt.Println("Too many arguments")
	} else {
		inputFile := os.Args[1]                //File with text to perform calculations                                                     //file that saves the text after manipulation
		if filepath.Ext(inputFile) != ".txt" { //To verify the arguments given is text file only
			fmt.Println("Only text file required!")
			os.Exit(1)
		}
		Stats(inputFile)
	}
}
func Stats(inputFile string) { //function to calculate average and print all the calculations
	count, sum := 0, 0
	var avg float64
	var result []int //resultant integer array
	var data []float64
	inputText, err := os.Open(inputFile) //reading the first file
	if err != nil {
		fmt.Println("Error opening input file: ", err.Error())
		os.Exit(1)
	}
	scanner := bufio.NewScanner(inputText) //reading input file's text
	for scanner.Scan() {
		num := scanner.Text()
		i, err := strconv.Atoi(num)
		if err != nil {
			fmt.Println("Error converting string to int")
		}
		result = append(result, i)
		data = append(data, float64(i))
		count = count + 1
	}
	for i := 0; i <= count-1; i++ {
		sum += result[i]
	}
	avg = float64(sum) / float64(count)
	avg = math.Round(avg)
	fmt.Println("Average:", avg)
	fmt.Println("Median:", median(data))
	v := math.Abs(variance(data))

	if v >= 1 {
		fmt.Printf("Variance: %d\n", int(variance(data)))
	} else {
		fmt.Printf("Variance: %.6\n", variance(data))
	}
	fmt.Println("Variance:", variance(data))
	fmt.Println("Standard Deviation:", standardDeviation(data))
}
func median(data []float64) float64 {
	dataCopy := make([]float64, len(data))
	copy(dataCopy, data)
	sort.Float64s(dataCopy)
	var median float64
	l := len(dataCopy)
	if l == 0 {
		return 0
	} else if l%2 == 0 {
		median = (dataCopy[l/2-1] + dataCopy[l/2]) / 2
	} else {
		median = dataCopy[l/2]
	}
	return math.Round(median)
}
func mean(nums []float64) float64 {
	var s float64 = 0
	for _, x := range nums {
		s += x
	}
	return s / float64(len(nums))
}
func variance(nums []float64) float64 {
	var res float64
	var m = mean(nums)
	var n = len(nums)
	for _, x := range nums {
		res += (x - m) * (x - m)
	}
	return math.Round(res / float64(n))
}
func standardDeviation(num []float64) float64 {
	var m, sd float64
	m = mean(num)
	n := len(num)
	for j := 0; j < n; j++ {
		sd += math.Pow(num[j]-m, 2)
	}
	sd = math.Sqrt(sd / float64(n))
	return math.Round(sd)
}
