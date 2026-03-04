package main

import (
	"fmt"
	"math/rand"
)

func main() {
	var code string
	found := false

	for i := 0; i < 5; i++ {
		fmt.Println("Введи значение")
		fmt.Scanf("%s\n", &code)
		// Проверка на число 300
		if code == "300" || code == "3 0 0" || code == "тристо" || code == "триста" {
			fmt.Println("Привет, отсоси у тракториста")
			found = true
			break
		}
		// Случайный выбор одной из двух фраз
		if rand.Intn(2) == 0 {
			fmt.Println("не то")
		} else {
			fmt.Println("попробуй ещё")
		}
	}

	if !found {
		fmt.Println("Ты не угадал за 5 попыток, лох! Пока!")
	}
}
