package main

import (
	"fmt"
	"log"

	"github.com/Destroycode/keyboard"
)

func main() {
	fmt.Println("Проверка Git!")
	fmt.Print("Напишите градусы в Фаренгейтах: ")
	fareng, err := keyboard.Getfloat()
	if err != nil {
		log.Fatal()
	}
	celsii := (fareng - 32) * 5 / 9
	fmt.Printf("%0.1f градуса по Фаренгейту это %0.1f по Цельсию ", fareng, celsii)

	/*
		fmt.Println("Укажите ваши баллы: ")

		grade, err := keyboard.Getfloat()
		if err != nil {
			log.Fatal()
		}

		var status string
		if grade >= 60 {
			status = "Сдал."
		} else {
			status = "Не сдал."
		}
		fmt.Println("Твои баллы:", grade, "Ты", status)
	*/
}
