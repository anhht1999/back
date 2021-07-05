package handlers

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"
)

const (
	div = "div"
)

func Div(w http.ResponseWriter, req *http.Request) {
	params := req.URL.Query()
	if div, ok := params[div]; ok {
		fmt.Fprintf(w, "div %s", div[0])

		text := strings.Split(div[0], "-")

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
