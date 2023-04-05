package main

import (
	"fmt"
	"os"
	"time"
)

func main() {
	args := os.Args[1:]

	if args[0] == "yearbar" {
		yearbar()
	} else {
		result, err := parseFloats(args)
		if err == nil {
			fmt.Println(result)
			fmt.Println(len(result))
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

func parseFloats(args []string) ([]float64, error) {
	result := make([]float64, len(args))
	fmt.Println(result)
	return result, nil
}
