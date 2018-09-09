package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
	"path/filepath"
	"quiz/question"
	"quiz/question/lifeexpectancy"
	"quiz/question/medianrents"
	"quiz/question/population/whatis"
	"strings"

	"github.com/gorilla/mux"
)

func StartServer(port int) {
	r := mux.NewRouter()
	r.HandleFunc("/api/questions/population", popQuestionFunc)
	r.HandleFunc("/{file:.*}", serveFunc)
	r.HandleFunc("/", serveFunc)
	srv := &http.Server{
		Addr:    fmt.Sprintf("0.0.0.0:%d", port),
		Handler: r, // Pass our instance of gorilla/mux in.
	}
	if err := srv.ListenAndServe(); err != nil {
		log.Println(err)
	}
}

func serveFunc(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	file := "/" + vars["file"]
	if file == "/" {
		file = "/index.html"
	}

	ext := filepath.Ext(file)
	if ext == "" {
		file = "/" + strings.Trim(file, "/") + "/index.html"
	}

	file = "public" + file

	log.Printf("Request for file: %s", file)

	var data []byte
	var err error

	data, err = ioutil.ReadFile(file)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	log.Printf("ext = %s", ext)

	switch ext {
	case ".html":
		w.Header().Set("Content-Type", "text/html")
	case ".css":
		w.Header().Set("Content-Type", "text/css")
	case ".js":
		w.Header().Set("Content-Type", "application/javascript")
	case ".json":
		w.Header().Set("Content-Type", "application/json")
	default:
		w.Header().Set("Content-Type", "text/html")
	}

	w.Write(data)

}

func popQuestionFunc(w http.ResponseWriter, r *http.Request) {

	rr := rand.Intn(100)

	var q *question.Question
	var err error

	if rr < 50 {
		q, err = whatis.New()
		log.Printf("Error: %v", err)
		if err != nil {
			w.WriteHeader(500)
			return
		}
	} else if rr < 75 {
		q, err = medianrents.New()
		log.Printf("Error: %v", err)
		if err != nil {
			w.WriteHeader(500)
			return
		}
	} else {
		q, err = lifeexpectancy.New()
		log.Printf("Error: %v", err)
		if err != nil {
			w.WriteHeader(500)
			return
		}
	}

	j, _ := json.Marshal(q)
	w.Header().Set("Content-type", "application/json")
	w.Write(j)

}
