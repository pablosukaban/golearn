package exercises

import (
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
	"time"
)

// странца 65
func Exercise31() {
	greetings := []string{"Hello", "Hola", "नमस्कार", "こんにちは", "Привет"}
	firstTwo := greetings[:2]
	secThirdFourth := greetings[1:4]
	fourthFifth := greetings[3:]

	fmt.Println(firstTwo, secThirdFourth, fourthFifth)
}

func Exercise32() {
	message := "Hi 👨 and 👩"
	fmt.Println(len(message))
	fmt.Println(message[3:7])
}

func Exercise33() {
	type employee struct {
		firstName string
		lastName  string
		id        int
	}

	bob := employee{
		"bob",
		"bar",
		1,
	}
	tom := employee{
		firstName: "tom",
		lastName:  "foo",
		id:        2,
	}
	var zack struct {
		firstName string
		lastName  string
		id        int
	}
	zack.firstName = "zach"
	zack.lastName = "maz"
	zack.id = 3

	fmt.Println(bob, tom, zack)
}

/*
читает из stdin (или заранее из массива) список чисел,
считает минимальное, максимальное и среднее значение,
выводит результат с форматированием.
*/
func Exercise34() {
	args := os.Args

	if len(args) < 2 {
		fmt.Println("задайте число!")
		return
	}

	min := math.MaxInt
	max := math.MinInt
	var sum int = 0

	for _, val := range args[1:] {

		intVal, err := strconv.Atoi(val)

		if err != nil {
			fmt.Println("ошибка во время конвертации строки в инт!", err)
			continue
		}

		if intVal < min {
			min = intVal
		}

		if intVal > max {
			max = intVal
		}

		sum += intVal
	}

	var average = float64(sum) / float64(len(args[1:]))

	fmt.Printf("min: %d, max: %d, average: %.2f \n", min, max, average)
}

/* Программа принимает строку, разбивает на слова (через strings.Fields) и считает, сколько раз встречается каждое слово. */
func Exercise35(param string) {
	words := strings.Fields(param)

	if len(words) == 0 {
		fmt.Println("param передайте!")
		return
	}

	wordsCount := map[string]int{}

	for _, word := range words {
		lowered := strings.ToLower(word)
		trimmed := strings.Trim(lowered, ".,!?;")

		wordsCount[trimmed]++
	}

	for word, count := range wordsCount {
		fmt.Printf("%s - %d \n", word, count)
	}
}

func isDatesEqual(first, second time.Time) bool {
	return first.Year() == second.Year() && first.Month() == second.Month() && first.Day() == second.Day()
}

func Exercise36() {
	const (
		colorRed   = "\033[0;31m"
		colorNone  = "\033[0m"
		colorWhite = "\033[97m"
		colorGray  = "\033[37m"
	)

	args := os.Args[1:]

	if len(args) < 2 {
		fmt.Println("введите месяц и год")
		return
	}

	monthNumber, err := strconv.Atoi(args[0])

	if err != nil {
		fmt.Println("ошибка с месяцем", err)
		return
	}

	yearNumber, err := strconv.Atoi(args[1])

	if err != nil {
		fmt.Println("ошибка с годом", err)
		return
	}

	date := time.Date(yearNumber, time.Month(monthNumber), 1, 0, 0, 0, 0, time.Local)
	daysInMonth := date.AddDate(0, 1, -1).Day()

	currentDate := time.Now()

	fmt.Println(date.Month(), yearNumber)
	fmt.Println("Su Mo Tu We Th Fr Sa")

	for range int(date.Weekday()) {
		fmt.Printf("%2s ", "")
	}

	for i := 1; i <= daysInMonth; i++ {
		isToday := isDatesEqual(date, currentDate)
		weekDay := int(date.Weekday())

		color := colorWhite
		suffix := " "

		switch {
		case isToday:
			color = colorRed
		case weekDay == 0, weekDay == 6:
			color = colorNone
		}

		if weekDay == 6 {
			suffix = "\n"
		}

		fmt.Printf(color+"%2d"+suffix+colorNone, i)
		date = date.AddDate(0, 0, 1)
	}

	fmt.Println()
}
