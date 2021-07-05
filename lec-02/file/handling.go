package file

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

func CheckPrimeNumber(num int) bool {

	if num < 2 {
		return false
	}

	for i := 2; i <= int(math.Sqrt(float64(num))); i++ {
		if num%i == 0 {
			return false
		}
	}
	return true
}

func IsPrime(arr []int) []int {
	var prime []int
	for i := 0; i < len(arr); i++ {
		if CheckPrimeNumber(arr[i]) {
			prime = append(prime, arr[i])
		}
	}

	return prime
}
func CheckNumber(search []int, arr []int) []int {

	var numberEsixt []int
	for i := 0; i < len(arr); i++ {
		for j := 0; j < len(search); j++ {
			if arr[i] == search[j] {
				numberEsixt = append(numberEsixt, arr[i])
			}
		}
	}

	return numberEsixt
}

func ConvertStringToInt(arr string) []int {

	var arrConvert []int
	text := strings.Split(arr, ",")

	for _, val := range text {

		value, _ := strconv.Atoi(val)
		arrConvert = append(arrConvert, value)
	}

	return arrConvert
}

func FindMinAndMax(a []int) (min int, max int) {
	min = a[0]
	max = a[0]
	for _, value := range a {
		if value < min {
			min = value
		}
		if value > max {
			max = value
		}
	}
	return min, max
}

func Avg(arr []int) {
	count := 0
	for _, val := range arr {
		count += val
	}
	fmt.Println("trung binh la:", float32(count)/float32(len(arr)))
}
