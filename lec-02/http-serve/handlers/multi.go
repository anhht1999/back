package handlers

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"
)

const (
	multi = "multi"
)

func Multi(w http.ResponseWriter, req *http.Request) {
	params := req.URL.Query()

	if multi, ok := params[multi]; ok {
		fmt.Fprintf(w, "multi %s", multi[0])

		text := strings.Split(multi[0], "*")

		var arrConvert []int
		for _, val := range text {

			value, _ := strconv.Atoi(val)
			arrConvert = append(arrConvert, value)
		}

		fmt.Fprintf(w, "= %d", arrConvert[0]*arrConvert[1])

	} else {
		fmt.Fprintln(w, `Hi guys. I don't know your name because you don't enter the your_name query param`)
	}
}
