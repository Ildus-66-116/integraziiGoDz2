package main

import (
	"fmt"
	"os"
	"time"
)

//Задание 1. Работа с файлами**
//Что нужно сделать:
//-Напишите программу, которая на вход получала бы строку, введённую пользователем,
//а в файл писала № строки, дату и сообщение в формате:
//2020-02-10 15:00:00 продам гараж.
//-При вводе слова exit программа завершает работу.

func main() {
	file, err := os.Create("histori.txt")
	if err != nil {
		fmt.Println("Не смогли создать файл, ", err)
		return
	}
	defer file.Close()
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
			file.WriteString(fmt.Sprintln(n, createdAt.Format(DateFormat), text))
			n++
			time.Sleep(1 * time.Second)
			fmt.Println("Данные записаны")
		}
	}
}
