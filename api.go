package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

// สร้าง struct ไว้เพื่อทำ js
type Movie struct {
	Id     string   `json:"id"`
	Title  string   `json:"title"`
	Year   int      `json:"year"`
	Rating float64  `json:"rating"`
	Gs     []string `json:"gs"`
	Iss    bool     `json:"iss"`
}

// func main จะอยู่ล่างสุดด เเพื่อง่ายต่อการทำงาน
var movie []Movie

func moviesHandler(w http.ResponseWriter, req *http.Request) {
	// req.Methot คือถามว่า mehod มาจากอะไร
	method := req.Method
	// วิธ๊ ยกประเเภท
	// เช็ค error
	if method == "POST" {
		body, err := ioutil.ReadAll(req.Body)
		if err != nil {
			// ผิดพลาดตรงไหนบ้างให้ใส่ coddeนี้
			w.WriteHeader(http.StatusBadRequest)
			fmt.Fprintf(w, "Error : %v ", err)
			return
		}
		//fmt.Println(string(body))
		// เปลี่ยนจาก js เป็น string
		t := Movie{}
		err = json.Unmarshal(body, &t)
		if err != nil {
			// ผิดพลาดตรงไหนบ้างให้ใส่ coddeนี้
			w.WriteHeader(http.StatusOK)
			fmt.Fprintf(w, "Error : %s ", err)
			return
		}
		// อันนี้คือนำของไปเก็บไว้ใน movie
		movie = append(movie, t)
		fmt.Fprintf(w, "HI %s save movies is POST", method)
		// ใส่เเื่อไม่ให้ลงไปด้านล่างต่อ
		return

	}

	// เปลี่ยนค่าจาก sstring เป็น js
	b, err := json.Marshal(movie)
	if err != nil {
		// ผิดพลาดตรงไหนบ้างให้ใส่ coddeนี้
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "Error : %s ", err)
		return
	}
	// กำหนดการส่งค่าให้ส่งเป็ฯ js
	w.Header().Set("Contect-Type", "application/json; charset=utf-8")
	// ส่งไปหา User
	// fmt.Fprintf(w, "%s ", string(b))
	// แต่วิธีที่ง่ายที่สุดคือ
	w.Write(b)
}
func main() {
	http.HandleFunc("/movies", moviesHandler)

	err := http.ListenAndServe("localhost:2565", nil)
	log.Fatal(err)
}
