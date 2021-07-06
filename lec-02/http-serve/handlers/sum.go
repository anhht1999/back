package handlers

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"
)

const (
	sum = "sum"
)

func Sum(w http.ResponseWriter, req *http.Request) {
	params := req.URL.Query()

	if sum, ok := params[sum]; ok {
		fmt.Fprintf(w, "sum %s", sum[0])

		text := strings.Split(sum[0], " ")

		var arrConvert []int
		for _, val := range text {

			value, _ := strconv.Atoi(val)
			arrConvert = append(arrConvert, value)
		}

		fmt.Fprintf(w, "= %d", arrConvert[0]+arrConvert[1])

	} else {
		fmt.Fprintln(w, `Hi guys. I don't know your name because you don't enter the your_name query param`)
	}
}
