/*
Упражнение 2.2. Напишите программу общего назначения для преобразования
единиц, аналогичную cf, которая считывает числа из аргументов командной
строки (или из стандартного ввода, если аргументы командной строки
отсутствуют) и преобразует каждое число в другие единицы, как
температуру — в градусы Цельсия и Фаренгейта, длину — в футы и метры,
вес — в фунты и килограммы и т.д.

В командной строке ожидаются параметры:

	t/l/w - температура, длина, вес;
	список значений для перевода.

Пример:

	task02 w 0 5 10
*/
package main

import (
	"bufio"
	"errors"
	"fmt"
	"golearn/chapter02/task02/conv"
	"os"
	"strconv"
	"strings"
)

func main() {
	typeOp, vals, err := getParams(os.Args[1:])
	if err != nil {
		fmt.Fprintf(os.Stderr, "main: %v\n", err)
		os.Exit(1)
	}
	print(typeOp, vals)
}

func getParams(args []string) (string, []string, error) {
	if len(args) == 0 {
		fmt.Println("Enter type op and values")
		text, err := bufio.NewReader(os.Stdin).ReadString('\n')
		if err != nil {
			return "", nil, err
		}
		args = strings.Split(text[:len(text)-2], " ")
	}

	if len(args) <= 1 {
		fmt.Println("getParams:len(args) == 1")
		return "", nil, errors.New("getParams:input data incorrect")
	} else {
		return args[0], args[1:], nil
	}
}

func funcPrint(typeOp string) (func(float64), error) {
	switch typeOp {
	case "t":
		return conv.PrintTKC, nil
	case "l":
		return conv.PrintLMFt, nil
	case "w":
		return conv.PrintWKgLb, nil
	default:
		return nil, errors.New("funcPrint:undefined type oper")
	}
}

func print(typeOp string, vals []string) {
	f, err := funcPrint(typeOp)
	if err != nil {
		fmt.Fprintf(os.Stderr, "print: %v\n", err)
		os.Exit(1)
	}
	for _, arg := range vals {
		t, err := strconv.ParseFloat(arg, 64)
		if err != nil {
			fmt.Fprintf(os.Stderr, "print: %v\n", err)
			os.Exit(1)
		}
		f(t)
	}
}
