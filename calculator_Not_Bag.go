package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

var znac string

func main() {

	fmt.Println("Нажмите s для старта, q для выхода")

	scannRestart := bufio.NewScanner(os.Stdin)

	for scannRestart.Scan() {

		exit := scannRestart.Text()

		if exit == "q" {
			break
		} else {

			scannerString := bufio.NewScanner(os.Stdin)
			fmt.Println("введите выражение:")
			scannerString.Scan()

			str := scannerString.Text()
			str = strings.Join(strings.Fields(str), "")
			//arr := strings.Split(str, "")

			arrResult, swh := MV1(str)
			if swh == true {

				//fmt.Println(arrResult)
				arrSwh := MV2(arrResult)
				if arrSwh[0] == true && arrSwh[1] == true {

					//fmt.Println(arrSwh)
					x := arrResult[0]
					y := arrResult[1]

					hundlerRome(x, y)

				} else if arrSwh[0] == false && arrSwh[1] == false {

					//fmt.Println(arrSwh)
					x, err1 := strconv.Atoi(arrResult[0])
					if err1 != nil {
						log.Fatal(err1)
					}

					y, err2 := strconv.Atoi(arrResult[1])
					if err2 != nil {
						log.Fatal(err2)
					}

					if x != 0 && y != 0 {

						hundlerArabic(x, y)

					} else {
						panic("Ноль нельзя!")
					}

				} else {
					panic("Нельзя одновременно использовать две системы счисления")
				}

			} else {
				panic("Несуществующий знак")
			}

			fmt.Println("Нажмите s для старта, q для выхода")

		}

	}

}

func MV1(str string) ([]string, bool) {

	if strings.Contains(str, "+") {
		arr := strings.Split(str, "+")
		znac = "+"
		return arr, true
	}
	if strings.Contains(str, "-") {
		arr := strings.Split(str, "-")
		znac = "-"
		return arr, true
	}
	if strings.Contains(str, "*") {
		arr := strings.Split(str, "*")
		znac = "*"
		return arr, true
	}
	if strings.Contains(str, "/") {
		arr := strings.Split(str, "/")
		znac = "/"
		return arr, true
	} else {
		arr := strings.Split(str, "")
		return arr, false
	}

}

func MV2(arr []string) []bool {

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

	if len(retArr) > 2 || len(retArr) < 2 {
		panic("Неправильная запись")
	}
	return retArr

}

func hundlerArabic(x, y int) {

	if znac == "+" {
		fmt.Println(x + y)
	}
	if znac == "-" {
		fmt.Println(x - y)
	}
	if znac == "*" {
		fmt.Println(x * y)
	}
	if znac == "/" {
		fmt.Println(x / y)
	}

}

func hundlerRome(x, y string) {

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

	if znac == "+" {

		c := (RomeToArab[x] + RomeToArab[y])

		if c > 0 {

			translateToRome(c)

		} else {
			panic("Отрицательных чисел в римской системе нет!")
		}

	}
	if znac == "-" {

		c := (RomeToArab[x] - RomeToArab[y])

		if c > 0 {

			translateToRome(c)

		} else {
			panic("Отрицательных чисел в римской системе нет!")
		}

	}
	if znac == "*" {

		c := (RomeToArab[x] * RomeToArab[y])

		if c > 0 {

			translateToRome(c)

		} else {
			panic("Отрицательных чисел в римской системе нет!")
		}

	}
	if znac == "/" {

		c := (RomeToArab[x] / RomeToArab[y])

		if c > 0 {

			translateToRome(c)

		} else {
			panic("Отрицательных чисел в римской системе нет!")
		}

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

	result := strings.Join(strings.Fields(arrThausen[arr[3]]+arrHundred[arr[2]]+arrTen[arr[1]]+arrOne[arr[0]]), "")
	fmt.Println(result)

}
