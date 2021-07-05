package main

import (
	"fmt"
	"io/ioutil"
	"log"

	"github.com/anhht1999/lec-02/algorithm"
	"github.com/anhht1999/lec-02/file"
)


func main()  {
	
	//Doc file test.txt
	content, err := ioutil.ReadFile("test.txt")

	if err != nil {
		log.Fatal(err)
	}
	text := string(content)

	fmt.Println(text)

	unSort := file.convertStringToInt(text)
	search := []int{3, 2}

	min, max := file.findMinAndMax(unSort)
	fmt.Println("Min: ", min)
	fmt.Println("Max: ", max)

	fmt.Println("Bubble Sorting:")
	fmt.Println(algorithm.bubbleSort(unSort))
	fmt.Println("Bubble Sorting:")
	fmt.Println(algorithm.mergeSort(unSort))
	fmt.Println("Quick Sorting:")
	fmt.Println(algorithm.quickSort(unSort, 0, len(unSort)-1))
	fmt.Println("Is Prime?:")
	fmt.Println(file.isPrime(unSort))
	fmt.Println("Check Number")
	fmt.Println(file.checkNumber(search, unSort))

	file.avg(unSort)
}