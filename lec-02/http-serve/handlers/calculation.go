package handlers

import (
	"fmt"
	"strconv"
	"net/http"
	
	"../calculations"
)


func Handler(w http.ResponseWriter, r *http.Request) {

	//lấy tham số từ url
	calculation := r.URL.Query().Get("calculation")
	operator1, _ := strconv.ParseFloat(r.URL.Query().Get("operator1"), 64)
	operator2, _ := strconv.ParseFloat(r.URL.Query().Get("operator2"), 64)

	//hiển thị kết quả
	fmt.Fprintln(w, "Phép tính: ", operator1,calculation,operator2)
	fmt.Fprintln(w,"Kết quả: ", calculations.Calculate(calculation, operator1, operator2))

}