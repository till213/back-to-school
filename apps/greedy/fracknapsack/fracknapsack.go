package main

import (
	"fmt"
	"math"
	"sort"
)

const cashAmount = 7
const goldAmount = 3
const jewelryAmount = 1
const maxWeight = 50

type item struct {
	name string
	// the value of this item
	value float64
	// the weight of this item
	weight float64
	// the (fractional) number of available items
	amount float64
}

// either the bank or the knapsack where we "store"
// the (maximum value) items
type storage struct {
	// the items in this storage
	items     []item
	// the total weight of this storage
	weight float64
	maxWeight float64
}

func initBank(bank *storage) {
	var item1, item2, item3 item

	item1.name = "Cash"
	item1.value = 1000
	item1.weight = 3
	item1.amount = cashAmount

	item2.name = "Gold"
	item2.value = 1500
	item2.weight = 7
	item2.amount = goldAmount

	item3.name = "Jewelry"
	item3.value = 5000
	item3.weight = 11
	item3.amount = jewelryAmount

	bank.items = make([]item, 3)
	bank.items[0] = item1
	bank.items[1] = item2
	bank.items[2] = item3

	for _,v := range bank.items {
		bank.weight += v.weight * v.amount
	}
	// tresor has unlimited capacity
	bank.maxWeight = math.MaxFloat64
}

func initKnapsack(knapsack *storage) {
	knapsack.items = make([]item, 0)
	knapsack.weight = 0
	// knapsack has limited capacity
	knapsack.maxWeight = maxWeight
}

// Solves the fractional knapsack problem which allows to
// grab fractional amounts, in a greedy manner: sorts
// items in a "value per weight" fashion
func grabAndRun(bank, knapsack *storage) {	
	sort.Slice(bank.items, func(i, j int) bool {
		return bank.items[i].value / bank.items[i].weight > bank.items[j].value / bank.items[j].weight
	})

	availableWeight := knapsack.maxWeight
	
	for i := 0; i < len(bank.items) && availableWeight > 0.0; i++ {
		bankItem := bank.items[i]
	
		totalItemWeight := bankItem.weight * bankItem.amount
		// the minimum weight that still fits into the knapsack
		minWeight := math.Min(totalItemWeight, availableWeight)
		// the (theoretical) amount that would still fit
		amount := minWeight / bankItem.weight
		// take what's there, with the most valuable items first ("greedy")
		minAmount := math.Min(amount, bankItem.amount)
		takenItem := item{bankItem.name, bankItem.value, bankItem.weight, minAmount}
		// and it's gone :)
		bankItem.amount -= minAmount
		knapsack.items = append(knapsack.items, takenItem)

		knapsack.weight += minAmount * takenItem.weight
		availableWeight = knapsack.maxWeight - knapsack.weight
	}
}

func printStorage(storage *storage) {
	for _, v := range storage.items {
		fmt.Printf("Item: %s value: %.2f weight: %.2f value per weight: %.2f amount: %.2f\n", v.name, v.value, v.weight, v.value / v.weight, v.amount)
	}
	fmt.Println("Weight:", storage.weight, "Max weight:", storage.maxWeight)
}

func main() {
	var bank, knapsack storage
	initBank(&bank)
	initKnapsack(&knapsack)
	fmt.Println("--- Before ---")
	fmt.Println("--- Bank ---")
	printStorage(&bank)
	fmt.Println("--- Knapsack ---")
	printStorage(&knapsack)

	grabAndRun(&bank, &knapsack)

	fmt.Println("--- After ---")
	fmt.Println("--- Bank ---")
	printStorage(&bank)
	fmt.Println("--- Knapsack ---")
	printStorage(&knapsack)
}
