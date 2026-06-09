package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	testString := "6, 5, 4, 3, 2, 9"
	operType := operType()
	fmt.Printf("Выбранный тип операции: %s\n", operType)
	integerSlice := stringToSlice(testString)
	fmt.Println(integerSlice)
	fmt.Printf("an AVG for this slice is: %.2f\n", calculateAVG(integerSlice))
	fmt.Printf("a SUM for this slice is: %d\n", calculateSUM(integerSlice))
	fmt.Printf("a sorted slice is: %v\n", sliceSorter(integerSlice))
	fmt.Printf("a MED for this slice is: %.2f\n", calculateMED(integerSlice))
}

// func scanedString reads a user input into a string
func scanedString() (scanedString string) {
	for {
		fmt.Print("Введите последовательность чисел для рассчета:\n")
		fmt.Scan(&scanedString)
		if len(scanedString) == 0 {
			fmt.Println("Введенная строка должна содержать элементы")
		} else {
			break
		}
	}
	return scanedString
}

// operType function read a user input and return desired operation type
func operType() (operType string) {
	for {
		fmt.Print("Введите тип операции AVG/SUM/MED:")
		fmt.Scan(&operType)
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
func stringToSlice(stringToTransform string) (intSlice []int) {
	elements := strings.Split(stringToTransform, ",")
	intSlice = make([]int, len(elements))
	for i, elem := range elements {
		// Trim spaces if the string might look like "5, 12, 67"
		trimmed := strings.TrimSpace(elem)
		num, err := strconv.Atoi(trimmed)
		if err != nil {
			fmt.Println("Ошибка преобразования:", err)
			return
		}
		intSlice[i] = num
	}
	return intSlice
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
func calculateSUM(intSlice []int) (calculatedValue int) {
	for _, value := range intSlice {
		calculatedValue += value
	}
	return calculatedValue
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
		calculatedValue = float64((intSlice[(len(intSlice)/2)-1] + intSlice[(len(intSlice)/2)]) / 2)
	} else {
		calculatedValue = float64(intSlice[(len(intSlice) / 2)])
	}
	return calculatedValue
}
