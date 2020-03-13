package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

//strconv.Atoi, os, append, fmt, len, cast, strings.Split

func main() {
	if len(os.Args[1:]) != 1 {
		fmt.Println("Error")
	} else {

		mem := strings.Split(os.Args[1], " ")
		var mem2 []string
		for _, v := range mem {
			if v != "" {
				mem2 = append(mem2, v)
			}
		}
		if !isVal(mem2) {
			fmt.Println("Error")
			return
		}
		fmt.Println(rpn(mem2))

	}

}

func rpn(mem2 []string) int {

	for isOp(mem2[len(mem2)-1]) {
		retNum := 0
		operator := ""
		oper := false

		num1 := 0
		firstNum := false

		num2 := 0
		secondNum := false

		order := 0

		for i := 0; i < len(mem2); i++ {
			if isOp(mem2[i]) && !oper {
				operator = mem2[i]
				order = i
				mem2[i] = ""
				oper = true
			}

			if oper {
				break
			}
		}

		for i := order; i >= 0; i-- {

			if mem2[i] == "" {
				continue
			}
			if !firstNum && mem2[i] != "" {
				num1, _ = strconv.Atoi(mem2[i])
				mem2[i] = ""
				firstNum = true
				continue
			}
			if !secondNum && mem2[i] != "" {
				num2, _ = strconv.Atoi(mem2[i])
				mem2[i] = ""
				secondNum = true
				continue
			}
			if secondNum && firstNum {
				break
			}

		}

		if oper && firstNum && secondNum {
			retNum = num1 + num2
			switch operator {
			case "+":
				retNum = num2 + num1
			case "-":
				retNum = num2 - num1
			case "*":
				retNum = num2 * num1
			case "/":
				retNum = num2 / num1
			case "%":
				retNum = num2 % num1
			}
			mem2[order] = strconv.Itoa(retNum)
		}

	}
	retNum, _ := strconv.Atoi(mem2[len(mem2)-1])
	return retNum
}

func isOp(s string) bool {
	if s == "+" || s == "-" || s == "*" || s == "/" || s == "%" {
		return true
	}
	return false
}

func isVal(arr []string) bool {

	no := 0
	nn := 0

	if !isOp(arr[len(arr)-1]) {
		return false
	}

	for _, v := range arr {
		if isOp(v) {
			no++
		} else {
			_, err := strconv.Atoi(v)
			if err != nil {
				return false
			}
			nn++
		}
	}
	if nn-no != 1 {
		return false
	}

	return true

}
