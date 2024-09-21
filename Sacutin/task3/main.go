package main

import "fmt"

func main() {
	fmt.Println("Назови себя сынок") // Вообще я хотел к 1 задаче, но 3 тоже об этом просит
	var (
		name string
		age  int
	)
	fmt.Scan(&name)
	fmt.Println("Имя понял, а шо по возврасту?")
	fmt.Scan(&age)
	fmt.Printf("Фу, тебя реально зовут %v, и тебе %v? Реально олух...", name, age)
	// ^ Вот тут меня кромешный пиздец был ошибки были потом вы выводило пробывл Scanf потом узнал что это фигня
}
