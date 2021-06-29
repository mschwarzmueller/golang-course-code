package main

import (
	"fmt"
	"io/fs"
	"io/ioutil"
)

type logger interface {
	log()
}

type logData struct {
	message  string
	fileName string
}

func (lData *logData) log() {
	err := ioutil.WriteFile(lData.fileName, []byte(lData.message), fs.FileMode(0644))

	if err != nil {
		fmt.Println("Storing the log data failed!")
		message := err.Error()
		fmt.Println(message)
	}
}

type loggableString string

func (text loggableString) log() {
	fmt.Println(text)
}

func main() {
	userLog := &logData{"User logged in", "user-log.txt"}
	// do more work ...
	createLog(userLog)

	message := loggableString("[INFO] User created!")
	// do more work ...
	createLog(message)

	// reader := bufio.NewReader(os.Stdin)
	outputValue(message)
	outputValue(userLog)

	firstNumber := 5
	secondNumber := 10.1
	numbers := []int{1, 2, 3}

	doubledFirstNumber := double(firstNumber)
	doubledSecondNumber := double(secondNumber)
	doubledNumbers := double(numbers)
	doubledString := double("Test")

	fmt.Println(doubledFirstNumber)
	fmt.Println(doubledSecondNumber)
	fmt.Println(doubledNumbers)
	fmt.Println(doubledString)
}

func createLog(data logger) {
	// more things we do
	data.log()
}

func outputValue(value interface{}) {
	fmt.Println(value)
}

func double(value interface{}) interface{} {
	switch val := value.(type) {
	case int:
		return val * 2
	case float64:
		return val * 2
	case []int:
		newNumbers := append(val, val...)
		return newNumbers
	default:
		return ""
	}
}
