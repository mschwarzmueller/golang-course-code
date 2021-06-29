package main

import (
	"fmt"
	"math/rand"
	"time"
)

var source = rand.NewSource(time.Now().Unix())
var randN = rand.New(source)

func main() {
	x := make(chan int)
	y := make(chan int)
	limiter := make(chan int, 3)

	go generateValue(x, limiter)
	go generateValue(y, limiter)

	var a int
	var b int

	select {
	case a = <-x:
		fmt.Printf("X finished first, value is %v", a)
	case b = <-y:
		fmt.Printf("Y finished first, value is %v", b)
	}
	// go generateValue(c, limiter)
	// go generateValue(c, limiter)

	// sum := 0
	// i := 0

	// for num := range c {
	// 	sum += num
	// 	i++
	// 	if i == 4 {
	// 		close(c)
	// 	}
	// }

	// fmt.Println(sum)
}

func generateValue(c chan int, limit chan int) int {
	limit <- 1
	fmt.Println("Generating value...")
	sleepTime := randN.Intn(3)
	time.Sleep(time.Duration(sleepTime) * time.Second)

	result := randN.Intn(10)
	c <- result
	<-limit
	return result
}

// func main() {
// 	greet()
// 	storeData("This is some dummy data!", "dummy-data.txt")

// 	channel := make(chan int)

// 	go storeMoreData(50000, "50000_1.txt", channel)
// 	go storeMoreData(50000, "50000_2.txt", channel)

// 	<-channel
// 	<-channel
// }

// func greet() {
// 	fmt.Println("Hi there!")
// }

// func storeData(storableText string, fileName string) {
// 	file, err := os.OpenFile(fileName,
// 		os.O_CREATE|os.O_APPEND|os.O_WRONLY,
// 		0666,
// 	)

// 	if err != nil {
// 		fmt.Println("Creating the file failed. Exiting.")
// 		return
// 	}

// 	defer file.Close()

// 	_, err = file.WriteString(storableText)

// 	if err != nil {
// 		fmt.Println("Writing to the file failed.")
// 	}
// }

// func storeMoreData(lines int, fileName string, c chan int) {
// 	for i := 0; i < lines; i++ {
// 		text := fmt.Sprintf("Line %v - Dummy Data\n", i)
// 		storeData(text, fileName)
// 	}

// 	fmt.Printf("-- Done storing %v lines of text --\n", lines)
// 	c <- 1
// }
