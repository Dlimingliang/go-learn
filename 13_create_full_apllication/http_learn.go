package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
)

var urls = []string{
	"http://www.baidu.com/",
}

func main() {
	http.HandleFunc("/", helloServer)
	http.HandleFunc("/add", add)
	err := http.ListenAndServe("localhost:9090", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err.Error())
	}

	//for _, url := range urls {
	//	resp, err := http.Head(url)
	//	if err != nil {
	//		fmt.Println("Error:", url, err)
	//	}
	//	fmt.Println(url, ": ", resp.Status)
	//}
}

func helloServer(w http.ResponseWriter, req *http.Request) {
	fmt.Println("Inside HelloServer handler")
	fmt.Fprintf(w, "Hello,"+req.URL.Path[1:])
}

func add(w http.ResponseWriter, req *http.Request) {
	_ = req.ParseForm()
	a, _ := strconv.Atoi(req.Form["a"][0])
	b, _ := strconv.Atoi(req.Form["b"][0])
	w.Header().Set("Content-Type", "application/json")
	jData, _ := json.Marshal(map[string]int{
		"data": a + b,
	})
	w.Write(jData)
}
