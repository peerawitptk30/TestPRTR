package main

import (
	"fmt"
	"strconv"
)

func encode(numerical string) string {
	encoded := ""
	for i := 0; i < len(numerical)-1; i++ {
		leftDigit, _ := strconv.Atoi(string(numerical[i]))
		rightDigit, _ := strconv.Atoi(string(numerical[i+1]))

		if leftDigit > rightDigit {
			encoded += "L"
		} else if leftDigit < rightDigit {
			encoded += "R"
		} else {
			encoded += "="
		}
	}

	return encoded
}

func decode(encoded string) string {
	result := ""
	prevSign := "="
	for _, symbol := range encoded {
		switch string(symbol) {
		case "L":
			if prevSign == "R" {
				result += "0"				
			}
			result += "1"
			prevSign = "L"

		case "R":
			if prevSign == "L" {
				result += "0"
			}
			result += "1"
			prevSign = "R"
			
		case "=":
			result += "0"
			prevSign = "="
		}

	}
	return result
}

func sumDigits(numerical string) int {
	sum := 0
	for _, digit := range numerical {
		digitInt, _ := strconv.Atoi(string(digit))
		sum += digitInt
	}
	return sum
}

func findMinimumSumEncoded(encodedStrings []string) string {
	minimumSum := -1
	minimumSumEncoded := ""
	for _, encoded := range encodedStrings {
		numerical := decode(encoded)
		sum := sumDigits(numerical)
		if minimumSum == -1 || sum < minimumSum {
			minimumSum = sum
			minimumSumEncoded = encoded
		}
	}
	return minimumSumEncoded
}

func main() {
	var numericalString string
	fmt.Print("Enter the numerical string: ")
	fmt.Scanln(&numericalString)
	_, err := strconv.Atoi(numericalString)
	if err != nil {
		fmt.Println("Input must be a number")
	} else {
		encodedString := encode(numericalString)
		minimumSumEncoded := findMinimumSumEncoded([]string{encodedString})
		fmt.Println("Encoded string with minimum sum of digits:", minimumSumEncoded)
	}

}
