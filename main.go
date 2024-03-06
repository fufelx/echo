package main

import (
	"fmt"
	"log"
)

func main() {
<<<<<<< HEAD
	n := ""
=======
	n := 0
>>>>>>> master
	fmt.Print("Введите целое число: ")
	_, err := fmt.Scan(&n)
	if err != nil {
		log.Fatal(err)
	}
<<<<<<< HEAD
	fmt.Printf("Вы ввели число: %v\n", n)
=======
	fmt.Printf("Вы ввели число: %d\n", n)
>>>>>>> master
}
