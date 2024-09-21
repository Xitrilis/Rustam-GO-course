package main

import "fmt"

func main() {
	a := 32             // Ну там всякие int32 int64 я не трогал
	b := 1.99           // float64 - float32
	c := "ElGribi"      // string
	d := 32 + 10        // Я хотел bool показать, ну тип вроде то
	e := []int{1, 2, 3} // Массив, ток я забыл как резать нужно будет напомнить
	f := 42 == d
	fmt.Println(a, b, c)
	fmt.Println(e)
	if f {
		fmt.Println("Есть") // Это булл тип??? надеюсь да
	}
}