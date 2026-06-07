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

	myArr := make([]int, 5, 10)

	fmt.Println(myArr[0])

}

func scanTransaction() float64 {
	var transaction float64
	fmt.Print("Введите сумму транзакции (0 для завершения): ")
	fmt.Scan(&transaction)
	return transaction
}

func banlanceCalculation(transactions []float64) (balance float64) {
	balance = 0.0
	for _, value := range transactions {
		balance += value
	}
	return balance
}
