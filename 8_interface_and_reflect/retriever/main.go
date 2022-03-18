package main

import (
	"fmt"
	"github.com/Dlimingliang/go-learn/8_interface_and_reflect/retriever/fake"
)

const url = "http://baidu.com"

type Retriever interface {
	Get(str string) string
}

func download(r Retriever) string {
	return r.Get(url)
}

type Poster interface {
	Post(url string, param map[string]string)
}

func post(p Poster) {
	p.Post(url, map[string]string{"aa": "bb"})
}

type RetrieverPost interface {
	Retriever
	Poster
}

func session(rp RetrieverPost) string {
	rp.Post(url, map[string]string{"context": "Another fake context"})
	return rp.Get(url)
}

func main() {
	//var retriever Retriever
	//retriever = real.Retriever{}
	//fmt.Println(download(retriever))
	var r Retriever
	retriever := &fake.Retriever{Context: "I am fake context"}
	fmt.Println("stringer", retriever)
	r = retriever
	fmt.Println(r.Get(url))
	fmt.Println(session(retriever))
}
