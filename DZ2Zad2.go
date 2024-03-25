package main

import (
	"fmt"
	"io"
	"os"
)

//Задание 2. Интерфейс io.Reader
//Что нужно сделать:
//-Напишите программу, которая читает и выводит в консоль строки из файла,
//созданного в предыдущей практике, без использования ioutil.
//Если файл отсутствует или пуст, выведите в консоль соответствующее сообщение.

func main() {
	f, err := os.Open("histori.txt")
	if err != nil {
		fmt.Println("Не смогли открыть файл, ", err)
		return
	}
	defer f.Close()

	stat, err := f.Stat()
	if err != nil {
		panic(err)
	}
	size := stat.Size()

	buf := make([]byte, size)
	if _, err := io.ReadFull(f, buf); err != nil {
		panic(err)
	}
	fmt.Printf("%s\n", buf)
}
