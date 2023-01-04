package main

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"

	"calc/numerals"
)

var (
	errorReadInput           = errors.New("an error occured while reading input")
	errorEmptyInput          = errors.New("input is empty")
	errorRomanNegative       = errors.New("roman result must be positive")
	errorNumeralSystem       = errors.New("mixing of different numeral systems")
	errorMathOperation       = errors.New("not a valid mathematical operation")
	errorMathOperationFormat = errors.New("format of mathematical operation does not implement the requirements: two operands (1 - 10) and one operator (+, -, /, *)")

	printOut = "Output:\n%v\n"
)

type calculator struct {
	operations map[string]func() int
	operator   string
	a, b       int
	reader     *bufio.Reader
	roman      bool
}

func newCalculator() *calculator {
	var c calculator
	c.reader = bufio.NewReader(os.Stdin)
	c.operations = map[string]func() int{
		"+": c.sum,
		"-": c.subtract,
		"*": c.multiply,
		"/": c.divide,
	}
	return &c
}

func (c *calculator) sum() int {
	return c.a + c.b
}
func (c *calculator) subtract() int {
	return c.a - c.b
}
func (c *calculator) multiply() int {
	return c.a * c.b
}
func (c *calculator) divide() int {
	return c.a / c.b
}

func (c *calculator) validateInput(s string) error {
	a, b, err := validateOperation(c, s)
	if err != nil {
		return err
	}

	err = validateOperands(c, a, b)
	if err != nil {
		return err
	}
	if c.a > 10 || c.b > 10 {
		return errorMathOperationFormat
	}

	return nil
}

func validateOperation(c *calculator, s string) (string, string, error) {
	var opCount int = 0
	var a, b string

	for i := 0; i < len(s); i++ {
		for k := range c.operations {
			if opCount >= 2 {
				return "", "", errorMathOperationFormat
			}
			if string(s[i]) == k {
				opCount += 1
				a, b = s[0:i], s[i+1:]
				c.operator = k
			}
		}
	}

	return a, b, nil
}

func validateOperands(c *calculator, a, b string) error {
	var errAa, errBa, errAr, errBr error
	c.a, errAa = strconv.Atoi(a)
	c.b, errBa = strconv.Atoi(b)

	if errAa != nil || errBa != nil {
		c.a, errAr = numerals.Rtoi(a)
		c.b, errBr = numerals.Rtoi(b)
		if errAr != nil || errBr != nil {
			if (errAa != nil && errAr != nil) || (errBa != nil && errBr != nil) {
				return errorMathOperation
			}
			return errorNumeralSystem
		}
		c.roman = true
		return nil
	}
	c.roman = false
	return nil
}

func calculate(c *calculator) error {
	for {
		fmt.Printf("Input:\n")
		input, err := c.reader.ReadString('\n')

		if err == io.EOF { // system interrupt signal (ctrl + C)
			return nil
		}
		if err != nil {
			return errorReadInput
		}
		if len(input) <= 2 {
			fmt.Printf("%v\n", errorEmptyInput)
			continue
		}

		input = strings.ReplaceAll(input, " ", "")
		input = strings.TrimRight(input, "\n\r")

		err = c.validateInput(input)
		if err != nil {
			return err
		}

		result := c.operations[c.operator]()

		switch c.roman {
		case true:
			if result < 1 {
				return errorRomanNegative
			}
			fmt.Printf(printOut, numerals.Itor(result))
		case false:
			fmt.Printf(printOut, result)
		}
	}
}

func main() {
	err := calculate(newCalculator())

	if err != nil {
		fmt.Printf(printOut, err)
	}
}
