package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	var (
		name, famil, otchestvo, numberstud string
	)
	fmt.Print("Введите имя студента:")
	fmt.Scan(&name)
	fmt.Print("Введите фамилию студента:")
	fmt.Scan(&famil)
	fmt.Print("Введите отчетсво студента:")
	fmt.Scan(&otchestvo)
	fmt.Print("Введите номер группы студента:")
	fmt.Scan(&numberstud)

	var fio string
	var lab string
	var ready string

	lab = "Лабораторная работа №1"
	fio = name + " " + famil + " " + otchestvo
	ready = "Выполнил(а): ст. гр." + numberstud

	var zvezda []rune
	var maxLen int
	if utf8.RuneCountInString(fio) > utf8.RuneCountInString(lab) && utf8.RuneCountInString(fio) > utf8.RuneCountInString(ready) {
		maxLen = utf8.RuneCountInString(fio)
	} else if utf8.RuneCountInString(lab) > utf8.RuneCountInString(fio) && utf8.RuneCountInString(lab) > utf8.RuneCountInString(ready) {
		maxLen = utf8.RuneCountInString(lab)
	} else {
		maxLen = utf8.RuneCountInString(ready)
	}

	for i := 0; i < maxLen; i++ {

		zvezda = append(zvezda, '*')
	}
	var maxFio int
	maxFio = maxLen - utf8.RuneCountInString(fio)
	for i := 0; i < maxFio; i++ {
		fio += " "
	}
	var maxLab int
	maxLab = maxLen - utf8.RuneCountInString(lab)
	for i := 0; i < maxLab; i++ {
		lab += " "
	}
	var maxReady int
	maxReady = maxLen - utf8.RuneCountInString(ready)
	for i := 0; i < maxReady; i++ {
		ready += " "
	}
	fmt.Println("*" + string(zvezda) + "*")
	fmt.Println("*" + lab + "*")
	fmt.Println("*" + ready + "*")
	fmt.Println("*" + fio + "*")
	fmt.Println("*" + string(zvezda) + "*")
}
