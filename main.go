package main

import (
	"flag"
	"fmt"
	"os"
	"sort"
	"strings"
	"unicode/utf8"
)

func main() {
	topN := flag.Int("top", 5, "сколько слов показать в топе")
	onlyLines := flag.Bool("l", false, "вывести только количество строк")
	flag.Parse()

	args := flag.Args()
	if len(args) != 1 {
		fmt.Println("Нужно передать один аргумент: путь к файлу")
		return
	}

	path := args[0]
	data, err := os.ReadFile(path)
	if err != nil {
		fmt.Println("Ошибка чтения файла:", err)
		return
	}

	text := string(data)
	lines := strings.Count(text, "\n") + 1

	if *onlyLines {
		fmt.Printf("Количество строк: %d\n", lines)
		return
	}

	chars := utf8.RuneCountInString(text)
	wordList := strings.Fields(text)
	words := len(wordList)

	fmt.Printf("--- Анализ файла: %s ---\n", path)
	fmt.Printf("Количество символов: %d\n", chars)
	fmt.Printf("Количество строк: %d\n", lines)
	fmt.Printf("Количество слов: %d\n", words)

	freq := make(map[string]int)
	for _, word := range wordList {
		freq[word]++
	}

	type item struct {
		word  string
		count int
	}

	list := make([]item, 0)
	for word, count := range freq {
		list = append(list, item{word: word, count: count})
	}

	sort.Slice(list, func(i, j int) bool {
		return list[i].count > list[j].count
	})

	n := *topN
	if n > len(list) {
		n = len(list)
	}

	fmt.Printf("--- Топ-%d слов ---\n", n)
	for i := 0; i < n; i++ {
		fmt.Printf("%s: %d\n", list[i].word, list[i].count)
	}
}
