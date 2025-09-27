package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

func main() {
	helloHandler := func(w http.ResponseWriter, req *http.Request) {
		io.WriteString(w, "Hello, World!\n")
	}

	postArticleHandler := func(w http.ResponseWriter, req *http.Request) {
		io.WriteString(w, "Posting Article...\n")
	}

	listArticlesHandler := func(w http.ResponseWriter, req *http.Request) {
		io.WriteString(w, "Article List\n")
	}

	getArticleHandler := func(w http.ResponseWriter, req *http.Request) {
		articleID := 1
		resString := fmt.Sprintf("Article No.%d\n", articleID)
		io.WriteString(w, resString)
	}

	niceArticleHandler := func(w http.ResponseWriter, req *http.Request) {
		io.WriteString(w, "Posting Nice...\n")
	}

	commentArticleHandler := func(w http.ResponseWriter, req *http.Request) {
		io.WriteString(w, "Posting Comment...\n")
	}

	http.HandleFunc("/hello", helloHandler)
	http.HandleFunc("/article", postArticleHandler)
	http.HandleFunc("/article/list", listArticlesHandler)
	http.HandleFunc("/article/1", getArticleHandler)
	http.HandleFunc("/article/nice", niceArticleHandler)
	http.HandleFunc("/comment", commentArticleHandler)

	log.Println("server start at port 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
