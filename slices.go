package main

import (
"fmt"
"math/rand"
"reflect"
"time"
)

type Slices struct {
	s		[]string
	i		[]int
	ui		[]uint64
}

var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

func randInt(min int, max int) int {
	return min + rand.Intn(max-min)
}

func getString() []string {
	strLen := randInt(1, 10)

	str := make([]string, strLen)

	for i := range str {
		runeIndex := randInt(0, 52)
		str[i] = string(letterRunes[runeIndex])
	}

	return str
}

func getInt() (numbers[] int) {
	intLen := randInt(10, 20)

	numbers = make([]int, intLen)

	for i := range numbers {
		numbers[i] = randInt(10, 20)
	}

	return
}

func getUint() (uints[] uint64) {
	uintLen := randInt(20, 30)

	uints = make([]uint64, uintLen)

	for i := range uints {
		uints[i] = rand.Uint64()
	}

	return
}

func initSlices() Slices {
	sl := Slices {}

	sl.s = getString()
	sl.i = getInt()
	sl.ui = getUint()

	return sl
}

func printSlicest(s Slices) {
	fields := reflect.TypeOf(s)
	values := reflect.ValueOf(s)

	num := fields.NumField()

	for i := 0; i < num; i++ {
		field := fields.Field(i)
		value := values.Field(i)
		fmt.Print("Type: ", field.Type, ", ", field.Name, " = ", value, "\n")
	}
}

func main()  {
	rand.Seed(time.Now().UTC().UnixNano())

	slices := initSlices()

	printSlicest(slices)
}
