package main

import (
	"fmt"
	"os"
	"strings"
	"unicode/utf8"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Нужно передать один аргумент: путь к файлу")
		return
	}

	path := os.Args[1]
	data, err := os.ReadFile(path)
	if err != nil {
		fmt.Println("Ошибка чтения файла:", err)
		return
	}

	text := string(data)
	chars := utf8.RuneCountInString(text)
	lines := strings.Count(text, "\n") + 1
	words := len(strings.Fields(text))

	fmt.Printf("--- Анализ файла: %s ---\n", path)
	fmt.Printf("Количество символов: %d\n", chars)
	fmt.Printf("Количество строк: %d\n", lines)
	fmt.Printf("Количество слов: %d\n", words)
}
