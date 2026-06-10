package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	var calcResults float64
	var operType string
	var integerSlice []int
	var testString string

	for {
		operType = readOperType()
		testString, err := scanedString()
		integerSlice, err = stringToSlice(testString)
		if err == nil {
			break
		} else {
			fmt.Println("Некорректный ввод чисел для рассчета")
		}
	}
	fmt.Println(testString)
	fmt.Println(integerSlice)

	switch operType {
	case "AVG":
		calcResults = calculateAVG(integerSlice)
	case "SUM":
		calcResults = calculateSUM(integerSlice)
	case "MED":
		calcResults = calculateMED(integerSlice)
	}
	fmt.Printf("Результат вычисления %s равен: %.2f\n", operType, calcResults)

}

// func scanedString reads a user input into a string
func scanedString() (scanedString string, err error) {
	fmt.Print("Введите последовательность чисел для рассчета:\n")
	reader := bufio.NewReader(os.Stdin)
	// Read until the user presses Enter (\n)
	scanedString, err = reader.ReadString('\n')
	if err != nil {
		fmt.Println("Error reading input:", err)
		return
	}
	// 	// scanedString = "2, 3, 4, 5, 6, 7, 8"
	// fmt.Printf("I've just read this tring: %s\n", scanedString)
	return scanedString, err
}

// operType function read a user input and return desired operation type
func readOperType() (operType string) {
	for {
		fmt.Print("Введите тип операции AVG/SUM/MED:")
		fmt.Scan(&operType)
		// operType = "SUM"
		if operType == "AVG" || operType == "SUM" || operType == "MED" {
			break
		} else {
			fmt.Println("Ошибка при выборе типа операции")
			fmt.Println("Повторите свой выбор")
		}
	}
	return operType
}

// stringToSlice gets a string and transform it into a slice of intiger
func stringToSlice(stringToTransform string) ([]int, error) {
	// Remove all spaces, tabs, and newlines from the string
	fmt.Printf("the string before transformation is: %s", stringToTransform)
	cleaned := strings.ReplaceAll(stringToTransform, " ", "")
	cleaned = strings.ReplaceAll(cleaned, "\t", "")
	cleaned = strings.ReplaceAll(cleaned, "\n", "")
	cleaned = strings.ReplaceAll(cleaned, "\r", "")
	fmt.Printf("the string before conversion is: %s", cleaned)

	// If the cleaned string is empty, return an empty slice immediately
	if cleaned == "" {
		return []int{}, nil
	}

	// Split the string into a slice of string tokens by the comma delimiter
	parts := strings.Split(cleaned, ",")
	intSlice := make([]int, 0, len(parts))

	// Iterate through each token and parse it into an integer
	for _, part := range parts {
		// Handle trailing or double commas which create empty elements
		if part == "" {
			continue
		}

		num, err := strconv.Atoi(part)
		if err != nil {
			return nil, fmt.Errorf("failed to parse %q as an integer: %w", part, err)
		}
		intSlice = append(intSlice, num)
	}

	return intSlice, nil

	// elements := strings.Split(stringToTransform, ",")
	// intSlice = make([]int, len(elements))
	// for i, elem := range elements {
	// 	// Trim spaces if the string might look like "5, 12, 67"
	// 	trimmed := strings.TrimSpace(elem)
	// 	num, err := strconv.Atoi(trimmed)
	// 	if err != nil {
	// 		fmt.Println("Ошибка преобразования:", err)
	// 		return
	// 	}
	// 	intSlice[i] = num
	// }
	// return intSlice
}

// calculateAVG function gets a slice of integer numbers and calculate an average for the slice
func calculateAVG(intSlice []int) (calculatedValue float64) {
	for _, value := range intSlice {
		calculatedValue += float64(value)
	}
	calculatedValue = calculatedValue / float64(len(intSlice))
	return calculatedValue
}

// calculateSUM function gets a slice of integer numbers and calculate a SUMM for the slice
func calculateSUM(intSlice []int) float64 {
	var calculatedValue int
	for _, value := range intSlice {
		calculatedValue += value
	}
	return float64(calculatedValue)
}

// sliceSorter gets a slice of integer, sort it, and returns a sorted slice
func sliceSorter(intSlice []int) (sortedSlice []int) {
	var intdigit int
	sliceLength := len(intSlice)
	for i := 0; i < (sliceLength - 1); i++ {
		for j := 0; j < (sliceLength - 1); j++ {
			intdigit = intSlice[j]
			if intdigit > intSlice[j+1] {
				intSlice[j] = intSlice[j+1]
				intSlice[j+1] = intdigit
			}

		}
	}
	return intSlice
}

// calculateMED function a slice of integer numbers and calculate a median for the slice
func calculateMED(intSlice []int) (calculatedValue float64) {
	intSlice = sliceSorter(intSlice)
	if len(intSlice)%2 == 0 {
		calculatedValue = (float64(intSlice[(len(intSlice)/2)-1]) + float64(intSlice[(len(intSlice)/2)])) / 2
	} else {
		calculatedValue = float64(intSlice[(len(intSlice) / 2)])
	}
	return calculatedValue
}
