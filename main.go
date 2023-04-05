package main

import (
	"fmt"
	"os"
	"strconv"
	"time"
)

func main() {
	args := os.Args[1:]

	if args[0] == "yearbar" {
		yearbar()
	} else {
		result, stopError := parseFloats(args)
		if !stopError {
			fmt.Println(result)
			fmt.Println(len(result))
		} else {
			fmt.Println("something went bad")
		}
	}
}

func yearbar() {
	yearday := time.Now().YearDay()
	porcentage := (float64(yearday) / float64(365)) * float64(100)
	bar := ""

	for i := 0; i < 20; i++ {
		if i*5 > int(porcentage) {
			bar += "▒"
		} else {
			bar += "▓"
		}
	}

	bar += " " + fmt.Sprintf("%.2f", porcentage) + "%"

	fmt.Println(bar)
}

func parseFloats(args []string) ([]float64, bool) {
	stopError := false
	result := make([]float64, len(args))
	for i, arg := range args {
		if number, err := strconv.ParseFloat(arg, 64); err == nil {
			result[i] = number
		} else {
			stopError = true
			break
		}
	}
	return result, stopError
}
