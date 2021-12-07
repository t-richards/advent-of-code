package data

import (
	"bufio"
	"io"
	"log"
	"os"
	"strconv"
	"strings"
)

func LoadInts() []int {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var input []int

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		num, err := strconv.Atoi(scanner.Text())
		if err != nil {
			panic(err)
		}
		input = append(input, num)
	}

	return input
}

func LoadSplitInts() []int {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var input []int
	data, _ := io.ReadAll(file)

	parts := strings.FieldsFunc(string(data), func(r rune) bool {
		return r == ',' || r == '\n'
	})
	for _, p := range parts {
		num, err := strconv.Atoi(p)
		if err != nil {
			panic(err)
		}
		input = append(input, num)
	}

	return input
}

func LoadStrings() []string {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var input []string

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		input = append(input, scanner.Text())
	}

	return input
}
