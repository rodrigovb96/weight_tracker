package main

import (
	"github.com/rodrigovb96/weight_tracker/logger"

	"net/http"
	"os"
	"strconv"
	"html/template"
	"path/filepath"

	"fmt"

)


func logWeight(w http.ResponseWriter, r * http.Request) {

	rawData := r.URL.Path[len("/input/"):]

	weight, err := strconv.ParseFloat(string(rawData),64)

	if err != nil {
		panic(err)
	}

	if logger.LogToFile(float32(weight)) {
		w.Write([]byte("Weight Logged with success"))
	} else {
		w.Write([]byte("ERROR!"))
	}



}

func determineListenAddress() (string,error) {
	port := os.Getenv("PORT")
	if port == "" {
		return "", fmt.Errorf("$PORT not set")
	}

	return ":" + port, nil
}

func plotServer() {
	plotServer := http.NewServeMux()
	plotServer.Handle("/files/", http.StripPrefix("/files/", http.FileServer(http.Dir("./files"))))
	plotServer.HandleFunc("/", func(w http.ResponseWriter, r * http.Request) {
		filePath, _ := filepath.Abs("files/static/graph.html")
		tmpl, err := template.ParseFiles(filePath)

		if err != nil {
			http.Error(w,err.Error(), http.StatusInternalServerError)
			return
		}

		type path struct {
			FilePath string
		}

		if err := tmpl.Execute(w,path{ FilePath: r.URL.Path[1:]}); err != nil {
			http.Error(w,err.Error(), http.StatusInternalServerError)
			return
		}
	})

	if err := http.ListenAndServe(":8090", plotServer); err != nil {
		panic(err)
	}

}


func main() {

	insertServer := http.NewServeMux()
	insertServer.HandleFunc("/input/",logWeight)

	port, err := determineListenAddress()

	go plotServer();

	if err != nil {
		panic(err)
	}

	if err := http.ListenAndServe(port, insertServer); err != nil {
		panic(err)
	}

}
