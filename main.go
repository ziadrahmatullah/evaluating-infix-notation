package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type node struct {
	val  string
	next *node
}

type stack struct {
	top *node
	len int
}

func (S *stack) push(item string) {
	newNode := &node{val: item, next: nil}
	if S.top == nil {
		S.top = newNode
		S.len++
	} else {
		newNode.next = S.top
		S.top = newNode
		S.len++
	}
}

func (S *stack) pop() string {
	if S.top != nil {
		popped := S.top.val
		S.top = S.top.next
		S.len--
		return popped
	}
	return ""
}

func operation(operator, number1, number2 string) string {
	a, _ := strconv.Atoi(number1)
	b, _ := strconv.Atoi(number2)
	var result int
	if operator == "+" {
		result = a + b
	} else if operator == "-" {
		result = a - b
	} else if operator == "/" {
		result = a / b
	} else if operator == "*" {
		result = a * b
	}
	str := strconv.Itoa(result)
	return str
}

func calculateIN(input []string) int {
	N := stack{} // Number Stack
	O := stack{} // Operator Stack
	for _, char := range input {
		if _, err := strconv.Atoi(char); err == nil {
			N.push(char)
		} else if char != "(" && char != ")"{
			for O.top != nil && (O.top.val == "*" || O.top.val == "/"){
				operator := O.pop()
				N.push(operation(operator, N.pop(), N.pop()))
			}
			O.push(char)
		}else if char == ")" {
			operator := O.pop()
			for operator != "(" {
				N.push(operation(operator, N.pop(), N.pop()))
				operator = O.pop()
			}
		} else { // + - ( )
			O.push(char)
		}
	}
	for O.top != nil {
		operator := O.pop()
		N.push(operation(operator, N.pop(), N.pop()))
	}

	result, _ := strconv.Atoi(N.pop())
	return result
}

func readInput() []string {
	reader := bufio.NewReader(os.Stdin)
	line, err := reader.ReadString('\n')
	if err != nil {
		log.Fatal(err)
	}
	return strings.Fields(line)
}

func main() {
	fmt.Print("Input: ")
	polishNotation := readInput()
	fmt.Println("Output: ", calculateIN(polishNotation))
}