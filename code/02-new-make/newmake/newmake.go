package newmake

import "fmt"

func main() {
	// new & make

	number := new(int)
	fmt.Println(number)
	fmt.Println(*number)

	anotherNumber := 0
	numberAddress := &anotherNumber

	fmt.Println(numberAddress)
	fmt.Println(*numberAddress)

	// hobbies := []string{"Sports", "Reading"}

	hobbies := make([]string, 2, 10)
	moreHobbies := new([]string)
	// evenMoreHobbies := []string{} // make([]string)

	fmt.Println(*moreHobbies)

	*moreHobbies = append(*moreHobbies, "Sports")

	fmt.Println(*moreHobbies)
	// aMap := make(map[string]int, 5)

	hobbies[0] = "Sports"
	hobbies[1] = "Reading"

	hobbies = append(hobbies, "Cooking", "Dancing")

	fmt.Println(hobbies)
}
