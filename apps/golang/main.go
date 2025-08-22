package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
)

func AddSix(i float64) float64 {
	return i + 6
}

func main() {
	if len(os.Args) < 2 {
		log.Fatal("Please provide a number as an argument")
	}

	num, err := strconv.ParseFloat(os.Args[1], 64)
	if err != nil {
		log.Fatal("Error:", err)
	}

	result := AddSix(num)
	fmt.Println(result)
}
