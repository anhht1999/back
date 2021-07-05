package handlers

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"
)

const (
	minus = "minus"
)

func Minus(w http.ResponseWriter, req *http.Request) {
	params := req.URL.Query()
	if minus, ok := params[minus]; ok {
		fmt.Fprintf(w, "minus %s", minus[0])

		text := strings.Split(minus[0], "-")

		var arrConvert []int
		for _, val := range text {

			value, _ := strconv.Atoi(val)
			arrConvert = append(arrConvert, value)
		}

		fmt.Fprintf(w, "= %d", arrConvert[0] - arrConvert[1])

	} else {
		fmt.Fprintln(w, `Hi guys. I don't know your name because you don't enter the your_name query param`)
	}
}
