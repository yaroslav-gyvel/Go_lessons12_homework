package main

import (
	"fmt"
	"sync"
)

// Конкурентно порахувати суму кожного слайсу int, та роздрукувати результат.
// Потрібно використовувати WaitGroup.
// Приклад:
// [ [ 4, 6 ], [ 7, 9 ] ]
// Результат друку:
// Порядок друку не важливий.
// “slice 1: 10”
// “slice 2: 16”
func main() {
	var wg sync.WaitGroup

	n := [][]int{
		{2, 6, 9, 24},
		{7, 3, 94, 3, 0},
		{4, 2, 8, 35},
	}

	wg.Add(len(n))

	for counter, sl := range n {
		go sum(&wg, sl, counter)
	}

	wg.Wait()
}

func sum(wg *sync.WaitGroup, slice []int, c int) {
	defer wg.Done()
	var summa int
	for _, number := range slice {
		summa += number
	}
	fmt.Printf("slice %v: %v\n", c+1, summa)
}
