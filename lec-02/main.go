package main

import (
	"fmt"
	"io/ioutil"
	"log"

	"github.com/anhht1999/back/lec-02/file"
	"github.com/anhht1999/back/lec-02/algorithm"
)


func main()  {
	
	//Doc file test.txt
	content, err := ioutil.ReadFile("test.txt")

	if err != nil {
		log.Fatal(err)
	}
	text := string(content)

	fmt.Println(text)

	unSort := file.ConvertStringToInt(text)
	search := []int{4,13}

	min, max := file.FindMinAndMax(unSort)
	fmt.Println("Min: ", min)
	fmt.Println("Max: ", max)

	fmt.Println("Bubble Sorting:")
	fmt.Println(algorithm.BubbleSort(unSort))
	fmt.Println("Merge Sorting:")
	fmt.Println(algorithm.MergeSort(unSort))
	fmt.Println("Quick Sorting:")
	fmt.Println(algorithm.QuickSort(unSort, 0, len(unSort)-1))
	fmt.Println("Is Prime?:")
	fmt.Println(file.IsPrime(unSort))
	fmt.Println("Check Number")
	fmt.Println(file.CheckNumber(search, unSort))

	file.Avg(unSort)
}