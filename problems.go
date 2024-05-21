package main

import "fmt"

func AverageInArray(slice []int) int {
	sum := 0
	for i := 0; i < len(slice); i++ {
		sum += slice[i]
	}
	return sum / len(slice)
}
func Age(number int) {
	if number > 100 {
		fmt.Println("Exceeded the age limit")
	}
	if number < 0 {
		fmt.Println("no negative age is applicable")
	}
	if number < 18 {
		fmt.Println("minor")
	} else if number < 25 {
		fmt.Println("Young Adult")
	} else {
		fmt.Println("Adult")
	}

}
func ReverseString(str string) string {
	rns := []rune(str)
	for i, j := 0, len(str)-1; i < j; i, j = i+1, j-1 {
		rns[i], rns[j] = rns[j], rns[i]
	}

	return string(rns)

}
func LargestNumber(arr []int) int {
	max := arr[0]
	for i := 1; i < len(arr); i++ {
		if arr[i] > max {
			max = arr[i]
		}
	}
	return max
}

type Counter struct {
	count int
}

func (counter Counter) add() Counter {
	counter.count++
	return counter
}
func (counter Counter) substract() Counter {
	counter.count--
	return counter
}
func (counter Counter) display() Counter {
	fmt.Println(counter)
	return counter
}

func main() {
	slice := []int{1, 2, 3, 4, 5}
	fmt.Println(AverageInArray(slice))

	Age(23)
	fmt.Println(ReverseString("hello"))
	arr := []int{3, 2, 4, 5, 67, 5}
	fmt.Println(LargestNumber(arr))
// counter program
	var counter = Counter{count: 0}
	counter = counter.add()
	counter = counter.add()
	counter = counter.display()
	counter = counter.substract()
	counter = counter.display()
	// fmt.Println(counter.count)
}
