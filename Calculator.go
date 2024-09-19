package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

var swth int

func main() {

	fmt.Println("Нажмите r для старта/рестарта, нажмите q для выхода")
	scanners := bufio.NewScanner(os.Stdin)
	for scanners.Scan() {
		exit := scanners.Text()
		if exit == "q" {
			break
		} else {

			scanner := bufio.NewScanner(os.Stdin)
			fmt.Print("Введите выражение: ")
			scanner.Scan()

			str := scanner.Text()
			str = strings.Join(strings.Fields(str), "")

			f := math(str, transform(rasp(pars(str)), pars(str)))

			if swth != 2 {

				fmt.Println(f)

			}
			if swth == 2 {

				translateToRome(f)

			}

			fmt.Println("Нажмите r для рестарта, нажмите q для выхода")
		}
	}

}

func pars(str string) []string {

	if strings.Contains(str, "+") {
		arr := strings.Split(str, "+")
		return arr
	}
	if strings.Contains(str, "-") {
		arr := strings.Split(str, "-")
		return arr
	}
	if strings.Contains(str, "*") {
		arr := strings.Split(str, "*")
		return arr
	}
	if strings.Contains(str, "/") {
		arr := strings.Split(str, "/")
		return arr
	} else {
		arr := strings.Split(str, "")
		return arr
	}

}

func rasp(arr []string) []bool {

	rome := [10]string{"I", "II", "III", "IV", "V", "VI", "VII", "VIII", "IX", "X"}
	retArr := []bool{false, false}

	d := false

	for i := len(arr); i > 0; i-- {

		for c := len(rome); c > 0; c-- {

			if arr[i-1] == rome[c-1] {
				d = true
				retArr[i-1] = d
				break
			} else {
				d = false
				retArr[i-1] = d
			}

		}

	}

	return retArr

}

func transform(arrB []bool, arr []string) []int {

	RomeToArab := map[string]int{
		"I":   1,
		"II":  2,
		"III": 3,
		"IV":  4,
		"V":   5,
		"VI":  6,
		"VII": 7,
		"IX":  9,
		"X":   10,
	}

	if arrB[0] == false && arrB[1] == false {

		x, err1 := strconv.Atoi(arr[0])
		if err1 != nil {
			log.Fatal(err1)
		}

		y, err2 := strconv.Atoi(arr[1])
		if err2 != nil {
			log.Fatal(err2)
		}

		arrC := []int{x, y}
		return arrC

	}

	if arrB[0] == true && arrB[1] == true {

		x := RomeToArab[arr[0]]
		y := RomeToArab[arr[1]]

		arrC := []int{x, y}
		swth = 2
		return arrC

	} else {
		x := 0
		y := 0
		arrC := []int{x, y}
		return arrC
	}

}

func math(str string, arrC []int) int {

	if strings.Contains(str, "+") {

		x := arrC[0]
		y := arrC[1]
		c := x + y

		return c

	}
	if strings.Contains(str, "-") {

		x := arrC[0]
		y := arrC[1]
		c := x - y

		return c

	}
	if strings.Contains(str, "*") {

		x := arrC[0]
		y := arrC[1]
		c := x * y

		return c

	}
	if strings.Contains(str, "/") {

		x := arrC[0]
		y := arrC[1]
		c := x / y

		return c

	} else {
		return 0
	}

}

func translateToRome(x int) {

	var arr [4]int

	arr[0] = x % 10
	arr[1] = (x % 100) / 10
	arr[2] = (x % 1000) / 100
	arr[3] = x / 1000

	arrOne := []string{" ", "I", "II", "III", "IV", "V", "VI", "VII", "VIII", "IX"}

	arrTen := []string{" ", "X", "XX", "XXX", "XL", "L", "LX", "LXX", "LXXX", "XC"}

	arrHundred := []string{" ", "C", "CC", "CCC", "CD", "D", "DC", "DCC", "DCCC", "CM"}

	arrThausen := []string{" ", "M", "MM", "MMM", "MMMM", "MMMMM", "MMMMMM", "MMMMMMM", "MMMMMMMM", "MMMMMMMMM"}

	if arr[1] == 0 {
		fmt.Println(arrOne[arr[0]])
	}
	if arr[2] == 0 {
		fmt.Println(arrTen[arr[1]], arrOne[arr[0]])
	}
	if arr[3] == 0 {
		fmt.Println(arrHundred[arr[2]], arrTen[arr[1]], arrOne[arr[0]])
	} else {
		fmt.Println(arrThausen[arr[3]], arrHundred[arr[2]], arrTen[arr[1]], arrOne[arr[0]])
	}

}
