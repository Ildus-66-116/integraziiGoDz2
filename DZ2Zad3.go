package main

import (
	"bufio"
	"fmt"
	"os"
)

//Задание 3. Уровни доступа
//Что нужно сделать:
//Напишите программу, создающую текстовый файл только для чтения, и проверьте,
//что в него нельзя записать данные.

func main() {
	file, err := os.Create("some1.txt")
	if err := os.Chmod("some1.txt", 0444); err != nil {
		fmt.Println(err)
	}
	writer := bufio.NewWriter(file)
	if err != nil {
		fmt.Println("Не смогли записать файл, ", err)
		return
	}
	defer file.Close()
	writer.WriteString("Say hi")
}
