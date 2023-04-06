package main

import (
	"fmt"
	"math"
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
			charts(result)
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

	bar += "\tDay " + strconv.Itoa(yearday) + ".\t" + fmt.Sprintf("%.2f", porcentage) + "%"

	fmt.Println(bar)
	fmt.Println(time.Now().Date())
}

func parseFloats(args []string) ([]float64, bool) {
	stopError := false
	result := make([]float64, len(args))

	for i, arg := range args {
		if number, err := strconv.ParseFloat(arg, 64); err == nil {
			sign := math.Signbit(number) // check if negative

			if sign {
				stopError = true
				break
			}

			result[i] = number
		} else {
			stopError = true // error parsing to float64
			break
		}
	}

	return result, stopError
}

func charts(args []float64) {
	max := max(args)

	for _, arg := range args {
		porcentage := (arg / max) * float64(100)
		fmt.Println("|")
		bar := "|"

		for i := 0; i < 20; i++ {
			if i*5 > int(porcentage) {
				bar += "▒"
			} else {
				bar += "▓"
			}
		}

		bar += "\t" + fmt.Sprintf("%.1f", arg) + "\t" + fmt.Sprintf("%.2f", porcentage) + "%"
		fmt.Println(bar)
	}

	fmt.Println("---------------------------------------")
	fmt.Println("                    100%")
}

func max(args []float64) float64 {
	max := args[0]

	for _, number := range args {
		if number > max {
			max = number
		}
	}

	return max
}
