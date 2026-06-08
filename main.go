package main

import "fmt"

func main() {
	// transactions := []float64{}
	// for {
	// 	transaction := scanTransaction()
	// 	if transaction == 0 {
	// 		break
	// 	}
	// 	transactions = append(transactions, transaction)
	// }
	// fmt.Println("Список транзакций:", transactions)
	// balance := banlanceCalculation(transactions)
	// fmt.Printf("Ваш баланс = %.2f\n", balance)

	// fmt.Println()("Длинна: ",len(arr),"Вместимость: ", cap(arr))

	// myArr := make([]int, 5, 10)

	// fmt.Println(myArr[0])
	operType := operType()
	fmt.Printf("Выбранный тип операции: %s\n", operType)

}

func scanTransaction() float64 {
	var transaction float64
	fmt.Print("Введите сумму транзакции (0 для завершения): ")
	fmt.Scan(&transaction)
	return transaction
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
