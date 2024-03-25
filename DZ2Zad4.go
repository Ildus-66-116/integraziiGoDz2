package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"os"
	"time"
)

// Задание 4. Пакет ioutil
//Что нужно сделать:
//-Перепишите задачи 1 и 2, используя пакет ioutil.

func main() {
	var b bytes.Buffer
	DateFormat := "2006-01-02 15:04:05"
	var n int = 1
	for {
		createdAt := time.Now()
		fmt.Println("Введите текст")
		var text string
		fmt.Scan(&text)
		if text == "exit" {
			fmt.Println("До свидания")
			break
		} else {
			b.WriteString(fmt.Sprintln(n, createdAt.Format(DateFormat), text))
			n++
			time.Sleep(1 * time.Second)
			fmt.Println("Данные запомнены")
		}
	}

	filename := "histori.txt"
	if err := ioutil.WriteFile(filename, b.Bytes(), 0666); err != nil {
		panic(err)
	}
	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	resultBytes, err := ioutil.ReadAll(file)
	if err != nil {
		panic(err)
	}
	fmt.Println("Сохраненый лог:")
	fmt.Println(string(resultBytes))
}
